package lib

import (
	"goweb2/app/models"
	"goweb2/helper"
	"net/http"
)

func CheckAuth(r *http.Request) bool {
	authSS := helper.GetSession("AuthSession", r)
	if authSS != "" {
		// isLogin := models.CheckLoginWithSession(authSS)
		return models.CheckLoginWithSession(authSS)
	}
	return false
}
