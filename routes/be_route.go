package routes

import (
	"github.com/julienschmidt/httprouter"
	"goweb2/app/controllers/admin"
)

func BeRoute(router *httprouter.Router) {
	adminproduct := admin.ProductAdmin
	router.GET("/admin/products", adminproduct.Perform(adminproduct.Index))
}
