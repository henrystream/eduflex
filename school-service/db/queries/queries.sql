-- name: CreateSchool :one
INSERT INTO schools (id, name, address, contact_email, contact_phone)
VALUES (gen_random_uuid(), $1, $2, $3, $4)
RETURNING *;

-- name: GetSchool :one
SELECT * FROM schools WHERE id = $1;

-- name: ListSchools :many
SELECT * FROM schools ORDER BY created_at;

-- name: UpdateSchool :one
UPDATE schools
SET name = $2, address = $3, contact_email = $4, contact_phone = $5
WHERE id = $1
RETURNING *;

-- name: DeleteSchool :exec
DELETE FROM schools WHERE id = $1;