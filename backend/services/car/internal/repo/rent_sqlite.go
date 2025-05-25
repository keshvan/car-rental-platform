package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
)

var (
	ErrRentNotFound = errors.New("rent not found")
)

type rentRepo struct {
	db *sqlx.DB
}

func NewRentRepository(db *sqlx.DB) RentRepository {
	return &rentRepo{db: db}
}

func (r *rentRepo) Create(ctx context.Context, rent *entity.Rent) error {
	if _, err := r.db.ExecContext(ctx, "INSERT INTO rents (user_id, car_id, start_date, end_date) VALUES ($1, $2, $3, $4)", rent.UserID, rent.CarID, rent.StartDate, rent.EndDate); err != nil {
		return fmt.Errorf("RentRepository - Create - r.db.ExecContext(): %w", err)
	}

	return nil
}

func (r *rentRepo) FindByID(ctx context.Context, id int64) (*entity.Rent, error) {
	var rent entity.Rent
	if err := r.db.GetContext(ctx, &rent, "SELECT r.*, c.name AS car_name, c.price_per_hour FROM rents r JOIN cars c ON r.car_id = c.id WHERE r.id = $1", id); err != nil {
		return nil, fmt.Errorf("RentRepository - FindByID - r.db.SelectContext(): %w", err)
	}
	return &rent, nil
}

func (r *rentRepo) FindByUserID(ctx context.Context, user_id int64) ([]entity.Rent, error) {
	var rents []entity.Rent
	if err := r.db.SelectContext(ctx, &rents, "SELECT r.*, c.name AS car_name, c.price_per_hour FROM rents r JOIN cars c ON r.car_id = c.id WHERE user_id = $1 ORDER BY r.status, r.start_date DESC", user_id); err != nil {
		return nil, fmt.Errorf("RentRepository - FindAll - r.db.SelectContext(): %w", err)
	}
	return rents, nil
}

func (r *rentRepo) FindAll(ctx context.Context) ([]entity.Rent, error) {
	var rents []entity.Rent
	if err := r.db.SelectContext(ctx, &rents, "SELECT r.*, c.name AS car_name, c.price_per_hour FROM rents r JOIN cars c ON r.car_id = c.id"); err != nil {
		return nil, fmt.Errorf("RentRepository - FindAll - r.db.SelectContext(): %w", err)
	}
	return rents, nil
}

func (r *rentRepo) CompleteRent(ctx context.Context, id int64, endDate time.Time, totalPrice int64) error {
	res, err := r.db.ExecContext(ctx, "UPDATE rents SET status = 'completed', end_date = $1, total_price = $2 WHERE id = $3", endDate, totalPrice, id)
	if err != nil {
		return fmt.Errorf("RentRepository - UpdateStatus - r.db.ExecContext: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("RentRepository - UpdateStatus - res.RowsAffected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrRentNotFound
	}

	return nil
}

func (r *rentRepo) Cancel(ctx context.Context, id int64, endDate time.Time) error {
	res, err := r.db.ExecContext(ctx, "UPDATE rents SET total_price = 0, end_date = $1, status = 'cancelled' WHERE id = $2", endDate, id)
	if err != nil {
		return fmt.Errorf("RentRepository - CancelRent - r.db.ExecContext: %w", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrRentNotFound
	}

	return nil
}

func (r *rentRepo) Delete(ctx context.Context, id int64) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM rents WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("RentRepository - DeleteRent - r.db.ExecContext: %w", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrRentNotFound
	}

	return nil
}
