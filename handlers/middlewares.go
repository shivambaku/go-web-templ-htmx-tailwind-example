package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/shivambaku/go-web-templ-htmx-tailwind-demo/internal/auth"
	"github.com/shivambaku/go-web-templ-htmx-tailwind-demo/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (h *Handler) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			responseError(w, http.StatusUnauthorized, "Couldn't find JWT")
			return
		}

		userIDStr, err := auth.ValidateJWT(token, h.JWTSecret)
		if err != nil {
			responseError(w, http.StatusUnauthorized, "Couldn't validate JWT")
			return
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			responseError(w, http.StatusUnauthorized, "Couldn't parse user ID")
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
