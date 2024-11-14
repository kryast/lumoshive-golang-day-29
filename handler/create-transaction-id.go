package handler

import (
	"day-29/library"
	"day-29/model"
	"day-29/service"
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

	phoneInt, _ := strconv.Atoi(phone)

	transaction := model.Transaction{
		TravelID: id,
		Name:     name,
		Email:    email,
		Phone:    phoneInt,
		Message:  message,
		Status:   status,
	}

	err := library.ValidateStruct(transaction)
	if err != nil {
		library.Response400(w, err.Error())
		return
	}

	transactionID, err := th.TransactionService.CreateTransactionByTravelID(transaction)
	if err != nil {
		library.Response404(w, "ID Not Found")
		return
	}

	library.Response200(w, transactionID)
}
