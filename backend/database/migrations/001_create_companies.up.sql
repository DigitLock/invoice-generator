CREATE TABLE companies (
    id          BIGSERIAL PRIMARY KEY,
    family_id   INTEGER NOT NULL,
    name        VARCHAR(255) NOT NULL,
    contact_person VARCHAR(255) NOT NULL,
    address     TEXT NOT NULL,
    phone       VARCHAR(50),
    vat_number  VARCHAR(50),
    reg_number  VARCHAR(50),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
