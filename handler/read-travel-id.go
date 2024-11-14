package handler

import (
	"day-29/library"
	"day-29/model"
	"fmt"
	"net/http"
)

func (h *TravelHandler) GetTravelByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := library.GetID(w, r)

	travel, err := h.travelService.GetTravelByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching travel: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if (travel == model.Travel{}) {
		http.Error(w, "Travel not found", http.StatusNotFound)
		return
	}

	library.Response200(w, travel)
}
