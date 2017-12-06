package views

import (
	"html/template"
	"log"
	"path/filepath"
)

type UserView struct {
	View
	Feed    Page
	Login   Page
	Contact Page
}

var User UserView

func UserFiles() []string {
	files, err := filepath.Glob("templates/users/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, LayoutFiles()...)
	files = append(files, LayoutFilesIncludes()...)
	return files
}

func init() {
	registerFiles := append(UserFiles(), "templates/users/register.html")
	loginFiles := append(UserFiles(), "templates/users/login.html")
	contactFiles := append(UserFiles(), "templates/users/contact.html")
	User.Create = Page{
		Template: template.Must(template.New("register").ParseFiles(registerFiles...)),
		Layout:   "my_layout",
	}
	User.Login = Page{
		Template: template.Must(template.New("login").ParseFiles(loginFiles...)),
		Layout:   "my_layout",
	}
	User.Contact = Page{
		Template: template.Must(template.New("contact").ParseFiles(contactFiles...)),
		Layout:   "my_layout",
	}
}
