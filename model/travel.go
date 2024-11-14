package model

import "time"

type Travel struct {
	ID            int       `json:"id,omitempty"`
	PlaceID       int       `json:"place_id,omitempty"`
	PlaceName     string    `json:"place_name"`
	Price         int       `json:"price"`
	Detail        string    `json:"detail"`
	PhotoURL      string    `json:"photo_url"`
	EventDate     time.Time `json:"event_date"`
	TransactionID int       `json:"transaction_id,omitempty"`
	Status        string    `json:"status,omitempty"`
	Rating        int       `json:"rating"`
}
