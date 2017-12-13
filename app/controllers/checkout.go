package controllers

import (
	"goweb2/app/models"
	"goweb2/helper"
	"goweb2/views"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Struct CartController
type CheckoutController struct {
	helper.Controller
}

var Checkout CheckoutController

// list cart
func (self CheckoutController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	order := helper.GetSession("order", r)
	if order == "" {
		http.Redirect(w, r, helper.BaseUrl(), http.StatusSeeOther)
	}
	user := models.GetAuth(r)
	idUser, _ := strconv.ParseInt(user.Id, 10, 64)
	if idUser == 0 {
		http.Redirect(w, r, helper.Url("login"), http.StatusSeeOther)
	}
	compact := map[string]interface{}{
		"Title": "THIS IS A CHECKOUT PAGE!",
	}
	return views.Checkout.Index.Render(w, r, compact)

}

func (self CheckoutController) Store(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	if r.Method == "POST" {
		user := models.GetAuth(r)
		idUser, _ := strconv.ParseInt(user.Id, 10, 64)
		address := r.FormValue("address")
		totalPrice, _ := strconv.ParseFloat(r.FormValue("totalprice"), 64)
		order := helper.GetSession("order", r)
		ok, _ := models.SetCurrentOrder(order, idUser, totalPrice, address)
		if ok == 0 {
			log.Println("error")
		}
		helper.ClearSession("order", w)
		http.Redirect(w, r, helper.BaseUrl(), http.StatusSeeOther)
	}
	return nil
}

func (self *CheckoutController) ReqKey(a helper.Action) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.FormValue("key") != "topsecret" {
			http.Error(w, "Invalid key.", http.StatusUnauthorized)
		} else {
			self.Controller.Perform(a)(w, r, ps)
		}
	})
}
