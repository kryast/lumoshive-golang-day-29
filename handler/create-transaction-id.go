package handler

import (
	"day-29/library"
	"day-29/model"
	"day-29/service"
	"fmt"
	"net/http"
	"strconv"
)

type TransactionHandler struct {
	TransactionService service.TransactionService
}

func NewTransactionHandler(service service.TransactionService) TransactionHandler {
	return TransactionHandler{TransactionService: service}
}

func (th *TransactionHandler) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	id := library.GetID(w, r)

	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	message := r.FormValue("message")
	status := "In Process"
	if name == "" || email == "" || phone == "" || message == "" || status == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	phoneInt, err := strconv.Atoi(phone)
	if err != nil {
		library.Response400(w, err.Error())
		return
	}

	transaction := model.Transaction{
		TravelID: id,
		Name:     name,
		Email:    email,
		Phone:    phoneInt,
		Message:  message,
		Status:   status,
	}

	transactionID, err := th.TransactionService.CreateTransactionByTravelID(transaction)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating transaction: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	library.Response200(w, transactionID)
}
