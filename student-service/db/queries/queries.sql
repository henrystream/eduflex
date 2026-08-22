-- name: CreateStudent :one
INSERT INTO students (first_name, last_name, date_of_birth, email, phone)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetStudent :one
SELECT * FROM students WHERE id = $1;

-- name: ListStudents :many
SELECT * FROM students ORDER BY created_at DESC;
