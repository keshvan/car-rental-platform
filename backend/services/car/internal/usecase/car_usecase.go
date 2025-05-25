package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/repo"
)

var (
	ErrCarNotFound   = errors.New("car not found")
	ErrBrandNotFound = errors.New("brand not found")
)

type carUsecase struct {
	carRepo repo.CarRepository
}

func NewCarUsecase(carRepo repo.CarRepository) CarUsecase {
	return &carUsecase{carRepo}
}

func (u *carUsecase) NewCar(ctx context.Context, car *entity.Car) error {
	brand, err := u.carRepo.BrandByID(ctx, car.BrandID)
	if err != nil {
		return fmt.Errorf("CarUsecase - NewCar - CarRepository.BrandByID: %w", err)
	}
	if brand == nil {
		return ErrBrandNotFound
	}

	car.Name = fmt.Sprintf("%s %s", brand.Name, car.Model)

	if err := u.carRepo.Create(ctx, car); err != nil {
		return fmt.Errorf("CarUsecase - NewCar - CarRepository.Create: %w", err)
	}
	return nil
}

func (u *carUsecase) Update(ctx context.Context, id int64, car *entity.Car) error {
	brand, err := u.carRepo.BrandByID(ctx, car.BrandID)
	if err != nil {
		return fmt.Errorf("CarUsecase - NewCar - CarRepository.BrandByID: %w", err)
	}
	if brand == nil {
		return ErrBrandNotFound
	}

	car.Name = fmt.Sprintf("%s %s", brand.Name, car.Model)

	if err := u.carRepo.Update(ctx, id, car); err != nil {
		return fmt.Errorf("CarUsecase - UpdateCar - CarRepository.Update: %w", err)
	}
	return nil
}

func (u *carUsecase) GetAll(ctx context.Context) ([]entity.Car, error) {
	cars, err := u.carRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("CarUsecase - GetBrands - CarRepository.FindAll: %w", err)
	}

	return cars, nil
}

func (u *carUsecase) Delete(ctx context.Context, id int64) error {
	if err := u.carRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("CarUsecase - Delete - CarRepository.Delete: %w", err)
	}
	return nil
}

func (u *carUsecase) GetBrands(ctx context.Context) ([]entity.Brand, error) {
	brands, err := u.carRepo.AllBrands(ctx)
	if err != nil {
		return nil, fmt.Errorf("CarUsecase - GetBrands - CarRepository.AllBrands: %w", err)
	}
	return brands, nil
}

func (u *carUsecase) GetByID(ctx context.Context, id int64) (*entity.Car, error) {
	car, err := u.carRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, repo.ErrCarNotFound) {
			return nil, ErrCarNotFound
		}
		return nil, fmt.Errorf("CarUsecase - GetByID - CarRepository.FindByID: %w", err)
	}
	return car, nil
}
