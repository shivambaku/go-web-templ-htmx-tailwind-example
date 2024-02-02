package handler

import (
	model "github.com/shivambaku/fuufu-app/models"
	"github.com/shivambaku/fuufu-app/views/user"

	"github.com/labstack/echo"
)

func handlerUserShow(c echo.Context) error {
	u := model.User{
		Email: "user@example.com",
	}
	return view(c, user.Show(u))
}
