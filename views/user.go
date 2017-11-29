package views

import (
	"html/template"
	"log"
	"path/filepath"

	"goweb2/helper"
)

type UserView struct {
	helper.View
	Feed helper.Page
}

var User UserView

func UserFiles() []string {
	files, err := filepath.Glob("templates/users/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, helper.LayoutFiles()...)
	files = append(files, helper.LayoutFilesIncludes()...)
	return files
}

func init() {
	indexFiles := append(UserFiles(), "templates/users/index.html")
	User.Index = helper.Page{
		Template: template.Must(template.New("index").ParseFiles(indexFiles...)),
		Layout:   "my_layout",
	}
}
