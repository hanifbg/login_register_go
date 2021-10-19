package repository

import (
	user "github.com/hanifbg/login_register/entity/user"
)

type UserRepository interface {
	Create(user.User) (user.User, error)
	FindByEmail(login user.LoginUser) user.User
}

type repository struct {
	opt Option
}

func NewUserRepository(opt Option) *repository {
	return &repository{
		opt: opt,
	}
}

func (r *repository) Create(user user.User) (user.User, error) {

	err := r.opt.DB.Create(&user).Error

	return user, err
}

func (r *repository) FindByEmail(login user.LoginUser) user.User {
	var result user.User
	r.opt.DB.Where("email = ?", login.Email).First(&result)

	return result
}
