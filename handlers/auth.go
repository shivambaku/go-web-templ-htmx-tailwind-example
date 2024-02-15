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

	type response struct {
		Username string `json:"username"`
		Token    string `json:"token"`
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

	accessToken, err := auth.MakeJWT(
		user.ID,
		h.JWTSecret,
		time.Hour*24,
		auth.TokenTypeAccess,
	)
	if err != nil {
		responseError(w, http.StatusInternalServerError, "Couldn't make access token")
		return
	}

	responseJSON(w, http.StatusOK, response{
		Username: user.Username,
		Token:    accessToken,
	})
}
