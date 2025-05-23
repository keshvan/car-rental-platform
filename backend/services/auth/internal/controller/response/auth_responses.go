package response

import "github.com/keshvan/car-rental-platform/backend/services/auth/internal/entity"

type (
	Tokens struct {
		AccessToken  string
		RefreshToken string
	}

	RegisterResponse struct {
		User   entity.User
		Tokens Tokens
	}

	LoginResponse struct {
		User   entity.User
		Tokens Tokens
	}

	CheckSessionResponse struct {
		User     *entity.User `json:"user"`
		IsActive bool         `json:"is_active"`
	}
)
