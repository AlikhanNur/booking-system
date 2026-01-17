package models

// table - столики в ресторанах и кафе
type Table struct {
	ID           int64  `json:"id"`
	RestaurantID int64  `json:"restaurant_id"`
	TableNumber  int64  `json:"table_number"`
	Seats        int64  `json:"seats"`
	Status       string `json:"status"` // занят - не занят
}
