package repository

import (
	"database/sql"
	"day-29/model"
	"fmt"
)

type TransactionRepositoryDB struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{DB: db}
}

func (r TransactionRepositoryDB) CreateTransactionByTravelID(t model.Transaction) (int, error) {
	query := `
		INSERT INTO transactions (travel_id, name, email, phone, message, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	var transactionID int
	err := r.DB.QueryRow(query, t.TravelID, t.Name, t.Email, t.Phone, t.Message, t.Status).Scan(&transactionID)
	if err != nil {
		return 0, fmt.Errorf("error creating transaction: %w", err)
	}

	return transactionID, nil
}
