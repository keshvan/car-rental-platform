package repo

import (
	"context"
	"time"

	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
)

type (
	CarRepository interface {
		Create(ctx context.Context, car *entity.Car) error
		FindAll(ctx context.Context) ([]entity.Car, error)
		FindByID(ctx context.Context, id int64) (*entity.Car, error)
		Update(ctx context.Context, id int64, car *entity.Car) error
		SetAvailability(ctx context.Context, carID int64, available bool) error
		Delete(ctx context.Context, id int64) error
		AllBrands(ctx context.Context) ([]entity.Brand, error)
		NewBrand(ctx context.Context, brandName string) error
		BrandByID(ctx context.Context, id int64) (*entity.Brand, error)
		DeleteBrand(ctx context.Context, brandId int64) error
	}

	RentRepository interface {
		Create(context.Context, *entity.Rent) error
		FindByID(ctx context.Context, id int64) (*entity.Rent, error)
		FindByUserID(ctx context.Context, user_id int64) ([]entity.Rent, error)
		FindAll(ctx context.Context) ([]entity.Rent, error)
		CompleteRent(ctx context.Context, id int64, endDate time.Time, totalPrice int64) error
		Cancel(ctx context.Context, id int64, endDate time.Time) error
		Delete(ctx context.Context, id int64) error
	}

	ReviewRepository interface {
		Create(ctx context.Context, review *entity.Review) error
		FindByCarID(ctx context.Context, carID int64) ([]entity.Review, error)
		FindUnverified(ctx context.Context) ([]entity.Review, error)
		VerifyReview(ctx context.Context, id int64) error
		Delete(ctx context.Context, id int64) error
	}

	BalanceRepository interface {
		UpdateBalance(context.Context, int64, int64) error
		GetBalance(context.Context, int64) (int64, error)
	}
)
