package usecase

import (
	"context"

	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/controller/response"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/entity"
)

type AuthUsecase interface {
	Register(ctx context.Context, email string, password string) (*response.RegisterResponse, error)
	Login(ctx context.Context, email, password, refreshToken string, cookieName string) (*response.LoginResponse, error)
	Refresh(ctx context.Context, refreshToken string) (*response.Tokens, error)
	Logout(ctx context.Context, refreshToken string) error
	CheckSession(ctx context.Context, refreshToken string) (*response.CheckSessionResponse, error)
}

type UserUsecase interface {
	GetAllUsers(ctx context.Context) ([]entity.User, error)
	UpdateUserRole(ctx context.Context, id int64, role string) error
	DeleteUser(ctx context.Context, id int64) error
}
