CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE banks (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name          TEXT NOT NULL,
    contact_email TEXT,
    contact_phone TEXT,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE loan_facilities (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bank_id       UUID NOT NULL REFERENCES banks(id),
    credit_limit  NUMERIC NOT NULL,
    interest_rate NUMERIC NOT NULL,
    start_date    DATE NOT NULL,
    end_date      DATE NOT NULL,
    status        TEXT NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE loan_drawdowns (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    facility_id   UUID NOT NULL REFERENCES loan_facilities(id),
    amount        NUMERIC NOT NULL,
    drawdown_date DATE NOT NULL,
    reference     TEXT NOT NULL
);

CREATE TABLE loan_repayments (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    drawdown_id UUID NOT NULL REFERENCES loan_drawdowns(id),
    amount      NUMERIC NOT NULL,
    paid_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    reference   TEXT NOT NULL
);
