package controllers

import (
	"fmt"
	"goweb2/app/models"
	"goweb2/helper"
	"goweb2/views"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/**
 * HomeController struct
 */
type HomeController struct {
	helper.Controller
}

var Homes HomeController

/**
 * Show home page
 */
func (self HomeController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	products, err := models.GetProductLimit()
	if err != nil {

		return err
	}
	order := helper.GetSession("order", r)
	orderId, _ := strconv.Atoi(order)
	fmt.Println("idorer", orderId)
	ShowCart, _ := models.ShowCart(orderId)
	compact := map[string]interface{}{
		"Title":    "THIS IS A HOME PAGE!",
		"Products": products,
		"Url":      helper.BaseUrl(),
		"Data":     ShowCart,
	}

	return views.Homes.Index.Render(w, r, compact)
}

func (self *HomeController) ReqKey(a helper.Action) httprouter.Handle {

	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.FormValue("key") != "topsecret" {
			http.Error(w, "Invalid key.", http.StatusUnauthorized)
		} else {
			self.Controller.Perform(a)(w, r, ps)
		}
	})
}
