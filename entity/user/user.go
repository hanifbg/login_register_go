package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name"  validate:"required"`
	Email        string `json:"email" validate:"required,email" gorm:"type:varchar(20)"`
	Phone_number string `json:"phone_number" validate:"required,number"`
	Password     string `json:"password"  validate:"required"`
	Address      string `json:"address"  validate:"required"`
	Role         int
	TokenHash    string
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
