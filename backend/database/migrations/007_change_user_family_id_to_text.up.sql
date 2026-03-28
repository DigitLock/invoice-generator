-- Change user_id and family_id from INTEGER to TEXT to support UUID values from Expense Tracker JWT
ALTER TABLE companies ALTER COLUMN family_id TYPE TEXT USING family_id::TEXT;
ALTER TABLE clients ALTER COLUMN family_id TYPE TEXT USING family_id::TEXT;
ALTER TABLE invoices ALTER COLUMN user_id TYPE TEXT USING user_id::TEXT;
ALTER TABLE invoices ALTER COLUMN family_id TYPE TEXT USING family_id::TEXT;
