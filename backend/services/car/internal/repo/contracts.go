package repo

import (
	"context"

	"github.com/keshvan/car-rental-platform/backend/services/car/internal/controller/request"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
)

type (
	CarRepository interface {
		Create(ctx context.Context, car *entity.Car) error
		FindAll(ctx context.Context) ([]entity.Car, error)
		FindByID(ctx context.Context, id int64) (*entity.Car, error)
		Update(ctx context.Context, id int64, req *request.UpdateCarRequest) error
		SetAvailability(ctx context.Context, carID int64, available bool) error
		Delete(ctx context.Context, id int64) error
		AllBrands(ctx context.Context) ([]entity.Brand, error)
		NewBrand(ctx context.Context, brandName string) error
		DeleteBrand(ctx context.Context, brandId int64) error
	}
	ReviewRepository interface{}
)
