package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/shivambaku/go-web-templ-htmx-tailwind-demo/internal/auth"
)

func (h *Handler) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := h.DB.GetUserByUsername(r.Context(), params.Username)
	if err != nil {
		responseError(w, http.StatusInternalServerError, "Couldn't get user")
		return
	}

	err = auth.ValidatePasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		responseError(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	auth.SetSessionToken(w, user.ID, time.Hour*24)
}

func (h *Handler) handlerLogout(w http.ResponseWriter, r *http.Request) {
	err := auth.ClearSessionToken(w, r)
	if err != nil {
		responseError(w, http.StatusBadRequest, "Couldn't clear session token")
		return
	}
}

func (h *Handler) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	err := auth.RefreshSessionToken(w, r, time.Hour*24)
	if err != nil {
		responseError(w, http.StatusUnauthorized, "Couldn't refresh session token")
		return
	}
}
