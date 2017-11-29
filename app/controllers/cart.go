package controllers

import (
	"goweb2/views"
	"net/http"

	"goweb2/helper"

	"github.com/julienschmidt/httprouter"
)

type CartController struct {
	helper.Controller
}

var Cart CartController

func (self CartController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	compact := map[string]interface{}{
		"Title": "THIS IS A CARTS PAGE!",
		"Other": []int{1, 2, 3},
	}
	return views.Carts.Index.Render(w, compact)
}

func (self *CartController) ReqKey(a helper.Action) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.FormValue("key") != "topsecret" {
			http.Error(w, "Invalid key.", http.StatusUnauthorized)
		} else {
			self.Controller.Perform(a)(w, r, ps)
		}
	})
}
