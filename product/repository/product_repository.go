package repository

import (
	"database/sql"

	model "goweb2/product"
)

// ProductRepository interface
type ProductRepository interface {
	GetLimit() ([]*model.Product, error)
}

// productRepository struct
type productRepository struct {
	DB *sql.DB
}

// GetLimit func
func (m *productRepository) GetLimit() ([]*model.Product, error) {
	const query = `SELECT id, name, description, image, price, category_id, created_at, updated_at FROM products LIMIT $1`
	products := make([]*model.Product, 0)
	var limit int
	limit = 10
	rows, err := m.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product model.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Image, &product.Price, &product.Category_id, &product.Created_at, &product.Updated_at )
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, err
}

// NewProductRepository func
func NewProductRepository(DB *sql.DB) ProductRepository {
	return &productRepository{DB}
}
