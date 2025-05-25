package usecase

import (
	"context"

	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
)

type (
	CarUsecase interface {
		GetAll(ctx context.Context) ([]entity.Car, error)
		GetBrands(ctx context.Context) ([]entity.Brand, error)
		GetByID(ctx context.Context, id int64) (*entity.Car, error)
		Delete(ctx context.Context, id int64) error
		NewCar(ctx context.Context, car *entity.Car) error
		Update(ctx context.Context, id int64, car *entity.Car) error
	}

	RentUsecase interface {
		NewRent(ctx context.Context, rent *entity.Rent) error
		GetAllRents(ctx context.Context) ([]entity.Rent, error)
		GetRentsByUserID(ctx context.Context, id int64) ([]entity.Rent, error)
		GetRentByID(ctx context.Context, id int64) (*entity.Rent, error)
		CompleteRent(ctx context.Context, id int64, review *entity.Review, userId int64, role string) error
		Cancel(ctx context.Context, id int64) error
		Delete(ctx context.Context, id int64) error
	}

	BalanceUsecase interface {
		UpdateBalance(ctx context.Context, balance *entity.Balance) error
	}

	ReviewUsecase interface {
		GetByCarID(ctx context.Context, carID int64) ([]entity.Review, error)
		GetUnverified(ctx context.Context) ([]entity.Review, error)
		VerifyReview(ctx context.Context, id int64) error
		Delete(ctx context.Context, id int64) error
	}
)
