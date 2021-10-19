package service

import (
	"errors"

	user "github.com/hanifbg/login_register/entity/user"
	"github.com/hanifbg/login_register/utils"
)

type UserServiceInterface interface {
	BindUser(user.User) user.User
	LoginUser(user.LoginUser) (token string, err error)
}

type userService struct {
	opt Option
}

func NewUserService(opt Option) UserServiceInterface {
	return &userService{
		opt: opt,
	}
}

func (us *userService) BindUser(u user.User) (user user.User) {
	hashedPassword, err := utils.EncryptPassword(u.Password)
	if err != nil {
		return
	}

	user = u
	user.Password = string(hashedPassword)
	us.opt.Repository.User.Create(user)

	return
}

func (us *userService) LoginUser(l user.LoginUser) (token string, err error) {
	result := us.opt.Repository.User.FindByEmail(l)

	if utils.ComparePassword(result.Password, l.Password) { //not finish yet
		return "JWT TOKEN", nil
	}

	return "", errors.New("wrong credentials")
}
