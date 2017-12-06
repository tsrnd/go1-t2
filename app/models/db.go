package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() {
	var err error
	address := ""
	port := "5432"
	val, ok := os.LookupEnv("DB_PORT")
	if ok {
		port = val
	}
	if host, ok := os.LookupEnv("DB_HOST"); ok && host != "" {
		address = fmt.Sprintf("host=%s port=%d ", host, port)
	}
	//dsn
	dsn := fmt.Sprintf("%s user=%s password=%s dbname=%s sslmode=disable", address, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	dsn = "host=localhost port=5432 user=thinhnguyenv. " +
		"password= dbname=goweb2 sslmode=disable"
	//open a database connection
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
}
