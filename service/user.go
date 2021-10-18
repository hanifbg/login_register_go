package service

import (
	"time"

	user "github.com/hanifbg/login_register/entity/user"
	"github.com/hanifbg/login_register/utils"
)

type UserServiceInterface interface {
	BindUser(user.User) user.User
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

	user.Name = u.Name
	user.Email = u.Email
	user.Phone_number = u.Phone_number
	user.Address = u.Address
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	us.opt.Repository.User.Create(user)

	return
}
