package views

import (
	"html/template"
	"log"
	"path/filepath"
)

// It returns the number of bytes written and any write error encountered.
type CartsView struct {
	View
	Feed Page
}

var Carts CartsView

// It returns the number of bytes written and any write error encountered.
func CartsFiles() []string {
	files, err := filepath.Glob("templates/carts/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, LayoutFiles()...)
	files = append(files, LayoutFilesIncludes()...)
	return files
}

func init() {
	indexFiles := append(CartsFiles(), "templates/carts/index.html")
	Carts.Index = Page{
		Template: template.Must(template.New("index").ParseFiles(indexFiles...)),
		Layout:   "my_layout",
	}
}
