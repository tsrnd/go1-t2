package views

import (
	"html/template"
	"log"
	"path/filepath"
)

type HomesView struct {
	View
	Feed Page
}

var Homes HomesView

func HomesFiles() []string {
	files, err := filepath.Glob("templates/homes/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, LayoutFiles()...)
	files = append(files, LayoutFilesIncludes()...)
	return files
}

func init() {
	indexFiles := append(HomesFiles(), "templates/homes/index.html")
	Homes.Index = Page{
		Template: template.Must(template.New("index").ParseFiles(indexFiles...)),
		Layout:   "my_layout",
	}
}
