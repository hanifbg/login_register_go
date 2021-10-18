package service

import "github.com/hanifbg/login_register/repository"

type Option struct {
	*repository.Repository
}

//all service object
type Services struct {
	UserService UserServiceInterface
}
