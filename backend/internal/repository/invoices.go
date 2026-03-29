package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"

	"github.com/DigitLock/invoice-generator/backend/internal/database/sqlc"
	"github.com/DigitLock/invoice-generator/backend/internal/dto"
)

var allowedTransitions = map[string][]string{
	"draft":          {"sent", "cancelled"},
	"sent":           {"partially_paid", "paid", "cancelled"},
	"partially_paid": {"paid", "cancelled"},
}

type InvoiceRepository struct {
	q    *sqlc.Queries
	pool *pgxpool.Pool
}

func (r *InvoiceRepository) GetByID(ctx context.Context, id int64, familyID string) (sqlc.Invoice, error) {
	return r.q.GetInvoice(ctx, sqlc.GetInvoiceParams{ID: id, FamilyID: familyID})
}

func (r *InvoiceRepository) List(ctx context.Context, familyID string, page, pageSize int) ([]sqlc.ListInvoicesRow, int64, error) {
	offset := (page - 1) * pageSize
	rows, err := r.q.ListInvoices(ctx, sqlc.ListInvoicesParams{
		FamilyID: familyID,
		Limit:    int32(pageSize),
		Offset:   int32(offset),
	})
	if err != nil {
		return nil, 0, err
	}
	total, err := r.q.CountInvoices(ctx, familyID)
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (r *InvoiceRepository) Create(ctx context.Context, userID string, familyID string, req dto.CreateInvoiceRequest) (sqlc.Invoice, []sqlc.InvoiceItem, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return sqlc.Invoice{}, nil, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	qtx := r.q.WithTx(tx)

	issueDate, err := time.Parse("2006-01-02", req.IssueDate)
	if err != nil {
		return sqlc.Invoice{}, nil, fmt.Errorf("invalid issue_date: %w", err)
	}
	var dueDateParam pgtype.Date
	if req.DueDate != nil && *req.DueDate != "" {
		dueDate, err := time.Parse("2006-01-02", *req.DueDate)
		if err != nil {
			return sqlc.Invoice{}, nil, fmt.Errorf("invalid due_date: %w", err)
		}
		dueDateParam = pgtype.Date{Time: dueDate, Valid: true}
	}

	invoiceNumber, err := r.generateNumber(ctx, qtx, userID, issueDate)
	if err != nil {
		return sqlc.Invoice{}, nil, fmt.Errorf("generate invoice number: %w", err)
	}

	vatRate := decimal.NewFromFloat(0)
	if req.VatRate != "" {
		vatRate, err = decimal.NewFromString(req.VatRate)
		if err != nil {
			return sqlc.Invoice{}, nil, fmt.Errorf("invalid vat_rate: %w", err)
		}
	}

	subtotal := decimal.Zero
	type itemCalc struct {
		desc  string
		qty   decimal.Decimal
		price decimal.Decimal
		total decimal.Decimal
	}
	items := make([]itemCalc, 0, len(req.Items))
	for _, it := range req.Items {
		qty, err := decimal.NewFromString(it.Quantity)
		if err != nil {
			return sqlc.Invoice{}, nil, fmt.Errorf("invalid quantity: %w", err)
		}
		price, err := decimal.NewFromString(it.UnitPrice)
		if err != nil {
			return sqlc.Invoice{}, nil, fmt.Errorf("invalid unit_price: %w", err)
		}
		total := qty.Mul(price).Round(2)
		subtotal = subtotal.Add(total)
		items = append(items, itemCalc{desc: it.Description, qty: qty, price: price, total: total})
	}

	vatAmount := subtotal.Mul(vatRate.Div(decimal.NewFromInt(100))).Round(2)
	grandTotal := subtotal.Add(vatAmount).Round(2)

	invoice, err := qtx.CreateInvoice(ctx, sqlc.CreateInvoiceParams{
		UserID:            userID,
		FamilyID:          familyID,
		CompanyID:         req.CompanyID,
		ClientID:          req.ClientID,
		BankAccountID:     req.BankAccountID,
		InvoiceNumber:     invoiceNumber,
		IssueDate:         pgtype.Date{Time: issueDate, Valid: true},
		DueDate:           dueDateParam,
		Currency:          req.Currency,
		VatRate:           vatRate,
		Subtotal:          subtotal,
		VatAmount:         vatAmount,
		Total:             grandTotal,
		ContractReference: textFromPtr(req.ContractReference),
		ExternalReference: textFromPtr(req.ExternalReference),
		Notes:             textFromPtr(req.Notes),
	})
	if err != nil {
		return sqlc.Invoice{}, nil, fmt.Errorf("create invoice: %w", err)
	}

	createdItems := make([]sqlc.InvoiceItem, 0, len(items))
	for _, it := range items {
		item, err := qtx.CreateInvoiceItem(ctx, sqlc.CreateInvoiceItemParams{
			InvoiceID:   invoice.ID,
			Description: it.desc,
			Quantity:    it.qty,
			UnitPrice:   it.price,
			Total:       it.total,
		})
		if err != nil {
			return sqlc.Invoice{}, nil, fmt.Errorf("create item: %w", err)
		}
		createdItems = append(createdItems, item)
	}

	if err := tx.Commit(ctx); err != nil {
		return sqlc.Invoice{}, nil, fmt.Errorf("commit: %w", err)
	}
	return invoice, createdItems, nil
}

func (r *InvoiceRepository) Update(ctx context.Context, id int64, familyID string, req dto.UpdateInvoiceRequest) (sqlc.Invoice, []sqlc.InvoiceItem, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return sqlc.Invoice{}, nil, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	qtx := r.q.WithTx(tx)

	issueDate, err := time.Parse("2006-01-02", req.IssueDate)
	if err != nil {
		return sqlc.Invoice{}, nil, fmt.Errorf("invalid issue_date: %w", err)
	}
	var dueDateParam pgtype.Date
	if req.DueDate != nil && *req.DueDate != "" {
		dueDate, err := time.Parse("2006-01-02", *req.DueDate)
		if err != nil {
			return sqlc.Invoice{}, nil, fmt.Errorf("invalid due_date: %w", err)
		}
		dueDateParam = pgtype.Date{Time: dueDate, Valid: true}
	}

	vatRate := decimal.NewFromFloat(0)
	if req.VatRate != "" {
		vatRate, err = decimal.NewFromString(req.VatRate)
		if err != nil {
			return sqlc.Invoice{}, nil, fmt.Errorf("invalid vat_rate: %w", err)
		}
	}

	subtotal := decimal.Zero
	type itemCalc struct {
		desc  string
		qty   decimal.Decimal
		price decimal.Decimal
		total decimal.Decimal
	}
	items := make([]itemCalc, 0, len(req.Items))
	for _, it := range req.Items {
		qty, err := decimal.NewFromString(it.Quantity)
		if err != nil {
			return sqlc.Invoice{}, nil, fmt.Errorf("invalid quantity: %w", err)
		}
		price, err := decimal.NewFromString(it.UnitPrice)
		if err != nil {
			return sqlc.Invoice{}, nil, fmt.Errorf("invalid unit_price: %w", err)
		}
		total := qty.Mul(price).Round(2)
		subtotal = subtotal.Add(total)
		items = append(items, itemCalc{desc: it.Description, qty: qty, price: price, total: total})
	}

	vatAmount := subtotal.Mul(vatRate.Div(decimal.NewFromInt(100))).Round(2)
	grandTotal := subtotal.Add(vatAmount).Round(2)

	invoice, err := qtx.UpdateInvoice(ctx, sqlc.UpdateInvoiceParams{
		ID:                id,
		FamilyID:          familyID,
		CompanyID:         req.CompanyID,
		ClientID:          req.ClientID,
		BankAccountID:     req.BankAccountID,
		InvoiceNumber:     req.InvoiceNumber,
		IssueDate:         pgtype.Date{Time: issueDate, Valid: true},
		DueDate:           dueDateParam,
		Currency:          req.Currency,
		VatRate:           vatRate,
		Subtotal:          subtotal,
		VatAmount:         vatAmount,
		Total:             grandTotal,
		ContractReference: textFromPtr(req.ContractReference),
		ExternalReference: textFromPtr(req.ExternalReference),
		Notes:             textFromPtr(req.Notes),
	})
	if err != nil {
		return sqlc.Invoice{}, nil, fmt.Errorf("update invoice: %w", err)
	}

	if err := qtx.DeleteInvoiceItemsByInvoice(ctx, id); err != nil {
		return sqlc.Invoice{}, nil, fmt.Errorf("delete old items: %w", err)
	}

	createdItems := make([]sqlc.InvoiceItem, 0, len(items))
	for _, it := range items {
		item, err := qtx.CreateInvoiceItem(ctx, sqlc.CreateInvoiceItemParams{
			InvoiceID:   id,
			Description: it.desc,
			Quantity:    it.qty,
			UnitPrice:   it.price,
			Total:       it.total,
		})
		if err != nil {
			return sqlc.Invoice{}, nil, fmt.Errorf("create item: %w", err)
		}
		createdItems = append(createdItems, item)
	}

	if err := tx.Commit(ctx); err != nil {
		return sqlc.Invoice{}, nil, fmt.Errorf("commit: %w", err)
	}
	return invoice, createdItems, nil
}

func (r *InvoiceRepository) UpdateStatus(ctx context.Context, id int64, familyID string, newStatus string) (sqlc.Invoice, error) {
	invoice, err := r.q.GetInvoice(ctx, sqlc.GetInvoiceParams{ID: id, FamilyID: familyID})
	if err != nil {
		return sqlc.Invoice{}, fmt.Errorf("get invoice: %w", err)
	}

	allowed, ok := allowedTransitions[invoice.Status]
	if !ok {
		return sqlc.Invoice{}, fmt.Errorf("status '%s' is terminal and cannot be changed", invoice.Status)
	}

	valid := false
	for _, s := range allowed {
		if s == newStatus {
			valid = true
			break
		}
	}
	if !valid {
		return sqlc.Invoice{}, fmt.Errorf("transition from '%s' to '%s' is not allowed", invoice.Status, newStatus)
	}

	return r.q.UpdateInvoiceStatus(ctx, sqlc.UpdateInvoiceStatusParams{
		ID: id, FamilyID: familyID, Status: newStatus,
	})
}

func (r *InvoiceRepository) UpdateOverdue(ctx context.Context, id int64, familyID string, isOverdue bool) (sqlc.Invoice, error) {
	invoice, err := r.q.GetInvoice(ctx, sqlc.GetInvoiceParams{ID: id, FamilyID: familyID})
	if err != nil {
		return sqlc.Invoice{}, fmt.Errorf("get invoice: %w", err)
	}
	if invoice.Status == "draft" {
		return sqlc.Invoice{}, fmt.Errorf("overdue flag cannot be set on draft invoices")
	}
	return r.q.UpdateInvoiceOverdue(ctx, sqlc.UpdateInvoiceOverdueParams{
		ID: id, FamilyID: familyID, IsOverdue: isOverdue,
	})
}

func (r *InvoiceRepository) Delete(ctx context.Context, id int64, familyID string) error {
	return r.q.DeleteInvoice(ctx, sqlc.DeleteInvoiceParams{ID: id, FamilyID: familyID})
}

func (r *InvoiceRepository) GetItems(ctx context.Context, invoiceID int64) ([]sqlc.InvoiceItem, error) {
	return r.q.ListInvoiceItems(ctx, invoiceID)
}

func (r *InvoiceRepository) GetBankAccount(ctx context.Context, id int64) (sqlc.BankAccount, error) {
	return r.q.GetBankAccount(ctx, id)
}

func (r *InvoiceRepository) generateNumber(ctx context.Context, q *sqlc.Queries, userID string, issueDate time.Time) (string, error) {
	dateStr := issueDate.Format("02012006")
	prefix := fmt.Sprintf("INV-%s-", dateStr)
	pattern := prefix + "%"

	maxSeqRaw, err := q.GetMaxInvoiceSequence(ctx, sqlc.GetMaxInvoiceSequenceParams{
		UserID:        userID,
		InvoiceNumber: pattern,
	})
	if err != nil {
		return "", fmt.Errorf("get max sequence: %w", err)
	}

	maxSeq := int64(0)
	if maxSeqRaw != nil {
		switch v := maxSeqRaw.(type) {
		case int32:
			maxSeq = int64(v)
		case int64:
			maxSeq = v
		}
	}

	return fmt.Sprintf("%s%03d", prefix, maxSeq+1), nil
}
