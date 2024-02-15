package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) Routes() *chi.Mux {
	router := chi.NewRouter()

	// Static files
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	// View routes
	router.Get("/", h.handlerUsersInfo)

	// Non API routes
	router.Post("/login", h.handlerLogin)

	// API routes
	apiRouter := chi.NewRouter()
	apiRouter.Post("/users", h.handlerUsersCreate)
	apiRouter.Get("/users/me", h.middlewareAuth(h.handlerUsersGet))
	router.Mount("/api", apiRouter)

	return router
}
