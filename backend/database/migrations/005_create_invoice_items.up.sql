CREATE TABLE invoice_items (
    id          BIGSERIAL PRIMARY KEY,
    invoice_id  BIGINT NOT NULL REFERENCES invoices(id) ON DELETE CASCADE,
    description TEXT NOT NULL,
    quantity    DECIMAL(10,2) NOT NULL CHECK (quantity > 0),
    unit_price  DECIMAL(15,2) NOT NULL CHECK (unit_price >= 0),
    total       DECIMAL(15,2) NOT NULL CHECK (total >= 0),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
