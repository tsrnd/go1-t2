package views

import (
	"html/template"
	"log"
	"path/filepath"
)

// It returns the number of bytes written and any write error encountered.
type ProductView struct {
	View
	Feed Page
}

var Product ProductView

// It returns the number of bytes written and any write error encountered.
func ProductFiles() []string {
	files, err := filepath.Glob("templates/products/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, LayoutFiles()...)
	files = append(files, LayoutFilesIncludes()...)
	return files
}

func init() {
	detailFiles := append(ProductFiles(), "templates/products/detail.html")
	Product.Show = Page{
		Template: template.Must(template.New("detail").ParseFiles(detailFiles...)),
		Layout:   "my_layout",
	}
}
