package viewAdmin

import (
	"html/template"
	"log"
	"path/filepath"
)

type AdminIndexView struct {
	Index Page
}

var AdminIndex AdminIndexView

func AdminIndexFiles() []string {
	files, err := filepath.Glob("templates/admin/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, LayoutAdminFiles()...)
	files = append(files, LayoutAdminFilesIncludes()...)

	return files
}

func init() {
	detailFiles := append(AdminIndexFiles(), "templates/admin/index.html")
	AdminIndex.Index = Page{
		Template: template.Must(template.New("index").ParseFiles(detailFiles...)),
		Layout:   "admin_layout",
	}
}
