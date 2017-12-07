package controllers

import (
	"goweb2/app/models"
	"goweb2/helper"
	"goweb2/views"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController struct {
	helper.Controller
}

var Product ProductController

func (self ProductController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	id := ps.ByName("id")
	product, err := models.ShowProduct(id)
	if err != nil {
		return err
	}
	compact := map[string]interface{}{
		"Title":   "THIS IS PRODUCT DETAIL PAGE!",
		"Product": product,
		"Url":     helper.BaseUrl(),
	}

	return views.Product.Show.Render(w, r, compact)
}

func (self *ProductController) ReqKey(a helper.Action) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.FormValue("key") != "topsecret" {
			http.Error(w, "Invalid key.", http.StatusUnauthorized)
		} else {
			self.Controller.Perform(a)(w, r, ps)
		}
	})
}
