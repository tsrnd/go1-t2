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
func GetProductLimit() ([]*Product, error) {
	query := "SELECT id, name, description, image, price, created_at, updated_at FROM products LIMIT 10"
	// Execute the query
	results, err := db.Query(query)
	// If there is an error opening the connection, handle it
	if err != nil {

		return nil, err
	}
	defer results.Close()

	products := make([]*Product, 0)
	for results.Next() {
		product := new(Product)
		// For each row, scan the result into our tag composite object
		err := results.Scan(&product.Id, &product.Name, &product.Description, &product.Image, &product.Price, &product.Created_at, &product.Updated_at)
		if err != nil {

			return nil, err
		}
		products = append(products, product)
	}
	if err = results.Err(); err != nil {

		return nil, err
	}

	return products, nil
}

func ShowProduct(id string) ([]*Product, error) {
	query := "SELECT id, name, description, image, price, created_at, updated_at FROM products WHERE id=?"
	results, err := db.Query(query, id)
	if err != nil {
		
		return nil, err
	}
	defer results.Close()

	products := make([]*Product, 0)
	for results.Next() {
		product := new(Product)
		// For each row, scan the result into our tag composite object
		err := results.Scan(&product.Id, &product.Name, &product.Description, &product.Image, &product.Price, &product.Created_at, &product.Updated_at)
		if err != nil {

			return nil, err
		}
		products = append(products, product)
	}
	if err = results.Err(); err != nil {

		return nil, err
	}

	return products, nil
}
