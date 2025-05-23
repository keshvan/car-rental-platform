package usecase

import (
	"context"

	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/controller/response"
)

type AuthUsecase interface {
	Register(ctx context.Context, email string, password string) (*response.RegisterResponse, error)
	Login(ctx context.Context, email, password, refreshToken string) (*response.LoginResponse, error)
	Refresh(ctx context.Context, refreshToken string) (*response.Tokens, error)
	Logout(ctx context.Context, refreshToken string) error
	CheckSession(ctx context.Context, refreshToken string) (*response.CheckSessionResponse, error)
}
