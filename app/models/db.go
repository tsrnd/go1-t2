package models

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var DB *sql.DB

func ConnectDB() {
	var err error
	// port := "5432"
	// val, ok := os.LookupEnv("DB_PORT")
	// if ok {
	// 	port = val
	// }
	dsn := fmt.Sprintf("host=%s port=%s user=%s password='%s' dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	db, err = gorm.Open(os.Getenv("DB_DRIVER"), dsn)
	if err != nil {
		panic(err.Error())
	}

}

func init() {
	var err error
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password='%s' dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	DB, err = sql.Open(os.Getenv("DB_DRIVER"), dsn)
	if err != nil {
		panic(err.Error())
	}
}
