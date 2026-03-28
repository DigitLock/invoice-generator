-- Revert user_id and family_id from TEXT back to INTEGER
ALTER TABLE invoices ALTER COLUMN family_id TYPE INTEGER USING family_id::INTEGER;
ALTER TABLE invoices ALTER COLUMN user_id TYPE INTEGER USING user_id::INTEGER;
ALTER TABLE clients ALTER COLUMN family_id TYPE INTEGER USING family_id::INTEGER;
ALTER TABLE companies ALTER COLUMN family_id TYPE INTEGER USING family_id::INTEGER;
