package viewAdmin

import (
	"goweb2/app/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

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

func (self *Page) Render(w http.ResponseWriter, r *http.Request, data interface{}) (err error) {
	sessionData := map[string]interface{}{
		"AuthData": models.GetAuth(r),
		"UrlPath":  r.URL.Path,
	}
	result := map[string]interface{}{
		"Data":        data,
		"PrivateData": sessionData,
	}
	if err := self.Template.ExecuteTemplate(w, self.Layout, result); err != nil {
		log.Printf("Failed to execute template: %v", err)
	}
	
	return err
}
