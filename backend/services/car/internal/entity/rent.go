package entity

import "time"

type Rent struct {
	ID           int64     `json:"id" db:"id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	CarID        int64     `json:"car_id" db:"car_id"`
	CarName      string    `json:"car_name" db:"car_name"`
	PricePerHour int64     `json:"price_per_hour" db:"price_per_hour"`
	StartDate    time.Time `json:"start_date" db:"start_date"`
	EndDate      time.Time `json:"end_date" db:"end_date"`
	TotalPrice   *int64    `json:"total_price" db:"total_price"`
	Status       string    `json:"status" db:"status"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
