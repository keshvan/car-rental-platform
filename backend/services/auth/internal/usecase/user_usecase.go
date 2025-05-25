package usecase

import (
	"context"

	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/entity"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/repo"
)

type userUsecase struct {
	userRepo repo.UserRepository
}

func NewUserUsecase(userRepo repo.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return u.userRepo.FindAll(ctx)
}

func (u *userUsecase) UpdateUserRole(ctx context.Context, id int64, role string) error {
	return u.userRepo.UpdateRole(ctx, id, role)
}

func (u *userUsecase) DeleteUser(ctx context.Context, id int64) error {
	return u.userRepo.Delete(ctx, id)
}
