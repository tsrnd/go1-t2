package controllers

import (
	"encoding/gob"
	"goweb2/app/models"
	"goweb2/helper"
	"goweb2/views"
	"net/http"

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
	// show cart
	session, err := store.Get(r, "carts")
	gob.Register(&Orders{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	oderId := session.Values["orders"]
	listCart, _ := models.ShowCart(oderId)
	//end
	compact := map[string]interface{}{
		"Title":    "THIS IS A HOME PAGE!",
		"Products": products,
		"Data":     listCart,
		"Url":      helper.BaseUrl(),
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
