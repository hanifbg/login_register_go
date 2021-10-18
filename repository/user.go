package repository

import (
	user "github.com/hanifbg/login_register/entity/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user.User) (user.User, error)
	FindByEmail(login user.LoginUser) user.User
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user user.User) (user.User, error) {

	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) FindByEmail(login user.LoginUser) user.User {
	var result user.User
	r.db.Where("email = ?", login.Email).First(&result)

	return result
}
