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
	return r
}

// addUserRoutes func
func addUserRoutes(r chi.Router, db *sql.DB, c cache.Cache) {
	repo := userRepo.NewUserRepository(db)
	uc := userCase.NewUserUsecase(repo)
	userDeliver.NewUserController(r, uc, c)
}
