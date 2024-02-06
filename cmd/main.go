package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/clay10j/beerdex/handlers"
	"github.com/clay10j/beerdex/views"
)

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

	e.GET("/", homepage)
	e.GET("/api/beers", handlers.HandleGetAllBeers)
	e.POST("/api/beers", handlers.HandleCreateBeer)
	e.GET("/api/beers/:id", handlers.HandleGetBeer)
	e.PUT("/api/beers/:id", handlers.HandleUpdateBeer)
	e.DELETE("/api/beers/:id", handlers.HandleDeleteBeer)

	e.GET("/api/breweries", handlers.HandleGetAllBreweries)
	e.POST("/api/breweries", handlers.HandleCreateBrewery)
	e.GET("/api/breweries/:id", handlers.HandleGetBrewery)
	e.PUT("/api/breweries/:id", handlers.HandleUpdateBrewery)
	e.DELETE("/api/breweries/:id", handlers.HandleDeleteBrewery)

	e.GET("/api/ratings", handlers.HandleGetAllRatings)
	e.POST("/api/ratings", handlers.HandleCreateRating)
	e.GET("/api/ratings/:id", handlers.HandleGetRating)
	e.PUT("/api/ratings/:id", handlers.HandleUpdateRating)
	e.DELETE("/api/ratings/:id", handlers.HandleDeleteRating)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func render(ctx echo.Context, statusCode int, tc templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return tc.Render(ctx.Request().Context(), ctx.Response())
}

func homepage(c echo.Context) error {
	return render(c, http.StatusOK, views.Home())
}
