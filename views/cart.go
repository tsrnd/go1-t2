package views

import (
	"html/template"
	"log"
	"path/filepath"

	"goweb2/helper"
)

// It returns the number of bytes written and any write error encountered.
type CartsView struct {
	helper.View
	Feed helper.Page
}

var Carts CartsView

// It returns the number of bytes written and any write error encountered.
func CartsFiles() []string {
	files, err := filepath.Glob("templates/carts/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, helper.LayoutFiles()...)
	files = append(files, helper.LayoutFilesIncludes()...)
	return files
}

func init() {
	indexFiles := append(CartsFiles(), "templates/carts/index.html")
	Carts.Index = helper.Page{
		Template: template.Must(template.New("index").ParseFiles(indexFiles...)),
		Layout:   "my_layout",
	}
}
