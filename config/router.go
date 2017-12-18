package config

import (
	"database/sql"
	
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"goweb2/services/cache"
	
	userRepo "goweb2/user/repository"
	userCase "goweb2/user/usecase"
	userDeliver "goweb2/user/delivery/http"

	productRepo "goweb2/product/repository"
	productCase "goweb2/product/usecase"
	productDeliver "goweb2/product/delivery/http"
)

// Router func
func Router(db *sql.DB, c cache.Cache) *chi.Mux{
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	addUserRoutes(r, db, c)
	addProductRoutes(r, db, c)
	return r
}

// addUserRoutes func
func addUserRoutes(r chi.Router, db *sql.DB, c cache.Cache) {
	repo := userRepo.NewUserRepository(db)
	uc := userCase.NewUserUsecase(repo)
	userDeliver.NewUserController(r, uc, c)
}

// addProductRoutes func
func addProductRoutes(r chi.Router, db *sql.DB, c cache.Cache) {
	repo := productRepo.NewProductRepository(db)
	uc := productCase.NewProductUsecase(repo)
	productDeliver.NewProductController(r, uc, c)
}
