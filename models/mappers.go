package model

import (
	"github.com/shivambaku/go-web-templ-htmx-tailwind-demo/internal/database"
)

func UserToUserDTO(user *database.User) UserDTO {
	return UserDTO{
		ID:        user.ID,
		UserName:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
