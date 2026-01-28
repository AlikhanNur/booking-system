package models

import "time"

type Restaurant struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Status      bool      `json:"status"` // открыто или закрыто
	Address     string    `json:"adsress"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
