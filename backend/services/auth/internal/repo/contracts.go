package repo

import (
	"context"

	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/entity"
)

type (
	UserRepository interface {
		Create(context.Context, *entity.User) (int64, error)
		Delete(context.Context, int64) error
		FindByEmail(context.Context, string) (*entity.User, error)
		FindByID(context.Context, int64) (*entity.User, error)
		GetRole(context.Context, int64) (string, error)
	}

	TokenRepository interface {
		Save(context.Context, string, int64) error
		Delete(context.Context, string) error
		GetUserID(context.Context, string) (int64, error)
	}
)
