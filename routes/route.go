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
	user := controllers.User
	cart := controllers.Cart
	port := "8080"
	if val, ok := os.LookupEnv("PORT"); ok && val != "" {
		port = val
	}

	router.GET("/", home.Perform(home.Index))
	router.GET("/register", user.Perform(user.Register))
	router.POST("/register", user.Perform(user.Store))
	router.GET("/login", user.Perform(user.LoginPage))
	router.GET("/contact", user.Perform(user.ShowContactPage))
	// router.POST("/login", user.Perform(user.Login))

	// controller page carts
	router.GET("/carts", cart.Perform(cart.Index))
	router.POST("/remove-cart", cart.Store)
	router.POST("/update-cart", cart.Update)
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	log.Println("Starting server on :", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
