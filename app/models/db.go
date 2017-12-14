package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDB() {
	var err error
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password='%s' dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	Db, err = sql.Open(os.Getenv("DB_DRIVER"), dsn)
	if err != nil {
		panic(err.Error())
	}
}

// func init() {
// 	var err error
// 	er := godotenv.Load()
// 	if er != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
// 	db, err = sql.Open(os.Getenv("DB_DRIVER"), dsn)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }
