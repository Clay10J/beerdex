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

func (cfg *HandlerConfig) HandleCreateBrewery(c echo.Context) error {
	brewery := new(models.Brewery)
	if err := c.Bind(brewery); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	params := database.CreateBreweryParams{
		BreweryName: brewery.Name,
		City:        brewery.City,
		State:       brewery.State,
	}
	newBrewery, err := cfg.DBQueries.CreateBrewery(c.Request().Context(), params)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return echo.NewHTTPError(http.StatusConflict, "brewery already exists")
		}
		return err
	}

	b := models.DatabaseBreweryToBrewery(newBrewery)
	return c.JSON(http.StatusCreated, b)
}

func (cfg *HandlerConfig) HandleGetAllBreweries(c echo.Context) error {
	breweries, err := cfg.DBQueries.GetBreweries(c.Request().Context())
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusOK, "")
		}
		return err
	}

	breweriesList := make([]models.Brewery, len(breweries))
	for _, brewery := range breweries {
		newBrewery := models.DatabaseBreweryToBrewery(brewery)
		breweriesList = append(breweriesList, newBrewery)
	}
	return c.JSON(http.StatusOK, breweriesList)
}

func (cfg *HandlerConfig) HandleGetBrewery(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	brewery, err := cfg.DBQueries.GetBrewery(c.Request().Context(), int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusOK, "")
		}
		return err
	}

	b := models.DatabaseBreweryToBrewery(brewery)
	return c.JSON(http.StatusOK, b)
}

func (cfg *HandlerConfig) HandleUpdateBrewery(c echo.Context) error {
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
	// Get the current state of the brewery record
	currentBrewery, err := qtx.GetBrewery(c.Request().Context(), int64(id))
	if err != nil {
		return err
	}

	// Parse the request body to update the brewery
	brewery := new(models.Brewery)
	if err := c.Bind(brewery); err != nil {
		return err
	}

	params := database.UpdateBreweryByIDParams{
		BreweryName: brewery.Name,
		City:        brewery.City,
		State:       brewery.State,
		BreweryID:   int64(id),
	}

	// overwrite empty fields to use the current db record values for update
	if params.BreweryName == "" {
		params.BreweryName = currentBrewery.BreweryName
	}
	if params.City == "" {
		params.City = currentBrewery.City
	}
	if params.State == "" {
		params.State = currentBrewery.State
	}

	// Update the brewery record based on the request
	updatedBrewery, err := cfg.DBQueries.UpdateBreweryByID(c.Request().Context(), params)
	if err != nil {
		return err
	}

	// Commit the transaction if everything is successful
	if err := tx.Commit(); err != nil {
		return err
	}

	b := models.DatabaseBreweryToBrewery(updatedBrewery)
	return c.JSON(http.StatusOK, b)
}

func (cfg *HandlerConfig) HandleDeleteBrewery(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	if err := cfg.DBQueries.DeleteBrewery(c.Request().Context(), int64(id)); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("brewery %s deleted", idStr))
}
