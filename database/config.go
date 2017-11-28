package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	db, err := sql.Open("mysql", "root:tuan123@/go_web_app?charset=utf8")
	rows, err := db.Query("SELECT * FROM test")
	fmt.Println(err)

	for rows.Next() {
		var name string
		var email string
		err = rows.Scan(&name, &email)
		fmt.Println(name)
		fmt.Println(email)
	}
}
