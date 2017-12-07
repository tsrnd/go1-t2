package controllers

import (
	"goweb2/app/models"
	"goweb2/helper"
	"goweb2/views"
	"net/http"
	"strconv"

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
	orderId := helper.GetSession("order", r)

	if orderId == "" {
		return nil
	}
	orderIdInt, _ := strconv.ParseInt(orderId, 10, 32)
	listCart, _ := models.ShowOrder(orderIdInt)
	compact := map[string]interface{}{
		"Title":   "THIS IS PRODUCT DETAIL PAGE!",
		"Product": product,
		"Data":    listCart,
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
