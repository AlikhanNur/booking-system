-- name: CreateBooking :one
INSERT INTO bookings (
    restaurant_id,
    table_id,
    user_id,
    date,
    time_from,
    time_to,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;
