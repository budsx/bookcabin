package controller

import (
	"encoding/json"
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
		ErrorResponse(w, http.StatusBadRequest, dto.BookCabinError{
			Code:    "INVALID_FLIGHT_ID",
			Message: "Invalid flight ID format",
		})
		return
	}
	seatMapRequest := dto.SeatMapRequest{
		FlightID: flightId,
	}
	seatMapResponse, err := c.service.GetSeatMap(r.Context(), seatMapRequest)
	if err != nil {
		HandleServiceError(w, err)
		return
	}
	JSONResponse(w, http.StatusOK, seatMapResponse)
}

func (c *BookCabinController) SelectSeat(w http.ResponseWriter, r *http.Request) {
	var selectSeatRequest dto.SeatSelectionRequest

	if err := json.NewDecoder(r.Body).Decode(&selectSeatRequest); err != nil {
		ErrorResponse(w, http.StatusBadRequest, dto.BookCabinError{
			Code:    "INVALID_REQUEST_BODY",
			Message: "Invalid request body format",
		})
		return
	}

	if selectSeatRequest.FlightID == 0 {
		ErrorResponse(w, http.StatusBadRequest, dto.BookCabinError{
			Code:    "MISSING_FLIGHT_ID",
			Message: "Flight ID is required",
		})
		return
	}
	if selectSeatRequest.SeatCode == "" {
		ErrorResponse(w, http.StatusBadRequest, dto.BookCabinError{
			Code:    "MISSING_SEAT_CODE",
			Message: "Seat code is required",
		})
		return
	}
	if selectSeatRequest.PassengerInfo.FirstName == "" || selectSeatRequest.PassengerInfo.LastName == "" {
		ErrorResponse(w, http.StatusBadRequest, dto.BookCabinError{
			Code:    "MISSING_PASSENGER_INFO",
			Message: "Passenger first name and last name are required",
		})
		return
	}

	selectSeatResponse, err := c.service.SelectSeat(r.Context(), selectSeatRequest)
	if err != nil {
		HandleServiceError(w, err)
		return
	}

	JSONResponse(w, http.StatusOK, selectSeatResponse)
}

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if bookCabinErr, ok := err.(dto.BookCabinError); ok {
		json.NewEncoder(w).Encode(bookCabinErr)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"error": err.Error(),
		},
	)
}

func HandleServiceError(w http.ResponseWriter, err error) {
	if bookCabinErr, ok := err.(dto.BookCabinError); ok {
		switch bookCabinErr.Code {
		case "FLIGHT_NOT_FOUND", "BOOKING_NOT_FOUND", "AIRCRAFT_NOT_FOUND", "PASSENGER_NOT_FOUND":
			ErrorResponse(w, http.StatusNotFound, bookCabinErr)
		case "SEAT_NOT_AVAILABLE", "SEAT_ALREADY_SELECTED", "INVALID_SEAT_CODE":
			ErrorResponse(w, http.StatusConflict, bookCabinErr)
		case "SEAT_MAP_UNAVAILABLE", "DATA_ACCESS_ERROR":
			ErrorResponse(w, http.StatusServiceUnavailable, bookCabinErr)
		default:
			ErrorResponse(w, http.StatusInternalServerError, bookCabinErr)
		}
		return
	}
	ErrorResponse(w, http.StatusInternalServerError, dto.ErrInternalServerError)
}
