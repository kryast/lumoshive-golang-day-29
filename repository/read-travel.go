package repository

import (
	"database/sql"
	"day-29/model"
	"fmt"
	"time"
)

type TravelRepositoryDB struct {
	DB *sql.DB
}

func NewTravelRepository(db *sql.DB) TravelRepositoryDB {
	return TravelRepositoryDB{DB: db}
}

func (r TravelRepositoryDB) GetTravel() ([]model.Travel, error) {
	query := `
		SELECT 
			t.id ,
			p.name AS place_name,
			p.price AS place_price,
			p.detail AS place_detail,
			e.date_event AS event_date,
			COALESCE(AVG(r.rating), 0) AS review_rating,
			g.photo_url AS photo_url
		FROM travels t
		JOIN places p ON t.place_id = p.id
		JOIN events e ON t.event_id = e.id
		LEFT JOIN reviews r ON t.review_id = r.id
		LEFT JOIN galleries g ON p.gallery_id = g.id
		GROUP BY t.id, p.name, p.price, p.detail, e.date_event, g.photo_url
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var travels []model.Travel

	for rows.Next() {
		var travel model.Travel
		var dateEvent time.Time
		var reviewRating sql.NullFloat64

		err := rows.Scan(
			&travel.ID,
			&travel.PlaceName,
			&travel.PlacePrice,
			&travel.PlaceDetail,
			&dateEvent,
			&reviewRating,
			&travel.PhotoURL,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		travel.EventDate = dateEvent
		if reviewRating.Valid {
			travel.ReviewRating = reviewRating.Float64
		} else {
			travel.ReviewRating = 0
		}

		travels = append(travels, travel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return travels, nil
}
