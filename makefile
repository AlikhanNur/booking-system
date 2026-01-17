
migrate_up:
	migrate -path internal/database/migrations -database "postgres://user:password@localhost:5432/booking_system?sslmode=disable" up

migrate_down:
	migrate -path internal/database/migrations -database "postgres://user:password@localhost:5432/booking_system?sslmode=disable" down