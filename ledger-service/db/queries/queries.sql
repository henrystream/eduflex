-- name: CreateLedgerEntry :one
INSERT INTO ledger_entries (
    event_type, event_id, source_service,
    debit_account, credit_account,
    amount, currency, occurred_at
)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: ListLedgerEntriesByEvent :many
SELECT * FROM ledger_entries
WHERE event_type = $1 AND event_id = $2
ORDER BY created_at DESC;

-- name: ListLedgerEntriesByService :many
SELECT * FROM ledger_entries
WHERE source_service = $1
ORDER BY created_at DESC;

-- name: ListLedgerEntries :many
SELECT * FROM ledger_entries
ORDER BY occurred_at DESC;
