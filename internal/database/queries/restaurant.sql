-- name: CreateRestaurant :one
INSERT INTO restaurants (name, address, phone_number, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, now(), now())
RETURNING id, name, address, phone_number, status, created_at, updated_at;

-- name: GetRestaurnatByID :one
SELECT * FROM restaurants
WHERE id = $1;

-- name: GetAllRestaurnats :many
SELECT * FROM restaurants
ORDER BY id;

-- name: UpdateRestaurnat :one
UPDATE restaurants
SET name = $2,
    address = $3,
    phone_number = $4,
    status = $5,
    updated_at = now()
Where id = $1
RETURNING id, name, address, phone_number, status, created_at, updated_at;

-- name: DeleteRestaurnat :exec
DELETE FROM restaurants
WHERE id = $1;
