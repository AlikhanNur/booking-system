package repository

import (
	"booking-system/internal/database/db"
	"booking-system/internal/models"
	"context"
)

type RestaurantRepository interface {
	Create(ctx context.Context, restaurant *models.Restaurant) (*models.Restaurant, error)
	GetByID(ctx context.Context, id int64) (*models.Restaurant, error)
	GetAll(ctx context.Context) ([]*models.Restaurant, error)
	Update(ctx context.Context, restaurant *models.Restaurant) (*models.Restaurant, error)
	Delete(ctx context.Context, id int64) error
}

type restaurantRepo struct {
	queries *db.Queries
}

func NewRestaurantRepo(queries *db.Queries) RestaurantRepository {
	return &restaurantRepo{
		queries: queries,
	}
}

func (r *restaurantRepo) Create(ctx context.Context, restaurant *models.Restaurant) (*models.Restaurant, error) {

	res, err := r.queries.CreateRestaurant(ctx, db.CreateRestaurantParams{
		Name:        restaurant.Name,
		Address:     toNullString(restaurant.Address),
		PhoneNumber: toNullString(restaurant.PhoneNumber),
		Status:      restaurant.Status,
	})

	if err != nil {
		return nil, err
	}

	return mapDBRestaurantToModel(&res), nil

}

func (r *restaurantRepo) GetByID(ctx context.Context, id int64) (*models.Restaurant, error) {
	res, err := r.queries.GetRestaurnatByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return mapDBRestaurantToModel(&res), nil
}

func (r *restaurantRepo) GetAll(ctx context.Context) ([]*models.Restaurant, error) {
	restaurants, err := r.queries.GetAllRestaurnats(ctx)

	if err != nil {
		return nil, err
	}

	result := make([]*models.Restaurant, 0, len(restaurants))

	for _, restaurant := range restaurants {
		result = append(result, mapDBRestaurantToModel(&restaurant))
	}

	return result, nil
}

func (r *restaurantRepo) Update(ctx context.Context, restaurant *models.Restaurant) (*models.Restaurant, error) {
	res, err := r.queries.UpdateRestaurnat(ctx, db.UpdateRestaurnatParams{
		ID:          restaurant.ID,
		Name:        restaurant.Name,
		Address:     toNullString(restaurant.Address),
		PhoneNumber: toNullString(restaurant.PhoneNumber),
		Status:      restaurant.Status,
	})

	if err != nil {
		return nil, err
	}

	return mapDBRestaurantToModel(&res), nil
}

func (r *restaurantRepo) Delete(ctx context.Context, id int64) error {
	return r.queries.DeleteRestaurnat(ctx, id)
}
