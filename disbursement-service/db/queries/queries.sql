-- name: CreateDisbursement :one
INSERT INTO eduflex_disbursements (
    school_id, invoice_id, amount, payment_method, reference, status
)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetDisbursement :one
SELECT * FROM eduflex_disbursements WHERE id = $1;

-- name: ListDisbursementsBySchool :many
SELECT * FROM eduflex_disbursements WHERE school_id = $1 ORDER BY disbursed_at DESC;

