package library

import (
	"day-29/model"
	"encoding/json"
	"net/http"
)

func Response400(w http.ResponseWriter, message string) {
	badResponse := model.Response{
		Success:    false,
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}

	response, _ := json.MarshalIndent(badResponse, "", " ")
	jsonResponse, _ := w.Write(response)
	json.NewEncoder(w).Encode(jsonResponse)
}

func Response404(w http.ResponseWriter, message string) {
	badResponse := model.Response{
		Success:    false,
		StatusCode: 404,
		Message:    message,
	}

	response, _ := json.MarshalIndent(badResponse, "", " ")
	jsonResponse, _ := w.Write(response)
	json.NewEncoder(w).Encode(jsonResponse)
}

func Response200(w http.ResponseWriter, data any) {
	successResponse := model.Response{
		Success:    true,
		StatusCode: 200,
		Data:       data,
	}
	response, _ := json.MarshalIndent(successResponse, "", " ")
	jsonResponse, _ := w.Write(response)
	json.NewEncoder(w).Encode(jsonResponse)
}
