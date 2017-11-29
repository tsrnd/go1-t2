package models

import (
	"database/sql"
	"goweb2/database"
)

var db *sql.DB

func init() {
	var err error
	db, err = database.ConnectDB("root", "tuan123", "go_web_app", "utf8")
	if err != nil {
		panic(err.Error())
	}
}

type Test struct {
	name  string
	email string
}

func AllTests() ([]*Test, error) {
	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Test, 0)
	for rows.Next() {
		bk := new(Test)
		err := rows.Scan(&bk.name, &bk.email)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}
