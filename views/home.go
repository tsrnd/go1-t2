package views

import (
	"html/template"
	"log"
	"path/filepath"

	"goweb2/helper"
)

type HomesView struct {
	helper.View
	Feed helper.Page
}

var Homes HomesView

func HomesFiles() []string {
	files, err := filepath.Glob("templates/homes/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, helper.LayoutFiles()...)
	files = append(files, helper.LayoutFilesIncludes()...)
	return files
}

func init() {
	indexFiles := append(HomesFiles(), "templates/homes/index.html")
	Homes.Index = helper.Page{
		Template: template.Must(template.New("index").ParseFiles(indexFiles...)),
		Layout:   "my_layout",
	}
}
