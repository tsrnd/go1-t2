package views

import (
	"fmt"
	"goweb2/app/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func LayoutFiles() []string {
	files, err := filepath.Glob("templates/layouts/*.html")
	if err != nil {
		log.Panic(err)
	}
	return files
}

func LayoutFilesIncludes() []string {
	filesInc, err := filepath.Glob("templates/layouts/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	return filesInc
}

func LayoutAdminFiles() []string {
	files, err := filepath.Glob("templates/layouts/admin/*.html")
	if err != nil {
		log.Panic(err)
	}
	return files
}

func LayoutAdminFilesIncludes() []string {
	filesInc, err := filepath.Glob("templates/layouts/admin/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	return filesInc
}

type View struct {
	Index  Page
	Show   Page
	New    Page
	Create Page
	Edit   Page
	Update Page
	Delete Page
}

type Page struct {
	Template *template.Template
	Layout   string
}

func (self *Page) Render(w http.ResponseWriter, r *http.Request, data interface{}) (a error) {

	sessionData := map[string]interface{}{
		"AuthData": models.GetAuth(r),
	}
	result := map[string]interface{}{
		"Data":        data,
		"PrivateData": sessionData,
	}
	fmt.Println(result)
	if err := self.Template.ExecuteTemplate(w, self.Layout, result); err != nil {
		log.Printf("Failed to execute template: %v", err)
	}
	return a
	// return self.Template.ExecuteTemplate(w, self.Layout, result)
}
