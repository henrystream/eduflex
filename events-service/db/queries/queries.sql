-- name: CreateEvent :one
INSERT INTO domain_events (
    event_type, source_service, aggregate_id, payload, occurred_at
)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: ListUnprocessedEvents :many
SELECT * FROM domain_events
WHERE processed = FALSE
ORDER BY occurred_at ASC;

-- name: MarkEventProcessed :exec
UPDATE domain_events
SET processed = TRUE
WHERE id = $1;
