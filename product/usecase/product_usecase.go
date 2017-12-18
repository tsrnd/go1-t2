package usecase

import (
	model "goweb2/product"
	repos "goweb2/product/repository"
)

// ProductUsecase interface
type ProductUsecase interface {
	GetLimit() ([]*model.Product, error)
}

// productUsecase struct
type productUsecase struct {
	productRepos repos.ProductRepository
}

// GetLimit func
func (us *productUsecase) GetLimit() ([]*model.Product, error) {
	return us.productRepos.GetLimit()
}

// NewProductUsecase func
func NewProductUsecase(us repos.ProductRepository) ProductUsecase {
	return &productUsecase{us}
}
