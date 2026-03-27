CREATE TABLE bank_accounts (
    id              BIGSERIAL PRIMARY KEY,
    company_id      BIGINT NOT NULL REFERENCES companies(id),
    bank_name       VARCHAR(255) NOT NULL,
    bank_address    TEXT NOT NULL,
    account_holder  VARCHAR(255) NOT NULL,
    iban            VARCHAR(50) NOT NULL
                    CHECK (LENGTH(iban) >= 15 AND LENGTH(iban) <= 34),
    swift           VARCHAR(20) NOT NULL
                    CHECK (LENGTH(swift) IN (8, 11)),
    currency        VARCHAR(3) NOT NULL,
    is_default      BOOLEAN NOT NULL DEFAULT false,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);
