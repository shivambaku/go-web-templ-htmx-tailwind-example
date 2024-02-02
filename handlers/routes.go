package handler

import "github.com/labstack/echo"

func Routes(e *echo.Echo) {
	e.Static("/", "assets")

	e.GET("/user", handlerUserShow)
}
