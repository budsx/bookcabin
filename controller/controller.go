package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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
	flightId, err := strconv.ParseInt(r.URL.Query().Get("flightId"), 10, 64)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, errors.New("invalid flight id"))
		return
	}
	seatMapRequest := dto.SeatMapRequest{
		FlightID: flightId,
	}
	seatMapResponse, err := c.service.GetSeatMap(r.Context(), seatMapRequest)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	JSONResponse(w, http.StatusOK, seatMapResponse)
}

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err.Error())
}
