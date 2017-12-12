package product

import (
	"html/template"
	"log"
	"path/filepath"
	"goweb2/views/viewAdmin"
)

type ProductAdminView struct {
	viewAdmin.View
	Feed viewAdmin.Page
}

var ProductAdmin ProductAdminView

func ProductAdminFiles() []string {
	files, err := filepath.Glob("templates/admin/product/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, viewAdmin.LayoutAdminFiles()...)
	files = append(files, viewAdmin.LayoutAdminFilesIncludes()...)
	
	return files
}

func init() {
	detailFiles := append(ProductAdminFiles(), "templates/admin/product/list.html")
	ProductAdmin.Index = viewAdmin.Page{
		Template: template.Must(template.New("list").ParseFiles(detailFiles...)),
		Layout:   "admin_layout",
	}
}
