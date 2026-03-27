-- Companies
CREATE INDEX idx_companies_family_id ON companies(family_id) WHERE deleted_at IS NULL;

-- Clients
CREATE INDEX idx_clients_family_id ON clients(family_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_clients_status ON clients(family_id, status) WHERE deleted_at IS NULL;

-- Bank accounts
CREATE INDEX idx_bank_accounts_company_id ON bank_accounts(company_id) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_bank_accounts_default_unique ON bank_accounts(company_id)
    WHERE is_default = true AND deleted_at IS NULL;

-- Invoices
CREATE INDEX idx_invoices_family_id ON invoices(family_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_invoices_company_date ON invoices(company_id, issue_date) WHERE deleted_at IS NULL;
CREATE INDEX idx_invoices_status ON invoices(family_id, status) WHERE deleted_at IS NULL;
CREATE INDEX idx_invoices_overdue ON invoices(family_id, is_overdue)
    WHERE deleted_at IS NULL AND is_overdue = true;
CREATE UNIQUE INDEX idx_invoices_number_unique ON invoices(company_id, invoice_number)
    WHERE deleted_at IS NULL;

-- Invoice items
CREATE INDEX idx_invoice_items_invoice_id ON invoice_items(invoice_id);

-- Updated_at triggers
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_companies_updated_at BEFORE UPDATE ON companies
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_clients_updated_at BEFORE UPDATE ON clients
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_bank_accounts_updated_at BEFORE UPDATE ON bank_accounts
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_invoices_updated_at BEFORE UPDATE ON invoices
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_invoice_items_updated_at BEFORE UPDATE ON invoice_items
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();
