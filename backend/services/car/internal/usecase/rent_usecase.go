package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/repo"
)

const (
	STATUS_COMPLETE = "complete"
)

var (
	ErrNoAccess = errors.New("no access")
)

type rentUsecase struct {
	rentRepo    repo.RentRepository
	carRepo     repo.CarRepository
	balanceRepo repo.BalanceRepository
	reviewRepo  repo.ReviewRepository
}

func NewRentUsecase(rentRepo repo.RentRepository, carRepo repo.CarRepository, balanceRepo repo.BalanceRepository, reviewRepo repo.ReviewRepository) RentUsecase {
	return &rentUsecase{rentRepo: rentRepo, carRepo: carRepo, balanceRepo: balanceRepo, reviewRepo: reviewRepo}
}

var (
	ErrNotEnoughBalance = errors.New("not enough balance")
	ErrCarNotAvailable  = errors.New("car is not available")
)

func (u *rentUsecase) NewRent(ctx context.Context, rent *entity.Rent) error {
	balance, err := u.balanceRepo.GetBalance(ctx, rent.UserID)
	if err != nil {
		return fmt.Errorf("RentUsecase - NewRent - u.BalanceRepository.GetBalance(): %w", err)
	}

	if balance <= 0 {
		return ErrNotEnoughBalance
	}

	car, err := u.carRepo.FindByID(ctx, rent.CarID)
	if err != nil {
		return fmt.Errorf("RentUsecase - NewRent - u.CarRepository.FindByID(): %w", err)
	}

	if !car.Available {
		return ErrCarNotAvailable
	}

	if err := u.rentRepo.Create(ctx, rent); err != nil {
		fmt.Println(err)
		return fmt.Errorf("RentUsecase - NewRent - u.RentRepository.Create(): %w", err)
	}

	if err := u.carRepo.SetAvailability(ctx, rent.CarID, false); err != nil {
		return fmt.Errorf("RentUsecase - NewRent - u.CarRepository.SetAvailability(): %w", err)
	}

	return nil
}

func (u *rentUsecase) GetAllRents(ctx context.Context) ([]entity.Rent, error) {
	rents, err := u.rentRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("RentUsecase - GetAllRents - u.RentRepository.FindAll(): %w", err)
	}

	return rents, nil
}

func (u *rentUsecase) GetRentByID(ctx context.Context, id int64) (*entity.Rent, error) {
	rent, err := u.rentRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("RentUsecase - GetRentByID - u.RentRepository.FindByID(): %w", err)
	}

	return rent, nil
}

func (u *rentUsecase) GetRentsByUserID(ctx context.Context, userId int64) ([]entity.Rent, error) {
	rents, err := u.rentRepo.FindByUserID(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("RentUsecase - GetRentsByUserID - u.RentRepository.FindByUserID(): %w", err)
	}

	return rents, nil
}

func (u *rentUsecase) CompleteRent(ctx context.Context, id int64, review *entity.Review, userId int64, role string) error {
	rent, err := u.rentRepo.FindByID(ctx, id)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return repo.ErrRentNotFound
		}
		return fmt.Errorf("RentUsecase - CompleteRent - RentRepository.FindByID: %w", err)
	}

	if role != "admin" || rent.UserID != userId {
		return ErrNoAccess
	}

	rent.EndDate = time.Now()
	total := int64(math.Round(float64(rent.EndDate.Sub(rent.StartDate).Hours()))) * rent.PricePerHour
	rent.TotalPrice = &total

	if err := u.balanceRepo.UpdateBalance(ctx, rent.UserID, -*rent.TotalPrice); err != nil {
		return fmt.Errorf("RentUsecase - CompleteRent - BalanceRepository.UpdateBalance: %w", err)
	}

	review.RentID = rent.ID
	review.UserID = rent.UserID
	review.CarID = rent.CarID

	if err := u.reviewRepo.Create(ctx, review); err != nil {
		return fmt.Errorf("RentUsecase - CompleteRent - ReviewRepository.Create: %w", err)
	}

	if err := u.rentRepo.CompleteRent(ctx, id, rent.EndDate, *rent.TotalPrice); err != nil {
		if errors.Is(err, repo.ErrRentNotFound) {
			return err
		}
		return fmt.Errorf("RentUsecase - CompleteRent - RentRepository.UpdateStatus: %w", err)
	}

	if err := u.carRepo.SetAvailability(ctx, rent.CarID, true); err != nil {
		return fmt.Errorf("RentUsecase - CompleteRent - CarRepository.SetAvailability: %w", err)
	}

	return nil
}

func (u *rentUsecase) Cancel(ctx context.Context, id int64) error {
	rent, err := u.rentRepo.FindByID(ctx, id)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return repo.ErrRentNotFound
		}
		return fmt.Errorf("RentUsecase - CompleteRent - RentRepository.FindByID: %w", err)
	}

	rent.EndDate = time.Now()
	total := int64(0)
	rent.TotalPrice = &total

	if err := u.rentRepo.Cancel(ctx, id, rent.EndDate); err != nil {
		if errors.Is(err, repo.ErrRentNotFound) {
			return err
		}
		return fmt.Errorf("RentUsecase - CompleteRent - RentRepository.UpdateStatus: %w", err)
	}

	if err := u.carRepo.SetAvailability(ctx, rent.CarID, true); err != nil {
		return fmt.Errorf("RentUsecase - CompleteRent - CarRepository.SetAvailability: %w", err)
	}

	return nil
}

func (u *rentUsecase) Delete(ctx context.Context, id int64) error {
	if err := u.rentRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("RentUsecase - DeleteRent - RentRepository.Delete: %w", err)
	}

	return nil
}
