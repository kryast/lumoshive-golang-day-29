package model

import (
	"time"
)

type Travel struct {
	ID               int       `json:"id"`
	PlaceName        string    `json:"place_name"`
	PlacePrice       int       `json:"place_price"`
	PlaceDetail      string    `json:"place_detail"`
	EventDate        time.Time `json:"event_date"`
	ReviewRating     float64   `json:"review_rating"`
	PhotoURL         string    `json:"photo_url,omitempty"`
	PhotoDescription string    `json:"photo_description,omitempty"`
}
