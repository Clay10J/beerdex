package models

import "time"

type Beer struct {
	Id          int       `json:"beer_id"`
	Name        string    `json:"beer_name"`
	BreweryId   int       `json:"brewery_id"`
	Abv         float64   `json:"abv"`
	CreatedAt   time.Time `json:"created_at"`
	PackageType string    `json:"package_type"`
}
