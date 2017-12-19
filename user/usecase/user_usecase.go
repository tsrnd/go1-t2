package usecase

import (
	model "goweb2/user"
	userRepo "goweb2/user/repository"
)

// UserUsecase interface
type UserUsecase interface {
	GetByID(id int64) (*model.User, error)
}

// userUsecase struct
type userUsecase struct {
	userRepo userRepo.UserRepository
}

// GetByID func
func (us *userUsecase) GetByID(id int64) (*model.User, error) {
	return us.userRepo.GetByID(id)
}

// NewUserUsecase func
func NewUserUsecase(us userRepo.UserRepository) UserUsecase {
	return &userUsecase{us}
}
