package models

import "time"

type Rating struct {
	BeerId    int       `json:"beer_id"`
	UserId    int       `json:"user_id"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}
