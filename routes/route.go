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
	product := controllers.Product
	port := "8080"
	if val, ok := os.LookupEnv("PORT"); ok && val != "" {
		port = val
	}

	router.GET("/", home.Perform(home.Index))
	router.GET("/register", user.Perform(user.Register))
	router.POST("/register", user.Perform(user.Store))
	router.GET("/login", user.Perform(user.LoginPage))
	router.GET("/contact", user.Perform(user.ShowContactPage))
	router.GET("/product/:id", product.Perform(product.Show))
	// router.POST("/login", user.Perform(user.Login))

	// controller page carts
	router.GET("/carts", cart.Perform(cart.Index))
	router.POST("/remove-cart", cart.Delete)
	router.POST("/update-cart", cart.Update)
	router.POST("/add-to-cart", cart.Store)
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	log.Println("Starting server on :", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
