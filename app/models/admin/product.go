package admin

import (
	"goweb2/app/models"
	"time"
	"fmt"
	"goweb2/helper"
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
	rows, err := models.Db.Query("SELECT id, name, description, image, price, created_at, updated_at FROM products")
	helper.Handle(err)
	for rows.Next() {
		product := &Product{}
		rows.Scan(&product.Id, &product.Name, &product.Description, &product.Image, &product.Price, &product.Created_at, &product.Updated_at)
		products = append(products, product)
	}
	return products, nil
}

func DeleteProductById(id string) {
	var products []*Product
	rows, err := models.Db.Query("SELECT id, name, description, image, price, created_at, updated_at FROM products WHERE id=$1", id)
	helper.Handle(err)
	for rows.Next() {
		product := &Product{}
		rows.Scan(&product.Id, &product.Name, &product.Description, &product.Image, &product.Price, &product.Created_at, &product.Updated_at)
		products = append(products, product)
		fmt.Println(product)
	}
	value2, err2 := models.Db.Exec("DELETE FROM cart_details WHERE product_id=$1", id)
	helper.Handle(err2)
	fmt.Println(value2)
	value1, err1 := models.Db.Exec("DELETE FROM products WHERE id=$1", id)
	helper.Handle(err1)
	fmt.Println(value1)
}
