package model

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID       int
	TravelID int
	Name     string
	Email    string
	Phone    int
	Message  string
	Status   string
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt sql.NullTime
}
