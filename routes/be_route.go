package routes

import (
	"goweb2/app/controllers/admin"

	"github.com/julienschmidt/httprouter"
)

func BeRoute(router *httprouter.Router) {
	adminproduct := admin.ProductAdmin
	adminIndex := admin.AdminIndexCtr
	router.GET("/admin/products", adminproduct.Perform(adminproduct.Index))
	// router.GET("/admin/product/:id", adminproduct.Perform(adminproduct.Edit))
	router.GET("/admin/product/create", adminproduct.Perform(adminproduct.Create))
	router.POST("/admin/product/", adminproduct.Perform(adminproduct.Store))
	router.GET("/admin", adminIndex.Perform(adminIndex.Index))
	router.GET("/admin/product/delete/:id", adminproduct.Perform(adminproduct.Delete))
	router.GET("/admin/product/showForm", adminproduct.Perform(adminproduct.ShowFormStore))
}
