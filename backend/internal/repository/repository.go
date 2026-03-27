package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DigitLock/invoice-generator/backend/internal/database/sqlc"
)

type Repositories struct {
	Companies    *CompanyRepository
	Clients      *ClientRepository
	BankAccounts *BankAccountRepository
	Invoices     *InvoiceRepository
	pool         *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repositories {
	queries := sqlc.New(pool)
	return &Repositories{
		Companies:    &CompanyRepository{q: queries},
		Clients:      &ClientRepository{q: queries},
		BankAccounts: &BankAccountRepository{q: queries, pool: pool},
		Invoices:     &InvoiceRepository{q: queries, pool: pool},
		pool:         pool,
	}
}
