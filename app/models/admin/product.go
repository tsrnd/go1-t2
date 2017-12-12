package admin

import (
	"goweb2/app/models"
	"time"
)

/**
 * Product struct
 */
type Product struct {
	Id          int
	Name        string
	Description string
	Image       string
	Price       float64
	Created_at  time.Time
	Updated_at  time.Time
}

/**
 * Get 10 products to show on admin home page
 */
func GetProductLimit() ([]*Product, error) {
	var products []*Product
	db := models.DB
	rows, err := db.Query("SELECT id, name, description, image, price, created_at, updated_at FROM products")
	if err != nil {

		return nil, err
	}
	for rows.Next() {
		product := &Product{}
		rows.Scan(&product.Id, &product.Name, &product.Description, &product.Image, &product.Price, &product.Created_at, &product.Updated_at)
		products = append(products, product)
	}

	return products, nil
}
