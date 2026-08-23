CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE financing_agreements (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    student_id     UUID NOT NULL,
    invoice_id     UUID NOT NULL,
    principal      NUMERIC NOT NULL,
    interest_rate  NUMERIC NOT NULL,
    service_fee    NUMERIC NOT NULL,
    total_payable  NUMERIC NOT NULL,
    term_months    INT NOT NULL,
    start_date     DATE NOT NULL,
    status         TEXT NOT NULL,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


CREATE TABLE monthly_installments (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    financing_id      UUID NOT NULL REFERENCES financing_agreements(id),
    installment_number INT NOT NULL,
    due_date          DATE NOT NULL,
    amount            NUMERIC NOT NULL,
    status            TEXT NOT NULL,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

