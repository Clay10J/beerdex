package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/clay10j/beerdex/internal/database"
	"github.com/clay10j/beerdex/models"
)

func (cfg *HandlerConfig) HandleCreateBeer(c echo.Context) error {
	beer := new(models.Beer)
	if err := c.Bind(beer); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	params := database.CreateBeerParams{
		BeerName:  beer.Name,
		BreweryID: int64(beer.BreweryId),
		Abv:       beer.Abv,
		BeerType:  beer.BeerType,
	}
	newBeer, err := cfg.DBQueries.CreateBeer(c.Request().Context(), params)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return echo.NewHTTPError(http.StatusConflict, "beer already exists")
		}
		return err
	}

	b := models.DatabaseBeerToBeer(newBeer)
	return c.JSON(http.StatusCreated, b)
}

func (cfg *HandlerConfig) HandleGetAllBeers(c echo.Context) error {
	beers, err := cfg.DBQueries.GetBeers(c.Request().Context())
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusOK, "")
		}
		return err
	}

	beersList := make([]models.Beer, len(beers))
	for _, beer := range beers {
		newBeer := models.DatabaseBeerToBeer(beer)
		beersList = append(beersList, newBeer)
	}
	return c.JSON(http.StatusOK, beersList)
}

func (cfg *HandlerConfig) HandleGetBeer(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	beer, err := cfg.DBQueries.GetBeer(c.Request().Context(), int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusOK, "")
		}
		return err
	}

	b := models.DatabaseBeerToBeer(beer)
	return c.JSON(http.StatusOK, b)
}

func (cfg *HandlerConfig) HandleUpdateBeer(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	// Start a transaction
	tx, err := cfg.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // Rollback the transaction if it's not committed

	qtx := cfg.DBQueries.WithTx(tx)
	// Get the current state of the beer record
	currentBeer, err := qtx.GetBeer(c.Request().Context(), int64(id))
	if err != nil {
		return err
	}

	// Parse the request body to update the beer
	beer := new(models.Beer)
	if err := c.Bind(beer); err != nil {
		return err
	}

	params := database.UpdateBeerByIDParams{
		BeerName:  beer.Name,
		BreweryID: int64(beer.BreweryId),
		Abv:       beer.Abv,
		BeerType:  beer.BeerType,
		BeerID:    int64(id),
	}

	// overwrite empty fields to use the current db record values for update
	if params.BeerName == "" {
		params.BeerName = currentBeer.BeerName
	}
	if params.BreweryID == 0 {
		params.BreweryID = currentBeer.BreweryID
	}
	if params.Abv == 0.0 {
		params.Abv = currentBeer.Abv
	}
	if params.BeerType == "" {
		params.BeerType = currentBeer.BeerType
	}

	// Update the beer record based on the request
	updatedBeer, err := cfg.DBQueries.UpdateBeerByID(c.Request().Context(), params)
	if err != nil {
		return err
	}

	// Commit the transaction if everything is successful
	if err := tx.Commit(); err != nil {
		return err
	}

	b := models.DatabaseBeerToBeer(updatedBeer)
	return c.JSON(http.StatusOK, b)
}

func (cfg *HandlerConfig) HandleDeleteBeer(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	if err := cfg.DBQueries.DeleteBeer(c.Request().Context(), int64(id)); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("beer %s deleted", idStr))
}
