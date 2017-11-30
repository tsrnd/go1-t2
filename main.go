package main

import (
	"goweb2/app/models"
	"goweb2/routes"
)

func main() {
	models.ConnectDB()
	routes.Route()
}
