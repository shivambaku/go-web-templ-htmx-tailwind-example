package handler

import (
	model "github.com/shivambaku/fuufu-app/models"
	user "github.com/shivambaku/fuufu-app/views/users"

	"github.com/labstack/echo"
)

func HandlerUserShow(c echo.Context) error {
	u := model.User{
		Email: "user@example.com",
	}
	return view(c, user.Show(u))
}

func HandlerUserGet(c echo.Context) error {
	u := model.User{
		Email: "test@example.com",
	}
	return c.JSON(200, u)
}
