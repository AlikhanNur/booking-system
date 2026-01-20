-- name: CreateBooking :one
INSERT INTO bookings (restaurant_id, table_id, user_id, date, time_from, time_to, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, now(), now())
RETURNING id, restaurant_id, table_id, user_id, date, time_from, time_to, status, created_at, updated_at;

-- name: GetBookingByID :one
SELECT id, restaurant_id, table_id, user_id, date, time_from, time_to, status, created_at, updated_at
FROM bookings
WHERE id = $1;

-- name: GetAllBookings :many
SELECT id, restaurant_id, table_id, user_id, date, time_from, time_to, status, created_at, updated_at
FROM bookings
ORDER BY date, time_from;

-- name: UpdateBooking :one
UPDATE bookings
SET restaurant_id = $2,
    table_id = $3,
    user_id = $4,
    date = $5,
    time_from = $6,
    time_to = $7,
    status = $8,
    updated_at = now()
WHERE id = $1
RETURNING id, restaurant_id, table_id, user_id, date, time_from, time_to, status, created_at, updated_at;

-- name: DeleteBooking :exec
DELETE FROM bookings
WHERE id = $1;
