package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) Routes() *chi.Mux {
	router := chi.NewRouter()

	// Static files
	fs := http.FileServer(http.Dir("assets"))
	router.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	// View routes
	router.Get("/", h.middlewareAuth(h.HandlerUsersInfoView))
	router.Get("/login", h.HandlerLoginView)

	// Auth routes
	authRouter := chi.NewRouter()
	authRouter.Post("/login", h.HandlerLogin)
	authRouter.Post("/logout", h.HandlerLogout)
	authRouter.Post("/refresh", h.HandlerRefresh)
	router.Mount("/auth", authRouter)

	// API routes
	apiRouter := chi.NewRouter()
	apiRouter.Post("/users", h.HandlerUsersCreate)
	apiRouter.Get("/users/me", h.middlewareAuth(h.HandlerUsersGet))
	router.Mount("/api", apiRouter)

	return router
}
