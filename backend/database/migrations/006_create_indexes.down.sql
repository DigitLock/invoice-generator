DROP TRIGGER IF EXISTS trg_invoice_items_updated_at ON invoice_items;
DROP TRIGGER IF EXISTS trg_invoices_updated_at ON invoices;
DROP TRIGGER IF EXISTS trg_bank_accounts_updated_at ON bank_accounts;
DROP TRIGGER IF EXISTS trg_clients_updated_at ON clients;
DROP TRIGGER IF EXISTS trg_companies_updated_at ON companies;
DROP FUNCTION IF EXISTS update_updated_at();

DROP INDEX IF EXISTS idx_invoice_items_invoice_id;
DROP INDEX IF EXISTS idx_invoices_number_unique;
DROP INDEX IF EXISTS idx_invoices_overdue;
DROP INDEX IF EXISTS idx_invoices_status;
DROP INDEX IF EXISTS idx_invoices_company_date;
DROP INDEX IF EXISTS idx_invoices_family_id;
DROP INDEX IF EXISTS idx_bank_accounts_default_unique;
DROP INDEX IF EXISTS idx_bank_accounts_company_id;
DROP INDEX IF EXISTS idx_clients_status;
DROP INDEX IF EXISTS idx_clients_family_id;
DROP INDEX IF EXISTS idx_companies_family_id;
