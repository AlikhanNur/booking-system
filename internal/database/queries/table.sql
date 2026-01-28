-- name: CreateTable :one
INSERT INTO tables (restaurant_id, table_number, seats, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, now(), now())
RETURNING id, restaurant_id, table_number, seats, status, created_at, updated_at;

-- name: GetTableByID :one
SELECT * FROM tables 
WHERE id = $1;

-- name: GetAllTables :many
SELECT * FROM tables
ORDER BY id;

-- name: UpdateTable :one
UPDATE tables
SET restaurant_id = $2,
    table_number = $3,
    seats = $4,
    status = $5,
    updated_at = now()
Where id = $1
RETURNING id, restaurant_id, table_number, seats, status, created_at, updated_at;

-- name: DeleteTable :exec
DELETE FROM tables
WHERE id = $1;
