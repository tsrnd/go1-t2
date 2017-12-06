package middleware

import (
	"goweb2/app/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var RedirectIfAuthenticated = func(f httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if models.CheckAuth(r) == true {
			http.Redirect(w, r, "/", 302)
			return
		}
		f(w, r, ps)
	}
}
