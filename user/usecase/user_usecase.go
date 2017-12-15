package usecase

import (
	model "goweb2/user"
	repos "goweb2/user/repository"
)

// UserUsecase interface
type UserUsecase interface {
	GetByID(id int64) (*model.User, error)
}

// userUsecase struct
type userUsecase struct {
	userRepos repos.UserRepository
}

// GetByID func
func (us *userUsecase) GetByID(id int64) (*model.User, error) {
	return us.userRepos.GetByID(id)
}

// NewUserUsecase func
func NewUserUsecase(us repos.UserRepository) UserUsecase {
	return &userUsecase{us}
}
