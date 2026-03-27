-- name: GetClient :one
SELECT * FROM clients
WHERE id = $1 AND family_id = $2 AND deleted_at IS NULL;

-- name: ListClients :many
SELECT * FROM clients
WHERE family_id = $1 AND deleted_at IS NULL
ORDER BY name;

-- name: ListClientsByStatus :many
SELECT * FROM clients
WHERE family_id = $1 AND status = $2 AND deleted_at IS NULL
ORDER BY name;

-- name: CreateClient :one
INSERT INTO clients (family_id, name, contact_person, email, address, vat_number, reg_number, contract_reference, contract_notes, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateClient :one
UPDATE clients
SET name = $3,
    contact_person = $4,
    email = $5,
    address = $6,
    vat_number = $7,
    reg_number = $8,
    contract_reference = $9,
    contract_notes = $10,
    status = $11
WHERE id = $1 AND family_id = $2 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteClient :exec
UPDATE clients
SET deleted_at = NOW()
WHERE id = $1 AND family_id = $2 AND deleted_at IS NULL;

-- name: ClientHasNonDraftInvoices :one
SELECT EXISTS (
    SELECT 1 FROM invoices
    WHERE client_id = $1 AND status != 'draft' AND deleted_at IS NULL
) AS has_invoices;
