package usecase

import (
	"context"
	"fmt"

	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/repo"
)

type balanceUsecase struct {
	balanceRepo repo.BalanceRepository
}

func NewBalanceUsecase(balanceRepo repo.BalanceRepository) BalanceUsecase {
	return &balanceUsecase{balanceRepo: balanceRepo}
}

func (u *balanceUsecase) UpdateBalance(ctx context.Context, balance *entity.Balance) error {
	if err := u.balanceRepo.UpdateBalance(ctx, balance.UserID, balance.Amount); err != nil {
		return fmt.Errorf("RentUsecase - UpdateBalance - u.rentRepo.UpdateBalance(): %w", err)
	}

	return nil
}
