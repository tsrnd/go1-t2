package controllers

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"goweb2/app/models"
	"goweb2/helper"
	"goweb2/views"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

// Struct CartController
type CartController struct {
	helper.Controller
}

// Struct Order session
type Orders struct {
	IdProduct int64
}

type StatusAjax struct {
	Status int `json:"statusCode"`
}

var Cart CartController

// init session
var store = sessions.NewCookieStore([]byte("secret-password"))

// list cart
func (self CartController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	session, err := store.Get(r, "carts")
	gob.Register(&Orders{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	if session.Values["orders"] == nil {
		orderId, _ := models.InsertOrder()
		session.Values["orders"] = orderId
		session.Save(r, w)
	}
	id := session.Values["orders"]
	// cartDetailId, _ := models.InsertCartDetail(100000, 1, 1, 3, id)
	listCart, _ := models.ShowCart(id)
	compact := map[string]interface{}{
		"Title": "THIS IS A CARTS PAGE!",
		"Data":  listCart,
		"Url":   helper.BaseUrl(),
	}
	// fmt.Println("cart id", cartDetailId)
	return views.Carts.Index.Render(w, compact)
}

func (self CartController) Store(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		session, _ := store.Get(r, "carts")
		orderId := session.Values["orders"]
		w.Header().Set("Content-Type", "application/json")
		ajax := StatusAjax{1}
		if orderId == nil {
			ajax = StatusAjax{0}
			js, _ := json.Marshal(ajax)
			w.Write(js)
		} else {
			detailCartId := r.FormValue("detailCartId")
			cartDetailId, _ := strconv.Atoi(detailCartId)
			result, _ := models.Remove(orderId, cartDetailId)
			if result == 0 {
				ajax = StatusAjax{0}
				js, _ := json.Marshal(ajax)
				w.Write(js)
			} else {
				js, _ := json.Marshal(ajax)
				w.Write(js)
				fmt.Println("js", result)
			}
		}
	} else {
		ajax := StatusAjax{0}
		js, _ := json.Marshal(ajax)
		w.Write([]byte(js))
	}

}
func (self CartController) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		session, _ := store.Get(r, "carts")
		orderId := session.Values["orders"]
		w.Header().Set("Content-Type", "application/json")
		ajax := StatusAjax{1}
		if orderId == nil {
			ajax = StatusAjax{0}
			js, _ := json.Marshal(ajax)
			w.Write(js)
		} else {
			detailCartId := r.FormValue("detailCartId")
			cartDetailId, _ := strconv.Atoi(detailCartId)
			quantity, _ := strconv.Atoi(r.FormValue("quantity"))
			totalPrice, _ := strconv.ParseFloat(r.FormValue("totalPrice"), 64)
			result, _ := models.Update(cartDetailId, quantity, totalPrice)
			if result == 0 {
				ajax = StatusAjax{0}
				js, _ := json.Marshal(ajax)
				w.Write(js)
			} else {
				js, _ := json.Marshal(ajax)
				w.Write(js)
				fmt.Println("js", result)
			}
		}
	} else {
		ajax := StatusAjax{0}
		js, _ := json.Marshal(ajax)
		w.Write([]byte(js))
	}
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
