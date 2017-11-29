package helper

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Action func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) error

type Controller struct{}

func (c *Controller) Perform(a Action) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := a(w, r, ps); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
