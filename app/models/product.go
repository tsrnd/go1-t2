package models

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
func GetProductLimit() ([]Product, error) {
	var products []Product
	results := db.Limit(10).Find(&products)
	if results.Error != nil {
		return nil, results.Error
	}
	return products, nil
}

func ShowProduct(id string) (Product, error) {
	var product Product
	results := db.First(&product, id)
	return product, results.Error
}
