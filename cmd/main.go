package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"github.com/clay10j/beerdex/handlers"
	"github.com/clay10j/beerdex/internal/database"
	"github.com/clay10j/beerdex/views"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("TURSO_DB_URL")
	if dbURL == "" {
		log.Fatal("TURSO_DB_URL environment variable is not set")
	}

	dbAuthToken := os.Getenv("TURSO_DB_AUTH_TOKEN")
	if dbAuthToken == "" {
		log.Fatal("TURSO_DB_AUTH_TOKEN environment variable is not set")
	}

	db, err := sql.Open("libsql", fmt.Sprintf("%v?authToken=%v", dbURL, dbAuthToken))
	if err != nil {
		log.Fatal("Failed to open db", err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}
	dbQueries := database.New(db)

	handlerCfg := &handlers.HandlerConfig{
		DB:        db,
		DBQueries: dbQueries,
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", homepage)
	e.GET("/api/beers", handlerCfg.HandleGetAllBeers)
	e.POST("/api/beers", handlerCfg.HandleCreateBeer)
	e.GET("/api/beers/:id", handlerCfg.HandleGetBeer)
	e.PUT("/api/beers/:id", handlerCfg.HandleUpdateBeer)
	e.DELETE("/api/beers/:id", handlerCfg.HandleDeleteBeer)

	e.GET("/api/breweries", handlerCfg.HandleGetAllBreweries)
	e.POST("/api/breweries", handlerCfg.HandleCreateBrewery)
	e.GET("/api/breweries/:id", handlerCfg.HandleGetBrewery)
	e.PUT("/api/breweries/:id", handlerCfg.HandleUpdateBrewery)
	e.DELETE("/api/breweries/:id", handlerCfg.HandleDeleteBrewery)

	e.GET("/api/ratings", handlerCfg.HandleGetAllRatings)
	e.POST("/api/ratings", handlerCfg.HandleCreateRating)
	e.GET("/api/ratings/:id", handlerCfg.HandleGetRating)
	e.PUT("/api/ratings/:id", handlerCfg.HandleUpdateRating)
	e.DELETE("/api/ratings/:id", handlerCfg.HandleDeleteRating)

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
