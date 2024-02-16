package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shivambaku/go-web-templ-htmx-tailwind-demo/internal/auth"
	"github.com/shivambaku/go-web-templ-htmx-tailwind-demo/internal/database"
	model "github.com/shivambaku/go-web-templ-htmx-tailwind-demo/models"
	page "github.com/shivambaku/go-web-templ-htmx-tailwind-demo/views/pages"
)

func (h *Handler) handlerUsersInfoView(w http.ResponseWriter, r *http.Request, user database.User) {
	responseView(w, r, page.Info(user))
}

func (h *Handler) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	responseJSON(w, http.StatusOK, model.UserToUserDTO(&user))
}

func (h *Handler) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var params parameters
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		responseError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		responseError(w, http.StatusInternalServerError, "Couldn't hash password")
		return
	}

	user, err := h.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		Username:       params.Username,
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	})
	if err != nil {
		responseError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	responseJSON(w, http.StatusCreated, model.UserToUserDTO(&user))
}
