package entity

import "time"

type Car struct {
	ID           int64     `json:"id" db:"id"`
	BrandID      int64     `json:"brand_id" db:"brand_id"`
	BrandName    string    `json:"brand_name" db:"brand_name"`
	Model        string    `json:"model" db:"model"`
	Name         string    `json:"name" db:"name"`
	Year         int64     `json:"year" db:"year"`
	PricePerHour int64     `json:"price_per_hour" db:"price_per_hour"`
	ImageURL     string    `json:"image_url,omitempty" db:"image_url"`
	Available    bool      `json:"available" db:"available"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
