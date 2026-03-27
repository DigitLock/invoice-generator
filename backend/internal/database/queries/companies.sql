-- name: GetCompany :one
SELECT * FROM companies
WHERE id = $1 AND family_id = $2 AND deleted_at IS NULL;

-- name: ListCompanies :many
SELECT * FROM companies
WHERE family_id = $1 AND deleted_at IS NULL
ORDER BY name;

-- name: CreateCompany :one
INSERT INTO companies (family_id, name, contact_person, address, phone, vat_number, reg_number)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateCompany :one
UPDATE companies
SET name = $3,
    contact_person = $4,
    address = $5,
    phone = $6,
    vat_number = $7,
    reg_number = $8
WHERE id = $1 AND family_id = $2 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteCompany :exec
UPDATE companies
SET deleted_at = NOW()
WHERE id = $1 AND family_id = $2 AND deleted_at IS NULL;

-- name: CompanyHasNonDraftInvoices :one
SELECT EXISTS (
    SELECT 1 FROM invoices
    WHERE company_id = $1 AND status != 'draft' AND deleted_at IS NULL
) AS has_invoices;
