package views

import (
	"html/template"
	"log"
	"path/filepath"
)

// It returns the number of bytes written and any write error encountered.
type CheckoutView struct {
	View
	Feed Page
}

var Checkout CheckoutView

// It returns the number of bytes written and any write error encountered.
func CheckoutsFiles() []string {
	files, err := filepath.Glob("templates/checkouts/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, LayoutFiles()...)
	files = append(files, LayoutFilesIncludes()...)
	return files
}

func init() {
	indexFiles := append(CartsFiles(), "templates/checkouts/index.html")
	Checkout.Index = Page{
		Template: template.Must(template.New("index").ParseFiles(indexFiles...)),
		Layout:   "my_layout",
	}
}
