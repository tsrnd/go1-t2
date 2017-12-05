package middleware

import (
	"goweb2/lib"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var RedirectIfAuthenticated = func(f httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if lib.CheckAuth(r) == true {
			http.Redirect(w, r, "/", 302)
			return
		}
		f(w, r, ps)
	}
}
