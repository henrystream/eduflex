-- name: CreateAgreement :one
INSERT INTO financing_agreements (
    student_id, invoice_id, principal, interest_rate, service_fee,
    total_payable, term_months, start_date, status
)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
RETURNING *;

-- name: GetAgreement :one
SELECT * FROM financing_agreements WHERE id = $1;

-- name: ListAgreementsByStudent :many
SELECT * FROM financing_agreements WHERE student_id = $1 ORDER BY created_at DESC;



-- name: CreateInstallment :one
INSERT INTO monthly_installments (
    financing_id, installment_number, due_date, amount, status
)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: ListInstallments :many
SELECT * FROM monthly_installments WHERE financing_id = $1 ORDER BY installment_number;

-- name: UpdateAgreementStatus :one
UPDATE financing_agreements
SET status = $2
WHERE id = $1
RETURNING *;

-- name: CreatePayment :one
INSERT INTO student_payments (installment_id, amount, payment_method, transaction_reference)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListPaymentsByStudent :many
SELECT p.*
FROM student_payments p
JOIN monthly_installments mi ON mi.id = p.installment_id
JOIN financing_agreements fa ON fa.id = mi.financing_id
WHERE fa.student_id = $1
ORDER BY p.paid_at DESC;
