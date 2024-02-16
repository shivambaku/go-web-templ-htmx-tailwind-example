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
	router.Get("/", h.middlewareAuth(h.handlerUsersInfoView))
	router.Get("/login", h.handlerLoginView)

	// Auth routes
	authRouter := chi.NewRouter()
	authRouter.Post("/login", h.handlerLogin)
	authRouter.Post("/logout", h.handlerLogout)
	authRouter.Post("/refresh", h.handlerRefresh)
	router.Mount("/auth", authRouter)

	// API routes
	apiRouter := chi.NewRouter()
	apiRouter.Post("/users", h.handlerUsersCreate)
	apiRouter.Get("/users/me", h.middlewareAuth(h.handlerUsersGet))
	router.Mount("/api", apiRouter)

	return router
}
