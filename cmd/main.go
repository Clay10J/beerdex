package main

import (
	"fmt"
	"log"
	"os"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/clay10j/beerdex/views"
)

func render(ctx echo.Context, tc templ.Component) error {
	return tc.Render(ctx.Request().Context(), ctx.Response())
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", helloWorld)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}

func helloWorld(c echo.Context) error {
	return render(c, views.Hello("Clay"))
}
