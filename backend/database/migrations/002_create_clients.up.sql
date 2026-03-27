CREATE TABLE clients (
    id                 BIGSERIAL PRIMARY KEY,
    family_id          INTEGER NOT NULL,
    name               VARCHAR(255) NOT NULL,
    contact_person     VARCHAR(255),
    email              VARCHAR(255),
    address            TEXT NOT NULL,
    vat_number         VARCHAR(50),
    reg_number         VARCHAR(50),
    contract_reference VARCHAR(255),
    contract_notes     TEXT,
    status             VARCHAR(20) NOT NULL DEFAULT 'active'
                       CHECK (status IN ('active', 'inactive')),
    created_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at         TIMESTAMPTZ
);
