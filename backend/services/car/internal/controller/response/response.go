package response

import "github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"

type CarWithReviews struct {
	Car     entity.Car      `json:"car"`
	Reviews []entity.Review `json:"reviews"`
}
