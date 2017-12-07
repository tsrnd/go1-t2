package controllers

import (
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
	order := helper.GetSession("order", r)
	orderId, _ := strconv.Atoi(order)
	listCart, _ := models.ShowCart(orderId)
	compact := map[string]interface{}{
		"Title": "THIS IS A CARTS PAGE!",
		"Url":   helper.BaseUrl(),
		"Data":  listCart,
	}
	// fmt.Println("cart id", cartDetailId)
	return views.Carts.Index.Render(w, r, compact)

}
func CreateOder(w http.ResponseWriter, r *http.Request) int64 {
	order := helper.GetSession("order", r)
	fmt.Println("Check order:", order)
	if order == "" || order == "0" {
		newOrder, _ := models.InsertOrder()
		fmt.Println("create order:", newOrder)
		helper.SetSession("order", strconv.Itoa(int(newOrder)), w)
		return newOrder
	}
	orderId, _ := strconv.ParseInt(order, 10, 32)
	return orderId
}
func (self CartController) Store(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		order := CreateOder(w, r)
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		quantity, _ := strconv.Atoi(r.FormValue("quantity"))
		idProduct, _ := strconv.ParseInt(r.FormValue("product_id"), 10, 32)
		cartDetailId, _ := models.InsertCartDetail(price, quantity, 1, idProduct, order)
		fmt.Println("add", cartDetailId)
		http.Redirect(w, r, helper.Url("carts"), http.StatusSeeOther)
	} else {
		http.Redirect(w, r, helper.BaseUrl(), http.StatusSeeOther)
	}
}

func (self CartController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		order := helper.GetSession("order", r)
		w.Header().Set("Content-Type", "application/json")
		ajax := StatusAjax{1}
		if order == "" {
			ajax = StatusAjax{0}
			js, _ := json.Marshal(ajax)
			w.Write(js)
		} else {
			detailCartId := r.FormValue("detailCartId")
			cartDetailId, _ := strconv.ParseInt(detailCartId, 10, 32)
			result, _ := models.Remove(cartDetailId)
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
		orderId := helper.GetSession("order", r)
		w.Header().Set("Content-Type", "application/json")
		ajax := StatusAjax{1}
		if orderId == "" {
			ajax = StatusAjax{0}
			js, _ := json.Marshal(ajax)
			w.Write(js)
		} else {
			detailCartId := r.FormValue("detailCartId")
			cartDetailId, _ := strconv.ParseInt(detailCartId, 10, 32)
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
