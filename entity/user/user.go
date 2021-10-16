package entity

import "time"

type User struct {
	ID           int
	Name         string `json:"name"  validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Phone_number string `json:"phone_number" validate:"required,number"`
	Password     string `json:"password"  validate:"required"`
	Address      string `json:"address"  validate:"required"`
	Role         int
	TokenHash    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
