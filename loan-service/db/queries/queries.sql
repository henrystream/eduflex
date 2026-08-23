-- BANKS
-- name: CreateBank :one
INSERT INTO banks (name, contact_email, contact_phone)
VALUES ($1,$2,$3)
RETURNING *;

-- name: ListBanks :many
SELECT * FROM banks ORDER BY created_at DESC;

-- FACILITIES
-- name: CreateFacility :one
INSERT INTO loan_facilities (bank_id, credit_limit, interest_rate, start_date, end_date, status)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: ListFacilitiesByBank :many
SELECT * FROM loan_facilities WHERE bank_id = $1 ORDER BY created_at DESC;

-- DRAWDOWNS
-- name: CreateDrawdown :one
INSERT INTO loan_drawdowns (facility_id, amount, drawdown_date, reference)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: ListDrawdownsByFacility :many
SELECT * FROM loan_drawdowns WHERE facility_id = $1 ORDER BY drawdown_date DESC;

-- REPAYMENTS
-- name: CreateRepayment :one
INSERT INTO loan_repayments (drawdown_id, amount, paid_at, reference)
VALUES ($1,$2,NOW(),$3)
RETURNING *;

-- name: ListRepaymentsByDrawdown :many
SELECT * FROM loan_repayments WHERE drawdown_id = $1 ORDER BY paid_at DESC;
