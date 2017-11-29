package controllers

import (
	"goweb2/helper"
	"goweb2/views"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	helper.Controller
}

var User UserController

func (self UserController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {

	return views.User.Index.Render(w, nil)

}
