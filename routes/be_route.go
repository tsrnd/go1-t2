package routes

import (
	"github.com/julienschmidt/httprouter"
	"goweb2/app/controllers/admin"
)

func BeRoute(router *httprouter.Router) {
	adminproduct := admin.ProductAdmin
	router.GET("/admin/products", adminproduct.Perform(adminproduct.Index))
	router.GET("/admin/product/delete/:id", adminproduct.Perform(adminproduct.Delete))
	router.GET("/admin/product/showForm", adminproduct.Perform(adminproduct.ShowFormStore))
}
