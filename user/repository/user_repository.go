package repository

import (
	model "goweb2/user"
)

// UserRepository interface
type UserRepository interface {
	GetByID(id int64) (*model.User, error)
}
