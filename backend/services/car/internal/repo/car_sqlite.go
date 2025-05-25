package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
)

type carRepo struct {
	db *sqlx.DB
}

var (
	ErrCarNotFound = errors.New("car not found")
)

func NewCarRepository(db *sqlx.DB) CarRepository {
	return &carRepo{db}
}

func (r *carRepo) Create(ctx context.Context, car *entity.Car) error {
	if _, err := r.db.ExecContext(ctx, "INSERT INTO cars (brand_id, model, name, year, price_per_hour, image_url) VALUES ($1, $2, $3, $4, $5, $6)", car.BrandID, car.Model, car.Name, car.Year, car.PricePerHour, car.ImageURL); err != nil {
		return fmt.Errorf("CarRepository - Create - db.ExecContext: %w", err)
	}
	return nil
}

func (r *carRepo) FindAll(ctx context.Context) ([]entity.Car, error) {
	rows, err := r.db.QueryxContext(ctx, `
	SELECT
    	c.*,
    	c.brand_id,
    	b.name AS brand_name
	FROM
    	cars c
	INNER JOIN
    	brands b ON c.brand_id = b.id`)
	if err != nil {
		return nil, fmt.Errorf("CarRepository - FindAll - db.QueryxContext: %w", err)
	}
	defer rows.Close()

	var cars []entity.Car
	var car entity.Car

	for rows.Next() {
		err := rows.StructScan(&car)
		if err != nil {
			return nil, fmt.Errorf("CarRepository - FindAll - rows.Next - rows.StructScan: %w", err)
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func (r *carRepo) FindByID(ctx context.Context, id int64) (*entity.Car, error) {
	row := r.db.QueryRowxContext(ctx, `
	SELECT
    	c.*,
    	c.brand_id,
    	b.name AS brand_name
	FROM
    	cars c
	INNER JOIN
    	brands b ON c.brand_id = b.id
	WHERE c.id = $1`, id)

	var car entity.Car
	err := row.StructScan(&car)

	switch err {
	case nil:
		return &car, nil
	case sql.ErrNoRows:
		return nil, ErrCarNotFound
	default:
		return nil, fmt.Errorf("CarRepository - FindByID - row.StructScan: %w", err)
	}
}

func (r *carRepo) Update(ctx context.Context, id int64, car *entity.Car) error {
	_, err := r.db.ExecContext(ctx, `
	UPDATE cars
	SET
		brand_id = COALESCE($1, brand_id),
		model = COALESCE($2, model),
		name = COALESCE($3, name),
		year = COALESCE($4, year),
		price_per_hour = COALESCE($5, price_per_hour),
		image_url = COALESCE($6, image_url),
		updated_at = CURRENT_TIMESTAMP
	WHERE id = $7
	`, car.BrandID, car.Model, car.Name, car.Year, car.PricePerHour, car.ImageURL, id)

	if err != nil {
		return fmt.Errorf("CarRepository - Update - row.ExecContext: %w", err)
	}
	return nil
}

func (r *carRepo) SetAvailability(ctx context.Context, carID int64, available bool) error {
	result, err := r.db.ExecContext(ctx, `
    UPDATE cars
    SET
		available = $1,
        updated_at = CURRENT_TIMESTAMP
    WHERE id = $2;
	`, available, carID)
	if err != nil {
		return fmt.Errorf("CarRepository - SetAvailability - db.ExecContext: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return ErrCarNotFound
	}

	return nil
}

func (r *carRepo) Delete(ctx context.Context, id int64) error {
	if _, err := r.db.ExecContext(ctx, "DELETE FROM cars WHERE id = $1", id); err != nil {
		return fmt.Errorf("CarRepository - Delete - db.ExecContext: %w", err)
	}
	return nil
}

func (r *carRepo) AllBrands(ctx context.Context) ([]entity.Brand, error) {
	rows, err := r.db.QueryxContext(ctx, "SELECT * FROM brands")
	if err != nil {
		return nil, fmt.Errorf("CarRepository - GetAllBrands - db.QueryxContext: %w", err)
	}

	var brands []entity.Brand
	var brand entity.Brand

	for rows.Next() {
		err := rows.StructScan(&brand)
		if err != nil {
			return nil, fmt.Errorf("CarRepository - GetAllBrands - rows.Next - rows.StructScan: %w", err)
		}
		brands = append(brands, brand)
	}
	return brands, nil
}

func (r *carRepo) BrandByID(ctx context.Context, id int64) (*entity.Brand, error) {
	var brand entity.Brand
	err := r.db.GetContext(ctx, &brand, "SELECT * FROM brands WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("CarRepository - BrandByID - db.GetContext: %w", err)
	}
	return &brand, nil
}

func (r *carRepo) NewBrand(ctx context.Context, brandName string) error {
	if _, err := r.db.ExecContext(ctx, "INSERT INTO brands (name) VALUES($1)", brandName); err != nil {
		return fmt.Errorf("CarRepository - NewBrand - db.ExecContext: %w", err)
	}
	return nil
}

func (r *carRepo) DeleteBrand(ctx context.Context, brandId int64) error {
	if _, err := r.db.ExecContext(ctx, "DELETE FROM brands WHERE id = $1", brandId); err != nil {
		return fmt.Errorf("CarRepository - DeleteBrand - db.ExecContext: %w", err)
	}
	return nil
}
