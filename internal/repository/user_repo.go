package repository

import (
	"context"

	"booking-system/internal/database/db"
	"booking-system/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, id int64) error
}

type userRepo struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) UserRepository {
	return &userRepo{queries: queries}
}

func (r *userRepo) Create(ctx context.Context, user *models.User) (*models.User, error) {
	u, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Name:        toNullString(user.Name),
		Email:       toNullString(user.Email),
		PhoneNumber: toNullString(user.PhoneNumber),
		Role:        toNullString(user.Role),
	})
	if err != nil {
		return nil, err
	}

	return mapDBUserToModel(&u), nil
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (*models.User, error) {
	u, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapDBUserToModel(&u), nil
}

func (r *userRepo) GetAll(ctx context.Context) ([]*models.User, error) {
	users, err := r.queries.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*models.User, 0, len(users))
	for _, u := range users {
		result = append(result, mapDBUserToModel(&u))
	}

	return result, nil
}

func (r *userRepo) Update(ctx context.Context, user *models.User) (*models.User, error) {
	u, err := r.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:          user.ID,
		Name:        toNullString(user.Name),
		Email:       toNullString(user.Email),
		PhoneNumber: toNullString(user.PhoneNumber),
		Role:        toNullString(user.Role),
	})
	if err != nil {
		return nil, err
	}

	return mapDBUserToModel(&u), nil
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	return r.queries.DeleteUser(ctx, id)
}
