package controller

import (
	"encoding/json"
	"net/http"

	"github.com/budsx/bookcabin/dto"
	"github.com/budsx/bookcabin/services"
)

type BookCabinController struct {
	service *services.BookCabinService
}

func NewBookCabinController(service *services.BookCabinService) *BookCabinController {
	return &BookCabinController{service: service}
}

func (c *BookCabinController) GetSeatMap(w http.ResponseWriter, r *http.Request) {
	seatMapRequest := dto.SeatMapRequest{
		AircraftCode: r.URL.Query().Get("aircraftCode"),
	}
	seatMapResponse, err := c.service.GetSeatMap(r.Context(), seatMapRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(seatMapResponse)
}
