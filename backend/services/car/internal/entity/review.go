package entity

import "time"

type Review struct {
	ID        int64     `json:"id" db:"id"`
	RentID    int64     `json:"rent_id" db:"rent_id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Email     string    `json:"email" db:"email"`
	CarID     int64     `json:"car_id" db:"car_id"`
	Rating    int64     `json:"rating" db:"rating"`
	Comment   string    `json:"comment,omitempty" db:"comment"`
	Verified  bool      `json:"-" db:"verified"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
