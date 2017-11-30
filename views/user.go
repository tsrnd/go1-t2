package views

import (
	"html/template"
	"log"
	"path/filepath"

	"goweb2/helper"
)

type UserView struct {
	helper.View
	Feed  helper.Page
	Login helper.Page
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
	registerFiles := append(UserFiles(), "templates/users/register.html")
	loginFiles := append(UserFiles(), "templates/users/login.html")
	User.Create = helper.Page{
		Template: template.Must(template.New("register").ParseFiles(registerFiles...)),
		Layout:   "my_layout",
	}
	User.Login = helper.Page{
		Template: template.Must(template.New("login").ParseFiles(loginFiles...)),
		Layout:   "my_layout",
	}
}
