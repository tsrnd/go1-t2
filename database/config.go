package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (db *sql.DB, err error) {

	address := ""
	port := "3306"
	val, ok := os.LookupEnv("DB_PORT")
	if ok {
		port = val
	}
	if host, ok := os.LookupEnv("DB_HOST"); ok && host != "" {
		address = fmt.Sprintf("(%s:%s)", host, port)
	}
	//dsn
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), address, os.Getenv("DB_DATABASE"), "utf8")

	//open a database connection
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db, err
}
