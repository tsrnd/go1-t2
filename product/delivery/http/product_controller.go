package http

import (
	"net/http"
	"encoding/json"
	
	"github.com/go-chi/chi"
	"goweb2/product/usecase"
	"goweb2/services/cache"
)

// ProductController type
type ProductController struct {
  	Usecase usecase.ProductUsecase
  	Cache   cache.Cache
}

// NewProductController func
func NewProductController(r chi.Router, uc usecase.ProductUsecase, c cache.Cache) *ProductController {
	handler := &ProductController{
		Usecase: uc,
		Cache:   c,
	}
	r.Get("/products", handler.ListProducts)
	return handler
}

// ListProducts func
func (ctrl *ProductController) ListProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	products, err := ctrl.Usecase.GetLimit()
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}