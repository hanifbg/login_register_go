package repository

import (
	user "github.com/hanifbg/login_register/entity/user"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user user.User) (user.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user user.User) (user.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
