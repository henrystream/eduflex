CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE ledger_entries (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_type      TEXT NOT NULL,   -- FINANCING_AGREEMENT, INSTALLMENT, DRAWDOWN, REPAYMENT, STUDENT_PAYMENT
    event_id        UUID NOT NULL,   -- id from source service
    source_service  TEXT NOT NULL,   -- financing-service, loan-service, student-service
    debit_account   TEXT NOT NULL,
    credit_account  TEXT NOT NULL,
    amount          NUMERIC NOT NULL,
    currency        TEXT NOT NULL DEFAULT 'AED',
    occurred_at     TIMESTAMPTZ NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
