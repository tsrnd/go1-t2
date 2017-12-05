package routes

import (
	"goweb2/app/controllers"
	"goweb2/app/controllers/admin"
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
	productAdmin := admin.Product
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
	router.GET("/admin/products", productAdmin.Perform(productAdmin.Index))

	// controller page carts
	router.GET("/carts", cart.Perform(cart.Index))
	router.POST("/remove-cart", cart.Delete)
	router.POST("/update-cart", cart.Update)
	router.POST("/add-to-cart", cart.Store)
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	FeRoute(router)
	BeRoute(router)
	log.Println("Starting server on :", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
