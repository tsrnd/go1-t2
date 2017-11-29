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
	port := "8080"
	if val, ok := os.LookupEnv("PORT"); ok && val != "" {
		port = val
	}

	router.GET("/", home.Perform(home.Index))

	router.ServeFiles("/public/*filepath", http.Dir("public"))
	log.Println("Starting server on :", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
