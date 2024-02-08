package models

import (
	"time"

	"github.com/clay10j/beerdex/internal/database"
)

type Beer struct {
	Id        int       `json:"beer_id"`
	Name      string    `json:"beer_name"`
	BreweryId int       `json:"brewery_id"`
	Abv       float64   `json:"abv"`
	CreatedAt time.Time `json:"created_at"`
	BeerType  string    `json:"beer_type"`
}

func DatabaseBeerToBeer(beer database.Beer) Beer {
	return Beer{
		Id:        int(beer.BeerID),
		Name:      beer.BeerName,
		BreweryId: int(beer.BreweryID),
		Abv:       beer.Abv,
		CreatedAt: beer.CreatedAt.Time,
		BeerType:  beer.BeerType,
	}
}
