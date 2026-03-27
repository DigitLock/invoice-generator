package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DigitLock/invoice-generator/backend/internal/database/sqlc"
)

type BankAccountRepository struct {
	q    *sqlc.Queries
	pool *pgxpool.Pool
}

func (r *BankAccountRepository) GetByID(ctx context.Context, id int64) (sqlc.BankAccount, error) {
	return r.q.GetBankAccount(ctx, id)
}

func (r *BankAccountRepository) ListByCompany(ctx context.Context, companyID int64) ([]sqlc.BankAccount, error) {
	return r.q.ListBankAccountsByCompany(ctx, companyID)
}

func (r *BankAccountRepository) Create(ctx context.Context, params sqlc.CreateBankAccountParams) (sqlc.BankAccount, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return sqlc.BankAccount{}, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	qtx := r.q.WithTx(tx)

	if params.IsDefault {
		if err := qtx.ClearDefaultBankAccount(ctx, sqlc.ClearDefaultBankAccountParams{
			CompanyID: params.CompanyID, ID: 0,
		}); err != nil {
			return sqlc.BankAccount{}, fmt.Errorf("clear default: %w", err)
		}
	}

	account, err := qtx.CreateBankAccount(ctx, params)
	if err != nil {
		return sqlc.BankAccount{}, fmt.Errorf("create: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return sqlc.BankAccount{}, fmt.Errorf("commit: %w", err)
	}
	return account, nil
}

func (r *BankAccountRepository) Update(ctx context.Context, id int64, params sqlc.UpdateBankAccountParams) (sqlc.BankAccount, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return sqlc.BankAccount{}, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	qtx := r.q.WithTx(tx)
	params.ID = id

	if params.IsDefault {
		companyID, err := qtx.GetBankAccountCompanyID(ctx, id)
		if err != nil {
			return sqlc.BankAccount{}, fmt.Errorf("get company: %w", err)
		}
		if err := qtx.ClearDefaultBankAccount(ctx, sqlc.ClearDefaultBankAccountParams{
			CompanyID: companyID, ID: id,
		}); err != nil {
			return sqlc.BankAccount{}, fmt.Errorf("clear default: %w", err)
		}
	}

	account, err := qtx.UpdateBankAccount(ctx, params)
	if err != nil {
		return sqlc.BankAccount{}, fmt.Errorf("update: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return sqlc.BankAccount{}, fmt.Errorf("commit: %w", err)
	}
	return account, nil
}

func (r *BankAccountRepository) Delete(ctx context.Context, id int64) error {
	hasInvoices, err := r.q.BankAccountHasNonDraftInvoices(ctx, id)
	if err != nil {
		return fmt.Errorf("check invoices: %w", err)
	}
	if hasInvoices {
		return fmt.Errorf("bank account is referenced by non-draft invoices and cannot be deleted")
	}
	return r.q.DeleteBankAccount(ctx, id)
}

func (r *BankAccountRepository) GetCompanyID(ctx context.Context, id int64) (int64, error) {
	return r.q.GetBankAccountCompanyID(ctx, id)
}
