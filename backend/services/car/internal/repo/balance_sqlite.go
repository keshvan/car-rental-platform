package repo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type balanceRepo struct {
	db *sqlx.DB
}

func NewBalanceRepository(db *sqlx.DB) BalanceRepository {
	return &balanceRepo{db: db}
}

func (r *balanceRepo) UpdateBalance(ctx context.Context, id int64, amount int64) error {
	if _, err := r.db.ExecContext(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", amount, id); err != nil {
		return fmt.Errorf("RentRepository - UpdateBalance - r.db.ExecContext(): %w", err)
	}

	return nil
}

func (r *balanceRepo) GetBalance(ctx context.Context, id int64) (int64, error) {
	var balance int64
	if err := r.db.GetContext(ctx, &balance, "SELECT balance FROM users WHERE id = $1", id); err != nil {
		return 0, fmt.Errorf("RentRepository - GetBalance - r.db.GetContext(): %w", err)
	}
	return balance, nil
}
