package routes

import (
	"goweb2/app/controllers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Route() {
	router := httprouter.New()
	home := controllers.Homes
	router.GET("/", home.Perform(home.Index))
	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
