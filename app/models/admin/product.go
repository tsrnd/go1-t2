package admin

import (
	"goweb2/app/models"
	"time"
	"fmt"
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
	db := models.ConnectDB()
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

func DeleteProductById(id string) {
	db := models.ConnectDB()
	var products []*Product
	rows, err := db.Query("SELECT id, name, description, image, price, created_at, updated_at FROM products WHERE id=$1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		product := &Product{}
		rows.Scan(&product.Id, &product.Name, &product.Description, &product.Image, &product.Price, &product.Created_at, &product.Updated_at)
		products = append(products, product)
		fmt.Println(product)
	}
	value2, err2 := db.Exec("DELETE FROM cart_details WHERE product_id=$1", id)
	fmt.Println(err2, value2)
	value1, err1 := db.Exec("DELETE FROM products WHERE id=$1", id)
	fmt.Println(err1, value1)
}
