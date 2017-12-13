package product

import (
	"html/template"
	"log"
	"path/filepath"
	"goweb2/views/viewAdmin"
)

type FormAddProductAdminView struct {
	viewAdmin.View
	Feed viewAdmin.Page
}

var FormAddProductAdmin FormAddProductAdminView

func FormAddProductAdminFiles() []string {
	files, err := filepath.Glob("templates/admin/product/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, viewAdmin.LayoutAdminFiles()...)
	files = append(files, viewAdmin.LayoutAdminFilesIncludes()...)
	
	return files
}

func init() {
	detailFiles := append(FormAddProductAdminFiles(), "templates/admin/product/showFormAdd.html")
	FormAddProductAdmin.Index = viewAdmin.Page{
		Template: template.Must(template.New("showFormAdd").ParseFiles(detailFiles...)),
		Layout:   "admin_layout",
	}
}
