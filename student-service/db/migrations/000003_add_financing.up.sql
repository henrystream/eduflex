CREATE TABLE financing_agreements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    student_id UUID NOT NULL REFERENCES students(id),
    invoice_id UUID NOT NULL,
    principal NUMERIC(12,2) NOT NULL,
    interest_rate NUMERIC(5,2) NOT NULL, -- e.g. 10.00
    service_fee NUMERIC(12,2) NOT NULL DEFAULT 0, -- fee charged to student
    total_payable NUMERIC(12,2) NOT NULL,
    term_months INT NOT NULL,
    start_date DATE NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('active','completed','defaulted')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_financing_agreements_student_id ON financing_agreements (student_id);
CREATE INDEX idx_financing_agreements_invoice_id ON financing_agreements (invoice_id);

CREATE TABLE monthly_installments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    financing_id UUID NOT NULL REFERENCES financing_agreements(id),
    installment_number INT NOT NULL,
    due_date DATE NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('pending','paid','overdue')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (financing_id, installment_number)
);

CREATE INDEX idx_monthly_installments_financing_id ON monthly_installments (financing_id);

CREATE TABLE student_payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    installment_id UUID NOT NULL REFERENCES monthly_installments(id),
    amount NUMERIC(12,2) NOT NULL,
    paid_at TIMESTAMP NOT NULL DEFAULT NOW(),
    payment_method TEXT,
    transaction_reference TEXT
);

CREATE INDEX idx_student_payments_installment_id ON student_payments (installment_id);