package main

import (
	"go_web_app/database"
	"go_web_app/routes"
)

func main() {
	routers.Route()
	database.Connect()
}
