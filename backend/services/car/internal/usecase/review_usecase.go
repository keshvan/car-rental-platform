package usecase

import (
	"context"
	"fmt"

	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/repo"
)

type reviewUsecase struct {
	reviewRepo repo.ReviewRepository
}

func NewReviewUsecase(reviewRepo repo.ReviewRepository) ReviewUsecase {
	return &reviewUsecase{reviewRepo: reviewRepo}
}

func (u *reviewUsecase) GetByCarID(ctx context.Context, carID int64) ([]entity.Review, error) {
	reviews, err := u.reviewRepo.FindByCarID(ctx, carID)
	if err != nil {
		return nil, fmt.Errorf("ReviewUsecase - GetByCarID - ReviewRepository.FindByCarID: %w", err)
	}
	return reviews, nil
}

func (u *reviewUsecase) GetUnverified(ctx context.Context) ([]entity.Review, error) {
	reviews, err := u.reviewRepo.FindUnverified(ctx)
	if err != nil {
		return nil, fmt.Errorf("ReviewUsecase - GetUnverified - ReviewRepository.FindUnverified: %w", err)
	}
	return reviews, nil
}

func (u *reviewUsecase) VerifyReview(ctx context.Context, id int64) error {
	err := u.reviewRepo.VerifyReview(ctx, id)
	if err != nil {
		return fmt.Errorf("ReviewUsecase - VerifyReview - ReviewRepository.VerifyReview: %w", err)
	}
	return nil
}

func (u *reviewUsecase) Delete(ctx context.Context, id int64) error {
	err := u.reviewRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("ReviewUsecase - Delete - ReviewRepository.Delete: %w", err)
	}
	return nil
}
