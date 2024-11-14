package service

import (
	"day-29/model"
	"day-29/repository"
	"fmt"
)

type TransactionService struct {
	repo repository.TransactionRepositoryDB
}

func NewTransactionService(repo repository.TransactionRepositoryDB) TransactionService {
	return TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) CreateTransactionByTravelID(transaction model.Transaction) (int, error) {
	transactionID, err := s.repo.CreateTransactionByTravelID(transaction)
	if err != nil {
		return 0, fmt.Errorf("failed to create transaction: %w", err)
	}

	return transactionID, nil
}
