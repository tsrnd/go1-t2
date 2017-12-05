package middleware

import (
	"goweb2/app/models"
	"goweb2/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// func AuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
var RedirectIfAuthenticated = func(f httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authSS := helper.GetSession("AuthSession", r)
		if authSS != "" {
			isLogin := models.CheckLoginWithSession(authSS)
			if isLogin == true {
				http.Redirect(w, r, "/", 302)
			}
		}
		f(w, r, ps)
	}
}
