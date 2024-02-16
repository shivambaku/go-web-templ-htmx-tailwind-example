package handler

import (
	"net/http"

	"github.com/shivambaku/go-web-templ-htmx-tailwind-demo/internal/auth"
	"github.com/shivambaku/go-web-templ-htmx-tailwind-demo/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (h *Handler) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := auth.GetSessionToken(r)
		if err != nil {
			responseError(w, http.StatusUnauthorized, "No session token")
			return
		}

		userID, err := auth.GetSessionUserId(sessionToken)
		if err != nil {
			responseError(w, http.StatusUnauthorized, "Invalid session token")
			return
		}

		user, err := h.DB.GetUserById(r.Context(), userID)
		if err != nil {
			responseError(w, http.StatusInternalServerError, "Couldn't get user")
			return
		}

		handler(w, r, user)
	}
}
