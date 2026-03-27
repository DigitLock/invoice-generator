-- name: ListInvoiceItems :many
SELECT * FROM invoice_items
WHERE invoice_id = $1
ORDER BY id;

-- name: CreateInvoiceItem :one
INSERT INTO invoice_items (invoice_id, description, quantity, unit_price, total)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteInvoiceItemsByInvoice :exec
DELETE FROM invoice_items
WHERE invoice_id = $1;
