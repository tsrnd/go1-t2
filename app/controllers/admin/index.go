package admin

import (
	"goweb2/helper"
	"goweb2/views/viewAdmin"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AdminIndexController struct {
	helper.Controller
}

var AdminIndexCtr AdminIndexController

func (self AdminIndexController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {

	compact := map[string]interface{}{}

	return viewAdmin.AdminIndex.Index.Render(w, r, compact)
}
