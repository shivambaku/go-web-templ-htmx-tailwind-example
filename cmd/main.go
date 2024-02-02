package main

import (
	"github.com/labstack/echo"
	handler "github.com/shivambaku/fuufu-app/apis/handlers"
)

func main() {
	e := echo.New()
	e.Static("/", "assets")
	handler.Routes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
