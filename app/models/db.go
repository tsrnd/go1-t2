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
	// address := ""
	// port := "3306"
	// val, ok := os.LookupEnv("DB_PORT")
	// if ok {
	// 	port = val
	// }
	if host, ok := os.LookupEnv("DB_HOST"); ok && host != "" {
		// address = fmt.Sprintf("(%s:%s)", host, port)
	}
	//dsn
	dsn := "host=localhost port=5432 user=tienphamd. " +
	"password='' dbname=goweb2 sslmode=disable"
	//open a database connection
	fmt.Println(dsn)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
}
