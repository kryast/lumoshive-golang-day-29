package repository

import (
	"database/sql"
	"day-29/model"
	"time"
)

func (r TravelRepositoryDB) GetTravelByID(id int) (model.Travel, error) {
	query := `
		SELECT 
			p.name AS place_name, 
			p.price, 
			p.detail, 
			e.date_event AS event_date, 
			r.rating, 
			p.photo_url
		FROM places p
		JOIN events e ON p.id = e.place_id
		JOIN transactions t ON e.id = t.event_id
		LEFT JOIN reviews r ON t.id = r.transaction_id
		WHERE p.id = $1`

	var travel model.Travel
	var dateEvent time.Time
	var rating sql.NullInt64

	err := r.DB.QueryRow(query, id).Scan(
		&travel.PlaceName,
		&travel.Price,
		&travel.Detail,
		&dateEvent,
		&rating,
		&travel.PhotoURL,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Travel{}, nil
		}
		return model.Travel{}, err
	}

	travel.EventDate = dateEvent

	if rating.Valid {
		travel.Rating = int(rating.Int64)
	} else {
		travel.Rating = 0
	}

	return travel, nil
}
