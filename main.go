package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	handler "github.com/shivambaku/go-web-templ-htmx-tailwind-demo/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable must be set")
	}

	e := echo.New()
	routes(e)
	e.Logger.Fatal(e.Start(":" + port))
}

func routes(e *echo.Echo) {
	e.Static("/", "assets")

	e.GET("/", handler.HandlerUserShow)
	e.GET("/user", handler.HandlerUserGet)
}
