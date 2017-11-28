package views

import (
	"html/template"
	"log"
	"path/filepath"

	"go_web_app/helper"
)

type HomesView struct {
	helper.View
	Feed helper.Page
}

var Homes HomesView

func HomesFiles() []string {
	files, err := filepath.Glob("templates/homes/includes/*.tmpl")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, helper.LayoutFiles()...)
	return files
}

func init() {
	indexFiles := append(HomesFiles(), "templates/homes/index.tmpl")
	Homes.Index = helper.Page{
		Template: template.Must(template.New("index").ParseFiles(indexFiles...)),
		Layout:   "my_layout",
	}
}
