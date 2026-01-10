package models

import "time"

// Booking — бронирование стола
type Booking struct {
	ID           int64 `json:"id"`
	RestaurantID int64 `json:"restaurant_id"`
	TableID      int64 `json:"table_id"`
	UserID       int64 `json:"user_id"`

	Date     time.Time `json:"date"`
	TimeFrom time.Time `json:"time_from"`
	TimeTo   time.Time `json:"time_to"`

	Status string `json:"status"` // pending / confirmed / cancelled

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
