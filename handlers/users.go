package handler

import (
	model "github.com/shivambaku/go-web-templ-htmx-tailwind-demo/models"
	user "github.com/shivambaku/go-web-templ-htmx-tailwind-demo/views/users"

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
