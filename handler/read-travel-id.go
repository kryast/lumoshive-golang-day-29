package handler

// import (
// 	"day-29/library"
// 	"fmt"
// 	"net/http"
// )

// func (th TravelHandler) GetTravelByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	id := library.GetID(w, r)

// 	travel, err := th.travelService.GetTravelByID(id)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error: %s", err.Error()), http.StatusInternalServerError)
// 		return
// 	}

// 	library.Response200(w, travel)
// }
