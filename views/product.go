package views

import (
	"html/template"
	"log"
	"path/filepath"
	"goweb2/helper"
)

// It returns the number of bytes written and any write error encountered.
type ProductView struct {
	helper.View
	Feed helper.Page
}

var Product ProductView

// It returns the number of bytes written and any write error encountered.
func ProductFiles() []string {
	files, err := filepath.Glob("templates/products/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, helper.LayoutFiles()...)
	files = append(files, helper.LayoutFilesIncludes()...)
	return files
}

func init() {
	detailFiles := append(ProductFiles(), "templates/products/detail.html")
	Product.Show = helper.Page{
		Template: template.Must(template.New("detail").ParseFiles(detailFiles...)),
		Layout:   "my_layout",
	}
}
