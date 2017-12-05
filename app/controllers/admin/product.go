package controllers

import (
	// "encoding/gob"
	//"goweb2/app/models"
	"goweb2/helper"
	// "goweb2/views"
	"goweb2/views/admin/product"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/**
 * HomeController struct
 */
type ProductController struct {
	helper.Controller
}

var Product ProductController

/**
 * Show home page
 */
func (self ProductController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	// products, err := models.GetProductLimit()
	// if err != nil {

	// 	return err
	// }
	// compact := map[string]interface{}{
	// 	"Title":    "THIS IS A HOME PAGE!",
	// 	"Products": products,
	// 	"Data":     listCart,
	// 	"Url":      helper.BaseUrl(),
	// }

	return product.ProductAdmin.Index.Render(w)
}