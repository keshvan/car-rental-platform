package repo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/entity"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *entity.User) (int64, error) {
	row := r.db.QueryRowContext(ctx, "INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id", user.Email, user.PasswordHash)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("UserRepository - Create - row.Scan(): %w", err)
	}

	return id, nil
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	if _, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id); err != nil {
		return fmt.Errorf("UserRepository - Delete - db.Exec(): %w", err)
	}
	return nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, email, role, password_hash, created_at FROM users WHERE email = $1", email)

	var u entity.User
	if err := row.Scan(&u.ID, &u.Email, &u.Role, &u.PasswordHash, &u.CreatedAt); err != nil {
		return nil, fmt.Errorf("UserRepository - GetByEmail - row.Scan(): %w", err)
	}

	return &u, nil
}

func (r *userRepo) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, email, role, created_at FROM users WHERE id = $1", id)

	var u entity.User
	if err := row.Scan(&u.ID, &u.Email, &u.Role, &u.CreatedAt); err != nil {
		return nil, fmt.Errorf("UserRepository - GetByID - row.Scan(): %w", err)
	}

	return &u, nil
}

func (r *userRepo) GetRole(ctx context.Context, id int64) (string, error) {
	row := r.db.QueryRowContext(ctx, "SELECT role FROM users WHERE id = $1", id)

	var role string
	if err := row.Scan(&role); err != nil {
		return "", fmt.Errorf("UserRepository - GetRole - row.Scan(): %w", err)
	}

	return role, nil
}
