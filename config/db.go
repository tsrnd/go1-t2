package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	db "goweb2/services/database/sql"
)

// DB func
func DB() *sql.DB {
	er := godotenv.Load()
	if er != nil {
		log.Print("Error loading .env file")
	}
	dbDlct := os.Getenv("DATABASE_DLCT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	dbSslmode := os.Getenv("DATABASE_SSL_MODE")
	db, err := db.Connect(dbDlct, dbUser, dbPass, dbHost, dbPort, dbName, dbSslmode)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
