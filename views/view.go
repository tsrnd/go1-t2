package views

import (
	"goweb2/app/models"
	"goweb2/helper"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func LayoutFiles() []string {
	files, err := filepath.Glob("templates/layouts/*.html")
	if err != nil {
		log.Panic(err)
	}
	return files
}

func LayoutFilesIncludes() []string {
	filesInc, err := filepath.Glob("templates/layouts/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	return filesInc
}

func LayoutAdminFiles() []string {
	files, err := filepath.Glob("templates/layouts/admin/*.html")
	if err != nil {
		log.Panic(err)
	}
	return files
}

func LayoutAdminFilesIncludes() []string {
	filesInc, err := filepath.Glob("templates/layouts/admin/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	return filesInc
}

type View struct {
	Index  Page
	Show   Page
	New    Page
	Create Page
	Edit   Page
	Update Page
	Delete Page
}

type Page struct {
	Template *template.Template
	Layout   string
}

func (self *Page) Render(w http.ResponseWriter, r *http.Request, data interface{}) (a error) {
	sessionData := map[string]interface{}{
		"AuthData": models.GetAuth(r),
		"UrlPath":  r.URL.Path,
	}
	result := map[string]interface{}{
		"Data":        data,
		"PrivateData": sessionData,
		"Cart":        GetCart(r),
		"Url":         helper.BaseUrl(),
	}
	if err := self.Template.ExecuteTemplate(w, self.Layout, result); err != nil {
		log.Printf("Failed to execute template: %v", err)
	}
	return a
}

func GetCart(r *http.Request) interface{} {
	order := helper.GetSession("order", r)
	orderId, _ := strconv.Atoi(order)
	listCart, _ := models.ShowCart(orderId)
	log.Println("list", listCart)
	return listCart
}
