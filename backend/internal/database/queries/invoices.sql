-- name: GetInvoice :one
SELECT * FROM invoices
WHERE id = $1 AND family_id = $2 AND deleted_at IS NULL;

-- name: ListInvoices :many
SELECT i.*,
       c.name AS company_name,
       cl.name AS client_name,
       (SELECT COUNT(*) FROM invoice_items WHERE invoice_id = i.id) AS items_count
FROM invoices i
JOIN companies c ON c.id = i.company_id
JOIN clients cl ON cl.id = i.client_id
WHERE i.family_id = $1 AND i.deleted_at IS NULL
ORDER BY i.issue_date DESC, i.created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountInvoices :one
SELECT COUNT(*) FROM invoices
WHERE family_id = $1 AND deleted_at IS NULL;

-- name: CreateInvoice :one
INSERT INTO invoices (
    user_id, family_id, company_id, client_id, bank_account_id,
    invoice_number, issue_date, due_date, currency, status, is_overdue,
    vat_rate, subtotal, vat_amount, total,
    contract_reference, external_reference, notes
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, 'draft', false, $10, $11, $12, $13, $14, $15, $16)
RETURNING *;

-- name: UpdateInvoice :one
UPDATE invoices
SET company_id = $3,
    client_id = $4,
    bank_account_id = $5,
    invoice_number = $6,
    issue_date = $7,
    due_date = $8,
    currency = $9,
    vat_rate = $10,
    subtotal = $11,
    vat_amount = $12,
    total = $13,
    contract_reference = $14,
    external_reference = $15,
    notes = $16
WHERE id = $1 AND family_id = $2 AND deleted_at IS NULL
RETURNING *;

-- name: UpdateInvoiceStatus :one
UPDATE invoices
SET status = $3
WHERE id = $1 AND family_id = $2 AND deleted_at IS NULL
RETURNING *;

-- name: UpdateInvoiceOverdue :one
UPDATE invoices
SET is_overdue = $3
WHERE id = $1 AND family_id = $2 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteInvoice :exec
UPDATE invoices
SET deleted_at = NOW()
WHERE id = $1 AND family_id = $2 AND deleted_at IS NULL;

-- name: GetMaxInvoiceSequence :one
SELECT COALESCE(MAX(
    CAST(SUBSTRING(invoice_number FROM '[0-9]+$') AS INTEGER)
), 0) AS max_seq
FROM invoices
WHERE user_id = $1
  AND invoice_number LIKE $2
  AND deleted_at IS NULL;
