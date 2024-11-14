package model

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID       int
	EventID  int
	Status   string
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt sql.NullTime
}
