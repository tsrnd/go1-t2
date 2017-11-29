package controllers

import (
	"goweb2/app/models"
	"goweb2/views"
	"net/http"

	"goweb2/helper"

	"github.com/julienschmidt/httprouter"
)

type HomeController struct {
	helper.Controller
}

var Homes HomeController

func (self HomeController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	val, err := models.AllTests()
	if err != nil {
		return err
	}
	return views.Homes.Index.Render(w, val)
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
