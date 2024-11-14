package repository

import (
	"database/sql"
	"day-29/model"
	"fmt"
	"time"
)

func (r TravelRepositoryDB) GetTravelByID(id int) (model.Travel, error) {
	query := `
		SELECT 
			t.id ,
			p.name AS place_name,
			p.price AS place_price,
			p.detail AS place_detail,
			e.date_event AS event_date,
			r.rating AS review_rating
		FROM travels t
		JOIN places p ON t.place_id = p.id
		JOIN events e ON t.event_id = e.id
		LEFT JOIN reviews r ON t.review_id = r.id
		WHERE t.id = $1
	`

	var travel model.Travel
	var dateEvent time.Time
	var reviewRating sql.NullFloat64

	err := r.DB.QueryRow(query, id).Scan(
		&travel.ID,
		&travel.PlaceName,
		&travel.PlacePrice,
		&travel.PlaceDetail,
		&dateEvent,
		&reviewRating,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Travel{}, nil
		}
		return model.Travel{}, fmt.Errorf("error executing query: %w", err)
	}

	travel.EventDate = dateEvent
	if reviewRating.Valid {
		travel.ReviewRating = reviewRating.Float64
	} else {
		travel.ReviewRating = 0
	}

	return travel, nil
}
