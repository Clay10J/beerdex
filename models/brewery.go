package models

import (
	"time"

	"github.com/clay10j/beerdex/internal/database"
)

type Brewery struct {
	Id        int       `json:"brewery_id"`
	Name      string    `json:"brewery_name"`
	CreatedAt time.Time `json:"created_at"`
	City      string    `json:"city"`
	State     string    `json:"state"`
}

func DatabaseBreweryToBrewery(brewery database.Brewery) Brewery {
	return Brewery{
		Id:        int(brewery.BreweryID),
		Name:      brewery.BreweryName,
		CreatedAt: brewery.CreatedAt.Time,
		City:      brewery.City,
		State:     brewery.State,
	}
}
