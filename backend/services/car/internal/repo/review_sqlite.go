package repo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
)

type reviewRepo struct {
	db *sqlx.DB
}

func NewReviewRepository(db *sqlx.DB) ReviewRepository {
	return &reviewRepo{db}
}

func (r *reviewRepo) Create(ctx context.Context, review *entity.Review) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO reviews (rent_id, user_id, car_id, rating, comment) VALUES ($1, $2, $3, $4, $5)", review.RentID, review.UserID, review.CarID, review.Rating, review.Comment)
	if err != nil {
		return fmt.Errorf("ReviewRepository - Create - r.db.ExecContext: %w", err)
	}

	return nil
}

func (r *reviewRepo) FindByCarID(ctx context.Context, carID int64) ([]entity.Review, error) {
	var reviews []entity.Review
	err := r.db.SelectContext(ctx, &reviews, "SELECT r.*, u.email FROM reviews r JOIN users u ON r.user_id = u.id WHERE r.car_id = $1 AND r.verified = true", carID)
	if err != nil {
		return nil, fmt.Errorf("ReviewRepository - FindByCarID - r.db.SelectContext: %w", err)
	}

	return reviews, nil
}

func (r *reviewRepo) FindUnverified(ctx context.Context) ([]entity.Review, error) {
	var reviews []entity.Review
	err := r.db.SelectContext(ctx, &reviews, "SELECT r.*, u.email FROM reviews r JOIN users u ON r.user_id = u.id WHERE r.verified = false")
	if err != nil {
		return nil, fmt.Errorf("ReviewRepository - FindUnverified - r.db.SelectContext: %w", err)
	}

	return reviews, nil
}

func (r *reviewRepo) VerifyReview(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, "UPDATE reviews SET verified = true WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("ReviewRepository - VerifyReview - r.db.ExecContext: %w", err)
	}

	return nil
}

func (r *reviewRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM reviews WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("ReviewRepository - Delete - r.db.ExecContext: %w", err)
	}
	return nil
}
