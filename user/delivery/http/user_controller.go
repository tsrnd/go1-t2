package http

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/go-chi/chi"
	"goweb2/user/usecase"
	"goweb2/user/delivery/http/request"
	"goweb2/services/cache"
)

// UserController type
type UserController struct {
  	Usecase usecase.UserUsecase
  	Cache   cache.Cache
}

// NewUserController func
func NewUserController(r chi.Router, uc usecase.UserUsecase, c cache.Cache) *UserController {
	handler := &UserController{
		Usecase: uc,
		Cache:   c,
	}
	r.Get("/user/{userID}", handler.Get)
	return handler
}

// Get func
func (ctrl *UserController) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	requestID := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(requestID)
	if err != nil {
		panic(err)
	}
	var request request.UserGetRequest
	request.ID = int64 (userID)
	user, err := ctrl.Usecase.GetByID(request.ID)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}