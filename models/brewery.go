package models

import "time"

type Brewery struct {
	Id        int       `json:"brewery_id"`
	Name      string    `json:"brewery_name"`
	CreatedAt time.Time `json:"created_at"`
	City      string    `json:"city"`
	State     string    `json:"state"`
}
