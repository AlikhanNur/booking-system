-- name: CreateUser :one
INSERT INTO users (name, email, phone_number, role, created_at, updated_at)
VALUES ($1, $2, $3, $4, now(), now())
RETURNING id, name, email, phone_number, role, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, name, email, phone_number, role, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetAllUsers :many
SELECT id, name, email, phone_number, role, created_at, updated_at
FROM users
ORDER BY id;

-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3, phone_number = $4, role = $5, updated_at = now()
WHERE id = $1
RETURNING id, name, email, phone_number, role, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
