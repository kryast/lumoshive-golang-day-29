package repository

import (
	"database/sql"
	"day-29/model"
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
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var travels []model.Travel
	for rows.Next() {
		var travel model.Travel
		var dateEvent time.Time
		var rating sql.NullInt64

		err := rows.Scan(
			&travel.PlaceName,
			&travel.Price,
			&travel.Detail,
			&dateEvent,
			&rating,
			&travel.PhotoURL,
		)
		if err != nil {
			return nil, err
		}

		travel.EventDate = dateEvent

		if rating.Valid {
			travel.Rating = int(rating.Int64)
		} else {
			travel.Rating = 0
		}

		travels = append(travels, travel)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return travels, nil
}
