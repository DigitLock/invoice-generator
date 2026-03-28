package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/DigitLock/invoice-generator/backend/internal/database/sqlc"
)

type CompanyRepository struct {
	q *sqlc.Queries
}

func (r *CompanyRepository) GetByID(ctx context.Context, id int64, familyID string) (sqlc.Company, error) {
	return r.q.GetCompany(ctx, sqlc.GetCompanyParams{ID: id, FamilyID: familyID})
}

func (r *CompanyRepository) List(ctx context.Context, familyID string) ([]sqlc.Company, error) {
	return r.q.ListCompanies(ctx, familyID)
}

func (r *CompanyRepository) Create(ctx context.Context, familyID string, name, contactPerson, address string, phone, vatNumber, regNumber *string) (sqlc.Company, error) {
	return r.q.CreateCompany(ctx, sqlc.CreateCompanyParams{
		FamilyID:      familyID,
		Name:          name,
		ContactPerson: contactPerson,
		Address:       address,
		Phone:         textFromPtr(phone),
		VatNumber:     textFromPtr(vatNumber),
		RegNumber:     textFromPtr(regNumber),
	})
}

func (r *CompanyRepository) Update(ctx context.Context, id int64, familyID string, name, contactPerson, address string, phone, vatNumber, regNumber *string) (sqlc.Company, error) {
	return r.q.UpdateCompany(ctx, sqlc.UpdateCompanyParams{
		ID:            id,
		FamilyID:      familyID,
		Name:          name,
		ContactPerson: contactPerson,
		Address:       address,
		Phone:         textFromPtr(phone),
		VatNumber:     textFromPtr(vatNumber),
		RegNumber:     textFromPtr(regNumber),
	})
}

func (r *CompanyRepository) Delete(ctx context.Context, id int64, familyID string) error {
	hasInvoices, err := r.q.CompanyHasNonDraftInvoices(ctx, id)
	if err != nil {
		return fmt.Errorf("check invoices: %w", err)
	}
	if hasInvoices {
		return fmt.Errorf("company has non-draft invoices and cannot be deleted")
	}
	return r.q.DeleteCompany(ctx, sqlc.DeleteCompanyParams{ID: id, FamilyID: familyID})
}

func textFromPtr(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}

func ptrFromText(t pgtype.Text) *string {
	if !t.Valid {
		return nil
	}
	return &t.String
}
