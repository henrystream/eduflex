CREATE TABLE student_payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    installment_id UUID NOT NULL REFERENCES monthly_installments(id),
    amount NUMERIC(12,2) NOT NULL,
    paid_at TIMESTAMP NOT NULL DEFAULT NOW(),
    payment_method TEXT,
    transaction_reference TEXT
);