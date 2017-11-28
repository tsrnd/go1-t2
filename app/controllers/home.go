package controllers

import (
	"go_web_app/views"
	"net/http"

	"go_web_app/helper"

	"github.com/julienschmidt/httprouter"
)

type HomeController struct {
	helper.Controller
}

var Homes HomeController

func (self HomeController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {

	// Pretend to lookup cats in some way.
	Hobbies := []string{"heathcliff", "garfield"}
	// render the view
	return views.Homes.Index.Render(w, Hobbies)
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
