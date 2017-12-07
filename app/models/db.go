package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func ConnectDB() {
	var err error
	port := "5432"
	val, ok := os.LookupEnv("DB_PORT")
	if ok {
		port = val
	}
	sslmode := "true"
	if disable, check := os.LookupEnv("DB_SSLMODE"); check {
		sslmode = disable
	}
	//dsn
	dsn := fmt.Sprintf("host='%s' port='%d'  user='%s' password='%s' dbname=%s sslmode=%s", os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), sslmode)
	dsn = "host=localhost port=5432 user=thinhnguyenv. " +
		"password='' dbname=golang2 sslmode=disable"
	//open a database connection
	db, err = gorm.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

}
