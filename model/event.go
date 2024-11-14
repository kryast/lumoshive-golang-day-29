package model

import "time"

type Event struct {
	ID        int
	PlaceID   int
	DateEvent time.Time
}
