package routes

import (
	"goweb2/app/controllers"
	"goweb2/app/middleware"
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
	var publicChain = []middleware.Middleware{
		middleware.RedirectIfAuthenticated,
	}
	router.GET("/", home.Perform(home.Index))
	router.GET("/register", middleware.BuildChain(user.Perform(user.Register), publicChain...))
	router.POST("/register", middleware.BuildChain(user.Perform(user.Store), publicChain...))
	router.GET("/login", middleware.BuildChain(user.Perform(user.LoginPage), publicChain...))
	router.POST("/login", middleware.BuildChain(user.Perform(user.Login), publicChain...))
	router.GET("/contact", user.Perform(user.ShowContactPage))
	router.GET("/product/:id", product.Perform(product.Show))

	// controller page carts
	router.GET("/carts", cart.Perform(cart.Index))
	router.POST("/remove-cart", cart.Delete)
	router.POST("/update-cart", cart.Update)
	router.POST("/add-to-cart", cart.Store)
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	log.Println("Starting server on :", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
