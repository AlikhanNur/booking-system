CREATE TABLE bookings (
    id BIGSERIAL PRIMARY KEY,
    restaurant_id BIGINT NOT NULL,
    table_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,

    date DATE NOT NULL,
    time_from TIMESTAMP NOT NULL,
    time_to TIMESTAMP NOT NULL,

    status TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT fk_booking_user
        FOREIGN KEY (user_id) REFERENCES users(id),

    CONSTRAINT fk_booking_restaurant
        FOREIGN KEY (restaurant_id) REFERENCES restaurants(id),

    CONSTRAINT fk_booking_table
        FOREIGN KEY (table_id) REFERENCES tables(id)
);
