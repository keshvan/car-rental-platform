package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/keshvan/car-rental-platform/backend/pkg/jwt"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/controller/response"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/entity"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/repo"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepo  repo.UserRepository
	tokenRepo repo.TokenRepository
	jwt       *jwt.JWT
}

var (
	ErrInvalidEmailOrPassword        = errors.New("invalid email or password")
	ErrRefreshTokenNotFoundOrInvalid = errors.New("refresh token not found or invalid")
	ErrInvalidRefreshToken           = errors.New("invalid refresh token")
	ErrInvalidRole                   = errors.New("invalid role")
)

func NewAuthUsecase(userRepo repo.UserRepository, tokenRepo repo.TokenRepository, jwt *jwt.JWT) AuthUsecase {
	return &authUsecase{userRepo: userRepo, tokenRepo: tokenRepo, jwt: jwt}
}

func (u *authUsecase) Register(ctx context.Context, email string, password string) (*response.RegisterResponse, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &entity.User{Email: email, PasswordHash: string(hashedPassword)}

	id, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	access, err := u.jwt.GenerateAccessToken(id, user.Role)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - Register - GenerateAccessToken: %w", err)
	}

	refresh, err := u.jwt.GenerateRefreshToken(id)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - Register - GenerateRefreshToken: %w", err)
	}

	err = u.tokenRepo.Save(ctx, refresh, id)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - Register - TokenRepository.save(): %w", err)
	}

	createdUser, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - Register - UserRepository.FindByID(): %w", err)
	}

	return &response.RegisterResponse{User: *createdUser, Tokens: response.Tokens{AccessToken: access, RefreshToken: refresh}}, nil
}

func (u *authUsecase) Login(ctx context.Context, email, password, refreshToken, cookieName string) (*response.LoginResponse, error) {
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - Login - UserRepository.FindByEmail(): %w", err)
	}

	if cookieName == "refresh_token_admin" && user.Role != "admin" {
		return nil, ErrInvalidRole
	}

	if refreshToken != "" {
		if err := u.tokenRepo.Delete(ctx, refreshToken); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrInvalidEmailOrPassword
			}
			return nil, fmt.Errorf("AuthUsecase - Login - TokenRepository.Delete(): %w", err)
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		return nil, ErrInvalidEmailOrPassword
	}

	access, err := u.jwt.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - Login - jwt.GenerateAccessToken(): %w", err)
	}

	refresh, err := u.jwt.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - Login - jwt.GenerateRefreshToken(): %w", err)
	}

	err = u.tokenRepo.Save(ctx, refresh, user.ID)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - Login - TokenRepository.Save(): %w", err)
	}

	return &response.LoginResponse{User: *user, Tokens: response.Tokens{AccessToken: access, RefreshToken: refresh}}, nil
}

func (u *authUsecase) Refresh(ctx context.Context, refreshToken string) (*response.Tokens, error) {
	claims, err := u.jwt.ParseToken(refreshToken)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	var refreshClaims entity.RefreshClaims
	if err := mapstructure.Decode(claims, &refreshClaims); err != nil {
		return nil, fmt.Errorf("AuthUsecase - Refresh - mapstrucutre.Decode(): %w", err)
	}

	userID, err := u.tokenRepo.GetUserID(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRefreshTokenNotFoundOrInvalid
		}
		return nil, fmt.Errorf("AuthUsecase - Refresh - TokenRepository.GetUserID(): %w", err)
	}

	if err := u.tokenRepo.Delete(ctx, refreshToken); err != nil {
		return nil, fmt.Errorf("AuthUsecase - Refresh - TokenRepository.Delete(): %w", err)
	}

	role, err := u.userRepo.GetRole(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - Refresh - UserRepository.GetRole(): %w", err)
	}

	newAccess, err := u.jwt.GenerateAccessToken(userID, role)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - Refresh - jwt.GenerateAccessToken(): %w", err)
	}

	newRefresh, err := u.jwt.GenerateRefreshToken(userID)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - Refresh - jwt.GenerateRefreshToken(): %w", err)
	}

	err = u.tokenRepo.Save(ctx, newRefresh, userID)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - Refresh - TokenRepository.Save(): %w", err)
	}

	return &response.Tokens{AccessToken: newAccess, RefreshToken: newRefresh}, nil
}

func (u *authUsecase) Logout(ctx context.Context, refreshToken string) error {
	if err := u.tokenRepo.Delete(ctx, refreshToken); err != nil {
		return fmt.Errorf("AuthUsecase - Logout - TokenRepository.Delete(): %w", err)
	}
	return nil
}

func (u *authUsecase) CheckSession(ctx context.Context, refreshToken string) (*response.CheckSessionResponse, error) {
	if refreshToken == "" {
		return &response.CheckSessionResponse{User: nil, IsActive: false}, nil
	}
	id, err := u.tokenRepo.GetUserID(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &response.CheckSessionResponse{User: nil, IsActive: false}, nil
		}
		return nil, fmt.Errorf("AuthUsecase - CheckSession - TokenRepository.GetUserID(): %w", err)
	}

	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("AuthUsecase - CheckSession - UserRepository.FindByID(): %w", err)
	}
	return &response.CheckSessionResponse{User: user, IsActive: true}, nil
}
