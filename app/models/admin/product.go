package admin

import (
	"fmt"
	"goweb2/app/models"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
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

func GetProduct(id int64) (Product, error) {
	var product Product
	err := models.DB.QueryRow("SELECT id, name, description, image, price, created_at, updated_at FROM products where id = $1 limit 1", id).Scan(&product.Id, &product.Name, &product.Description, &product.Image, &product.Price, &product.Created_at, &product.Updated_at)
	if err != nil {
		return product, err
	}
	return product, nil
}

func CreateProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	r.ParseMultipartForm(0) // Parses the request body
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	categoryId, _ := strconv.ParseInt(r.FormValue("category_id"), 10, 32)
	file, handler, err := r.FormFile("image")
	image := ""
	if err == nil {
		defer file.Close()
		image = updateImage(file, handler)
	}
	fmt.Println(image)
	_, createErr := models.DB.Exec("INSERT INTO products(name, description, price, image, category_id) VALUES ($1, $2, $3, $4, $5);", name, description, price, image, categoryId)
	if createErr != nil {
		return createErr
	}
	return nil
}

func updateImage(file multipart.File, handler *multipart.FileHeader) string {

	f, err := os.OpenFile("./public/images/product/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer f.Close()
	io.Copy(f, file)
	return handler.Filename
}
