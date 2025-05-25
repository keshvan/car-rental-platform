package entity

import "time"

type User struct {
	ID           int64     `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	Role         string    `json:"role" db:"role"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Balance      int64     `json:"balance" db:"balance"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
