package admin

import (
	"goweb2/app/models/admin"
	"goweb2/helper"
	"goweb2/views/viewAdmin/product"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type ProductAdminController struct {
	helper.Controller
}

var ProductAdmin ProductAdminController

func (self ProductAdminController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	products, err := admin.GetProductLimit()
	if err != nil {
		return err
	}
	compact := map[string]interface{}{
		"Title":    "THIS IS A HOME PAGE!",
		"Product":  products,
		"Url":      helper.BaseUrl(),
	}
	
	return product.ProductAdmin.Index.Render(w, r, compact)
}