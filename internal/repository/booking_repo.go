package repository

import (
	"booking-system/internal/database/db"
	"booking-system/internal/models"
	"context"
)

type BookingRepository interface {
	Create(ctx context.Context, booking *models.Booking) (*models.Booking, error)
	GetByID(ctx context.Context, id int64) (*models.Booking, error)
	GetAll(ctx context.Context) ([]*models.Booking, error)
	Update(ctx context.Context, booking *models.Booking) (*models.Booking, error)
	Delete(ctx context.Context, id int64) error
}

type bookingRepo struct {
	queries *db.Queries
}

func NewBookingRepository(queries *db.Queries) BookingRepository {
	return &bookingRepo{
		queries: queries,
	}
}

func (r *bookingRepo) Create(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	b, err := r.queries.CreateBooking(ctx, db.CreateBookingParams{
		RestaurantID: booking.RestaurantID,
		TableID:      booking.TableID,
		UserID:       booking.UserID,
		Date:         booking.Date,
		TimeFrom:     booking.TimeFrom,
		TimeTo:       booking.TimeTo,
		Status:       booking.Status,
	})
	if err != nil {
		return nil, err
	}
	return &models.Booking{
		ID:           b.ID,
		RestaurantID: b.RestaurantID,
		TableID:      b.TableID,
		UserID:       b.UserID,
		Date:         b.Date,
		TimeFrom:     b.TimeFrom,
		TimeTo:       b.TimeTo,
		Status:       b.Status,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}, nil
}

func (r *bookingRepo) GetByID(ctx context.Context, id int64) (*models.Booking, error) {
	b, err := r.queries.GetBookingByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.Booking{
		ID:           b.ID,
		RestaurantID: b.RestaurantID,
		TableID:      b.TableID,
		UserID:       b.UserID,
		Date:         b.Date,
		TimeFrom:     b.TimeFrom,
		TimeTo:       b.TimeTo,
		Status:       b.Status,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}, nil
}

func (r *bookingRepo) GetAll(ctx context.Context) ([]*models.Booking, error) {
	bookings, err := r.queries.GetAllBookings(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Booking, len(bookings))
	for i, b := range bookings {
		result[i] = &models.Booking{
			ID:           b.ID,
			RestaurantID: b.RestaurantID,
			TableID:      b.TableID,
			UserID:       b.UserID,
			Date:         b.Date,
			TimeFrom:     b.TimeFrom,
			TimeTo:       b.TimeTo,
			Status:       b.Status,
			CreatedAt:    b.CreatedAt,
			UpdatedAt:    b.UpdatedAt,
		}
	}
	return result, nil
}

func (r *bookingRepo) Update(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	b, err := r.queries.UpdateBooking(ctx, db.UpdateBookingParams{
		ID:           booking.ID,
		RestaurantID: booking.RestaurantID,
		TableID:      booking.TableID,
		UserID:       booking.UserID,
		Date:         booking.Date,
		TimeFrom:     booking.TimeFrom,
		TimeTo:       booking.TimeTo,
		Status:       booking.Status,
	})
	if err != nil {
		return nil, err
	}
	return &models.Booking{
		ID:           b.ID,
		RestaurantID: b.RestaurantID,
		TableID:      b.TableID,
		UserID:       b.UserID,
		Date:         b.Date,
		TimeFrom:     b.TimeFrom,
		TimeTo:       b.TimeTo,
		Status:       b.Status,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}, nil
}

func (r *bookingRepo) Delete(ctx context.Context, id int64) error {
	return r.queries.DeleteBooking(ctx, id)
}
