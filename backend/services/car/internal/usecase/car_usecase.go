package usecase

import (
	"context"

	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/repo"
)

type CarUsecase struct {
	carRepo repo.CarRepository
}

func NewCarUsecase(carRepo repo.CarRepository) *CarUsecase {
	return &CarUsecase{carRepo}
}

func (u *CarUsecase) GetAll(ctx context.Context) ([]entity.Car, error) {
	cars, err := u.carRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return cars, nil
}
