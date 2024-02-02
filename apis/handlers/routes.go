package handler

import "github.com/labstack/echo"

func Routes(e *echo.Echo) {
	e.GET("/user", handlerUserShow)
}
