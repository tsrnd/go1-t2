package admin

import (
	"goweb2/app/models/admin"
	"goweb2/helper"
	"goweb2/views/viewAdmin/product"
	"net/http"
	"strconv"

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
		"Title":   "THIS IS A HOME PAGE!",
		"Product": products,
		"Url":     helper.BaseUrl(),
	}

	return product.ProductAdmin.Index.Render(w, r, compact)
}

func (self ProductAdminController) Edit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	id, _ := strconv.ParseInt(ps.ByName("id"), 10, 32)
	productDt, _ := admin.GetProduct(id)
	compact := map[string]interface{}{
		"Title":   "Update Product",
		"Product": productDt,
		"Url":     helper.BaseUrl(),
	}

	return product.ProductAdmin.Edit.Render(w, r, compact)
}

func (self ProductAdminController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	categories, _ := admin.GetCategory()
	compact := map[string]interface{}{
		"Title":      "Create Product",
		"Categories": categories,
		"Url":        helper.BaseUrl(),
	}

	return product.ProductAdmin.Create.Render(w, r, compact)
}

func (self ProductAdminController) Store(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {

	admin.CreateProduct(w, r, ps)
	http.Redirect(w, r, "/admin/products", 302)
	return nil
}
