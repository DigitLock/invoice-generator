package repository

import (
	"context"
	"fmt"

	"github.com/DigitLock/invoice-generator/backend/internal/database/sqlc"
)

type ClientRepository struct {
	q *sqlc.Queries
}

func (r *ClientRepository) GetByID(ctx context.Context, id int64, familyID int32) (sqlc.Client, error) {
	return r.q.GetClient(ctx, sqlc.GetClientParams{ID: id, FamilyID: familyID})
}

func (r *ClientRepository) List(ctx context.Context, familyID int32) ([]sqlc.Client, error) {
	return r.q.ListClients(ctx, familyID)
}

func (r *ClientRepository) ListByStatus(ctx context.Context, familyID int32, status string) ([]sqlc.Client, error) {
	return r.q.ListClientsByStatus(ctx, sqlc.ListClientsByStatusParams{FamilyID: familyID, Status: status})
}

func (r *ClientRepository) Create(ctx context.Context, familyID int32, params sqlc.CreateClientParams) (sqlc.Client, error) {
	params.FamilyID = familyID
	if params.Status == "" {
		params.Status = "active"
	}
	return r.q.CreateClient(ctx, params)
}

func (r *ClientRepository) Update(ctx context.Context, id int64, familyID int32, params sqlc.UpdateClientParams) (sqlc.Client, error) {
	params.ID = id
	params.FamilyID = familyID
	return r.q.UpdateClient(ctx, params)
}

func (r *ClientRepository) Delete(ctx context.Context, id int64, familyID int32) error {
	hasInvoices, err := r.q.ClientHasNonDraftInvoices(ctx, id)
	if err != nil {
		return fmt.Errorf("check invoices: %w", err)
	}
	if hasInvoices {
		return fmt.Errorf("client has non-draft invoices and cannot be deleted")
	}
	return r.q.DeleteClient(ctx, sqlc.DeleteClientParams{ID: id, FamilyID: familyID})
}
