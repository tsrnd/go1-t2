package admin

import (
	"goweb2/app/models"
	"time"
)

/**
 * Product struct
 */
type Category struct {
	Id         int
	Name       string
	Created_at time.Time
	Updated_at time.Time
}

/**
 * Get 10 products to show on admin home page
 */
func GetCategory() ([]*Category, error) {
	var categories []*Category
	db := models.DB
	rows, err := db.Query("SELECT id, name, created_at, updated_at FROM categories")
	if err != nil {

		return nil, err
	}
	for rows.Next() {
		category := &Category{}
		rows.Scan(&category.Id, &category.Name, &category.Created_at, &category.Updated_at)
		categories = append(categories, category)
	}

	return categories, nil
}
