package models

import "time"

// table - столики в ресторанах и кафе
type Table struct {
	ID           int64     `json:"id"`
	RestaurantID int64     `json:"restaurant_id"`
	TableNumber  string    `json:"table_number"`
	Seats        int64     `json:"seats"`
	Status       bool      `json:"status"` // занят - не занят
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
