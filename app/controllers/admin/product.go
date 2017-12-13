package admin

import (
	"goweb2/app/models/admin"
	"goweb2/helper"
	"goweb2/views/viewAdmin/product"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"regexp"
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

func (self ProductAdminController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	reg, _ := regexp.Compile(`\d+$`)
	productId := reg.FindString(r.URL.Path)
	admin.DeleteProductById(productId)
	compact := map[string]interface{}{
		"Title":    "THIS IS A HOME PAGE!",
		"Url":      helper.BaseUrl(),
	}
	
	return product.ProductAdmin.Index.Render(w, r, compact)
}

func (self ProductAdminController) ShowFormStore(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	compact := map[string]interface{}{
		"Title":    "THIS IS A HOME PAGE!",
		"Url":      helper.BaseUrl(),
	}
	
	return product.FormAddProductAdmin.Index.Render(w, r, compact)
}