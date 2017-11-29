package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(username string, password string, db_name string, charset string) (db *sql.DB, err error) {

	//dsn
	dsn := fmt.Sprintf("%s:%s@/%s?charset=%s", username, password, db_name, charset)

	//open a database connection
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db, err
}
