package model

import "time"

type Travel struct {
	TravelID     int       `json:"travel_id"`
	PlaceName    string    `json:"place_name"`
	PlacePrice   int       `json:"place_price"`
	PlaceDetail  string    `json:"place_detail"`
	EventDate    time.Time `json:"event_date"`
	ReviewRating int       `json:"review_rating"`
}
