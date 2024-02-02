package main

import (
	"github.com/labstack/echo"
	handler "github.com/shivambaku/fuufu-app/handlers"
)

func main() {
	e := echo.New()
	handler.Routes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
