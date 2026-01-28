package repository

import (
	"booking-system/internal/database/db"
	"booking-system/internal/models"
	"context"
)

type TableRepository interface {
	Create(ctx context.Context, table *models.Table) (*models.Table, error)
	GetByID(ctx context.Context, id int64) (*models.Table, error)
	Update(ctx context.Context, table *models.Table) (*models.Table, error)
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]*models.Table, error)
}

type tableRepo struct {
	queries *db.Queries
}

func NewTableRepostory(queries *db.Queries) TableRepository {
	return &tableRepo{
		queries: queries,
	}
}

func (t *tableRepo) Create(ctx context.Context, table *models.Table) (*models.Table, error) {
	tab, err := t.queries.CreateTable(ctx, db.CreateTableParams{
		RestaurantID: table.RestaurantID,
		TableNumber:  table.TableNumber,
		Seats:        int32(table.Seats),
		Status:       table.Status,
	})

	if err != nil {
		return nil, err
	}

	return mapDBTableToModel(&tab), nil
}

func (t *tableRepo) GetByID(ctx context.Context, id int64) (*models.Table, error) {
	tab, err := t.queries.GetTableByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return mapDBTableToModel(&tab), nil
}

func (t *tableRepo) Update(ctx context.Context, table *models.Table) (*models.Table, error) {
	tab, err := t.queries.UpdateTable(ctx, db.UpdateTableParams{
		ID:           table.ID,
		RestaurantID: table.RestaurantID,
		TableNumber:  table.TableNumber,
		Seats:        int32(table.Seats),
		Status:       table.Status,
	})

	if err != nil {
		return nil, err
	}

	return mapDBTableToModel(&tab), nil
}

func (t *tableRepo) Delete(ctx context.Context, id int64) error {
	return t.queries.DeleteTable(ctx, id)
}

func (t *tableRepo) GetAll(ctx context.Context) ([]*models.Table, error) {
	tabs, err := t.queries.GetAllTables(ctx)

	if err != nil {
		return nil, err
	}

	result := make([]*models.Table, 0, len(tabs))

	for _, tab := range tabs {
		result = append(result, mapDBTableToModel(&tab))
	}

	return result, nil
}
