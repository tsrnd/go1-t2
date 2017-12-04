package controllers

import (
	"fmt"
	"goweb2/app/models"
	"goweb2/helper"
	"goweb2/views"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	helper.Controller
}
type RegisterStatus struct {
	Status   bool
	Messeage string
	Title    string
}

var User UserController

func (self UserController) Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	compact := RegisterStatus{false, "", "Register"}
	return views.User.Create.Render(w, compact)

}

func (self UserController) Store(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	var compact RegisterStatus
	ok, errMsg := models.StoreUser(r)
	if ok {
		compact = RegisterStatus{true, "", "Register Success!"}
	} else {
		compact = RegisterStatus{false, errMsg, "Register Error"}
	}
	return views.User.Create.Render(w, compact)
}

func (self UserController) LoginPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {

	flashData := helper.GetFlash("auth_err_msg", w, r)

	compact := map[string]interface{}{
		"Title":     "Login",
		"FlashData": flashData,
	}
	fmt.Println(flashData)
	return views.User.Login.Render(w, compact)
}

func (self UserController) ShowContactPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	compact := map[string]interface{}{
		"Title": "Contact page",
	}
	return views.User.Contact.Render(w, compact)
}

func (self UserController) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (a error) {
	if ok, errMsg := models.Login(w, r); ok == false {
		helper.SetFlash(errMsg, w, r)
		http.Redirect(w, r, "/login", 302)
		return a
	}

	http.Redirect(w, r, "/", 302)
	return a
}
