CREATE TABLE tables (
    id BIGSERIAL PRIMARY KEY,
    restaurant_id BIGINT NOT NULL,
    table_number VARCHAR(50) NOT NULL,
    seats INTEGER NOT NULL CHECK (seats > 0),
    status BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT fk_tables_restaurant
        FOREIGN KEY (restaurant_id)
        REFERENCES restaurants(id)
        ON DELETE CASCADE
);