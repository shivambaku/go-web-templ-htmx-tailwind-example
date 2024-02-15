package main

import (
	"net/http"

	"github.com/go-chi/chi"
	handler "github.com/shivambaku/go-web-templ-htmx-tailwind-demo/handlers"
)

func (s *server) routes() *chi.Mux {
	handler := handler.Handler{
		DB: s.DB,
	}

	router := chi.NewRouter()

	// Static files
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	// View routes
	router.Get("/", handler.HandlerUserInfo)

	// API routes
	apiRouter := chi.NewRouter()
	apiRouter.Post("/users", handler.HandlerUserCreate)

	router.Mount("/api", apiRouter)

	return router
}
