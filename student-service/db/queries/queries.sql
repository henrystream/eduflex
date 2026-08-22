-- name: CreateStudent :one
INSERT INTO students (first_name, last_name, date_of_birth, email, phone)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetStudent :one
SELECT * FROM students WHERE id = $1;

-- name: ListStudents :many
SELECT * FROM students ORDER BY created_at DESC;

-- name: CreateEnrollment :one
INSERT INTO student_school_enrollments (student_id, school_id, enrollment_date, status)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetEnrollment :one
SELECT * FROM student_school_enrollments WHERE id = $1;

-- name: ListEnrollmentsByStudent :many
SELECT * FROM student_school_enrollments WHERE student_id = $1 ORDER BY created_at DESC;


-- name: CreatePayment :one
INSERT INTO student_payments (installment_id, amount, payment_method, transaction_reference)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListPaymentsByStudent :many
SELECT p.*
FROM student_payments p
JOIN monthly_installments mi ON mi.id = p.installment_id
WHERE mi.financing_id IN (
    SELECT id FROM financing_agreements WHERE student_id = $1
)
ORDER BY p.paid_at DESC;

