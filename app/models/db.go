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
	dsn = "host=ec2-107-20-176-7.compute-1.amazonaws.com port=5432 user=csjzeohasarmuf" +
		"password='269122355cad43ad45ce1131f80ec194b4051ba20cde0d51cc94912d0e907ff2' dbname=d63dcu920h728n"
	//open a database connection
	db, err = gorm.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

}
