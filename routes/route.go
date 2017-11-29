package routes

import (
	"goweb2/app/controllers"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Route() {
	router := httprouter.New()
	home := controllers.Homes

	router.GET("/", home.Perform(home.Index))
	port := os.Getenv("PORT")
	log.Println("Starting server on :", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
