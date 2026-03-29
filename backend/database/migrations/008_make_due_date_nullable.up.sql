-- Make due_date optional
ALTER TABLE invoices DROP CONSTRAINT IF EXISTS invoices_due_date_check;
ALTER TABLE invoices ALTER COLUMN due_date DROP NOT NULL;
