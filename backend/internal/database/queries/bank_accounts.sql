-- name: GetBankAccount :one
SELECT * FROM bank_accounts
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListBankAccountsByCompany :many
SELECT * FROM bank_accounts
WHERE company_id = $1 AND deleted_at IS NULL
ORDER BY is_default DESC, bank_name;

-- name: CreateBankAccount :one
INSERT INTO bank_accounts (company_id, bank_name, bank_address, account_holder, iban, swift, currency, is_default)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateBankAccount :one
UPDATE bank_accounts
SET bank_name = $2,
    bank_address = $3,
    account_holder = $4,
    iban = $5,
    swift = $6,
    currency = $7,
    is_default = $8
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteBankAccount :exec
UPDATE bank_accounts
SET deleted_at = NOW()
WHERE id = $1 AND deleted_at IS NULL;

-- name: ClearDefaultBankAccount :exec
UPDATE bank_accounts
SET is_default = false
WHERE company_id = $1 AND id != $2 AND deleted_at IS NULL;

-- name: BankAccountHasNonDraftInvoices :one
SELECT EXISTS (
    SELECT 1 FROM invoices
    WHERE bank_account_id = $1 AND status != 'draft' AND deleted_at IS NULL
) AS has_invoices;

-- name: GetBankAccountCompanyID :one
SELECT company_id FROM bank_accounts
WHERE id = $1 AND deleted_at IS NULL;
