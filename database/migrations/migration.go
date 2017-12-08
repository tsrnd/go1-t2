package main

import (
	"fmt"
	"goweb2/app/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	var db *gorm.DB
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	db, _ = gorm.Open(os.Getenv("DB_DRIVER"), dsn)
	if err != nil {
		panic(err.Error())
	}
	//create table
	db.AutoMigrate(
		models.User{},
		models.Product{},
		models.Category{},
		models.Order{},
		models.CartDetail{},
	)

	// // //Add Foreign Key
	db.Model(models.Product{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Order{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(models.CartDetail{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
	db.Model(models.CartDetail{}).AddForeignKey("order_id", "orders(id)", "RESTRICT", "RESTRICT")

	// // //Add index
	db.Model(models.Product{}).AddIndex("idx_category_id", "category_id")
	db.Model(models.Order{}).AddIndex("idx_user_id", "user_id")
	db.Model(models.CartDetail{}).AddIndex("idx_product_id", "product_id")
	db.Model(models.CartDetail{}).AddIndex("idx_order_id", "order_id")

	// // //Add index unique
	db.Model(models.User{}).AddUniqueIndex("idx_user_name", "username")
	db.Model(models.User{}).AddUniqueIndex("idx_email", "email")
}
