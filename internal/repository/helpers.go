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
