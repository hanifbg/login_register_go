package repository

import "gorm.io/gorm"

//object that repo needed
type Option struct {
	DB *gorm.DB
}

//all repository object
type Repository struct {
	User UserRepository
}
