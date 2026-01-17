package models

type Restaurant struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"` // открыто или закрыто
	Address     string `json:"adsress"`
	PhoneNumber int64  `json:"phone_number"`
}
