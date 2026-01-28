package repository

import (
	"booking-system/internal/database/db"
	"booking-system/internal/models"
	"database/sql"
)

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func fromNullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func mapDBUserToModel(u *db.User) *models.User {
	return &models.User{
		ID:          u.ID,
		Name:        fromNullString(u.Name),
		Email:       fromNullString(u.Email),
		PhoneNumber: fromNullString(u.PhoneNumber),
		Role:        fromNullString(u.Role),
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func mapDBRestaurantToModel(res *db.Restaurant) *models.Restaurant {
	return &models.Restaurant{
		ID:          res.ID,
		Name:        res.Name,
		Status:      res.Status,
		Address:     fromNullString(res.Address),
		PhoneNumber: fromNullString(res.PhoneNumber),
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}
}

func mapDBTableToModel(table *db.Table) *models.Table {
	return &models.Table{
		ID:           table.ID,
		RestaurantID: table.RestaurantID,
		TableNumber:  table.TableNumber,
		Seats:        int64(table.Seats),
		Status:       table.Status,
		CreatedAt:    table.CreatedAt,
		UpdatedAt:    table.UpdatedAt,
	}
}
