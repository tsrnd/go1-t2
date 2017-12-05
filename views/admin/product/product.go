package views

import (
	"html/template"
	"log"
	"path/filepath"
	"goweb2/helper"
)

type ProductAdminView struct {
	helper.View
	Feed helper.Page
}

var ProductAdmin ProductAdminView

func ProductAdminFiles() []string {
	files, err := filepath.Glob("templates/admin/product/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, helper.LayoutAdminFiles()...)
	files = append(files, helper.LayoutAdminFilesIncludes()...)
	return files
}

func init() {
	detailFiles := append(ProductAdminFiles(), "templates/admin/product/list.html")
	ProductAdmin.Index = helper.Page{
		Template: template.Must(template.New("list").ParseFiles(detailFiles...)),
		Layout:   "admin_layout",
	}
}
