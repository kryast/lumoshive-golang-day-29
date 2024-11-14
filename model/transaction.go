package model

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID       int          `json:"id"`
	TravelID int          `json:"travel_id" validate:"required"`
	Name     string       `json:"name" validate:"required,min=3,max=100"`
	Email    string       `json:"email" validate:"required,email"`
	Phone    int          `json:"phone" validate:"required"`
	Message  string       `json:"message" validate:"required,max=500"`
	Status   string       `json:"status" validate:"required"`
	CreateAt time.Time    `json:"created_at"`
	UpdateAt time.Time    `json:"updated_at"`
	DeleteAt sql.NullTime `json:"deleted_at,omitempty"`
}
