package handler

import (
	"day-29/library"
	"day-29/service"
	"net/http"
)

type TravelHandler struct {
	travelService service.TravelService
}

func NewTravelHandler(service service.TravelService) TravelHandler {
	return TravelHandler{travelService: service}
}

func (th TravelHandler) GetTravelHandler(w http.ResponseWriter, r *http.Request) {

	travels, err := th.travelService.GetTravel()
	if err != nil {
		library.Response400(w, err.Error())
		return
	}

	library.Response200(w, travels)
}
