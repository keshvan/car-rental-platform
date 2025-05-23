package repo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type tokenRepo struct {
	db *sqlx.DB
}

func NewTokenRepo(db *sqlx.DB) TokenRepository {
	return &tokenRepo{db: db}
}

func (r *tokenRepo) Save(ctx context.Context, token string, userID int64) error {
	if _, err := r.db.ExecContext(ctx, "INSERT INTO tokens (token, user_id) VALUES($1, $2)", token, userID); err != nil {
		return fmt.Errorf("TokenRepository - Save - pg.Pool.Exec(): %w", err)
	}
	return nil
}

func (r *tokenRepo) Delete(ctx context.Context, token string) error {
	if _, err := r.db.ExecContext(ctx, "DELETE FROM tokens WHERE token = $1", token); err != nil {
		return fmt.Errorf("TokenRepository - Delete - pg.Pool.Exec(): %w", err)
	}
	return nil
}

func (r *tokenRepo) GetUserID(ctx context.Context, token string) (int64, error) {
	row := r.db.QueryRowContext(ctx, "SELECT user_id FROM tokens WHERE token = $1", token)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("TokenRepository - GetUserID - row.Scan(): %w", err)
	}

	return id, nil
}
