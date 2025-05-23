package entity

import "time"

type Review struct {
	ID        int64     `json:"id"`
	RentID    int64     `json:"rent_id"`
	UserID    int64     `json:"user_id"`
	CarID     int64     `json:"car_id"`
	Rating    int64     `json:"rating"`
	Comment   string    `json:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
