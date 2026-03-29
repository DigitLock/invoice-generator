ALTER TABLE invoices ALTER COLUMN due_date SET NOT NULL;
ALTER TABLE invoices ADD CONSTRAINT invoices_due_date_check CHECK (due_date >= issue_date);
