CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE eduflex_disbursements (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    school_id       UUID NOT NULL,
    invoice_id      UUID NOT NULL,
    amount          NUMERIC NOT NULL,
    disbursed_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    payment_method  TEXT NOT NULL,
    reference       TEXT NOT NULL,
    status          TEXT NOT NULL, -- PENDING, COMPLETED, FAILED
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
