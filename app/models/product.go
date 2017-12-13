package models

import (
	"log"
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
	Created_at  string
	Updated_at  string
}

/**
 * Get 10 products to show on home page
 */
func GetProductLimit() ([]*Product, error) {
	products := make([]*Product, 0)
	rows, _ := db.Query("select id, name, description, image, price, created_at, updated_at FROM products ORDER BY id DESC LIMIT 10")
	for rows.Next() {
		product := new(Product)
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Image, &product.Price, &product.Created_at, &product.Updated_at)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	return products, nil
}

func ShowProduct(id string) (Product, error) {
	var product Product
	err := db.QueryRow("select id, name, description, image, price, created_at, updated_at FROM products WHERE id = $1", id).Scan(&product.Id, &product.Name, &product.Description, &product.Image, &product.Price, &product.Created_at, &product.Updated_at)
	return product, err

}
