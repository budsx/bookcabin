package services

import (
	"context"
	"fmt"
	"time"

	"github.com/budsx/bookcabin/dto"
	"go.uber.org/zap"
)

func (s *BookCabinService) GetSeatMap(ctx context.Context, seatMapRequest dto.SeatMapRequest) (dto.SeatMapResponse, error) {
	seatMapResponse := dto.SeatMapResponse{}

	// Read Flight
	bookingFlight, err := s.repo.DBReadWriter.ReadBookingFlightByID(ctx, seatMapRequest.FlightID)
	if err != nil {
		s.logger.Error("Failed to read flight by id", zap.Error(err), zap.Int64("flightId", seatMapRequest.FlightID))
		return seatMapResponse, dto.ErrFlightNotFound
	}

	// Read Booking
	booking, err := s.repo.DBReadWriter.ReadBookingByID(ctx, bookingFlight.BookingID)
	if err != nil {
		s.logger.Error("Failed to read booking by id", zap.Error(err), zap.Int64("bookingId", bookingFlight.BookingID))
		return seatMapResponse, dto.ErrBookingNotFound
	}

	// Read Aircraft
	aircraft, err := s.repo.DBReadWriter.ReadAircraftsByCode(ctx, booking.Equipment)
	if err != nil {
		s.logger.Error("Failed to read aircrafts by code", zap.Error(err), zap.String("equipment", booking.Equipment))
		return seatMapResponse, dto.ErrAircraftNotFound
	}

	// Read Cabins
	cabins, err := s.repo.DBReadWriter.ReadCabinsByAircraftID(ctx, aircraft.ID)
	if err != nil {
		s.logger.Error("Failed to read cabins by aircraft id", zap.Error(err), zap.Int64("aircraftId", aircraft.ID))
		return seatMapResponse, dto.ErrSeatMapUnavailable
	}
	cabinsIDs := make([]int64, 0)
	for _, cabin := range cabins {
		cabinsIDs = append(cabinsIDs, cabin.ID)
	}

	// Read Seat Columns
	seatColumns, err := s.repo.DBReadWriter.ReadSeatColumnsByCabinIDs(ctx, cabinsIDs)
	if err != nil {
		s.logger.Error("Failed to read seat columns by cabin ids", zap.Error(err), zap.Any("cabinIds", cabinsIDs))
		return seatMapResponse, dto.ErrSeatMapUnavailable
	}
	seatColumnsDTO := make([]string, 0)
	for _, seatColumn := range seatColumns {
		seatColumnsDTO = append(seatColumnsDTO, seatColumn.ColumnCode)
	}

	// Read Seat Rows
	seatRows, err := s.repo.DBReadWriter.ReadSeatRowsByCabinIDs(ctx, cabinsIDs)
	if err != nil {
		s.logger.Error("Failed to read seat rows by cabin ids", zap.Error(err), zap.Any("cabinIds", cabinsIDs))
		return seatMapResponse, dto.ErrSeatMapUnavailable
	}
	seatRowsIDs := make([]int64, 0)
	for _, seatRow := range seatRows {
		seatRowsIDs = append(seatRowsIDs, seatRow.ID)
	}

	// Read Seats
	seats, err := s.repo.DBReadWriter.ReadSeatsBySeatRowIDs(ctx, seatRowsIDs)
	if err != nil {
		s.logger.Error("Failed to read seats by seat row ids", zap.Error(err), zap.Any("seatRowIds", seatRowsIDs))
		return seatMapResponse, dto.ErrSeatMapUnavailable
	}
	seatsIDs := make([]int64, 0)
	for _, seat := range seats {
		seatsIDs = append(seatsIDs, seat.ID)
	}

	// Read Seat Characteristics
	seatCharacteristics, err := s.repo.DBReadWriter.ReadSeatCharacteristicsBySeatIDs(ctx, seatsIDs)
	if err != nil {
		s.logger.Error("Failed to read seat characteristics by seat ids", zap.Error(err), zap.Any("seatIds", seatsIDs))
		return seatMapResponse, dto.ErrSeatMapUnavailable
	}

	// Read Raw Seat Characteristics
	rawSeatCharacteristics, err := s.repo.DBReadWriter.ReadRawSeatCharacteristicsBySeatIDs(ctx, seatsIDs)
	if err != nil {
		s.logger.Error("Failed to read raw seat characteristics by seat ids", zap.Error(err), zap.Any("seatIds", seatsIDs))
		return seatMapResponse, dto.ErrSeatMapUnavailable
	}

	// Read Seat Prices
	seatPrices, err := s.repo.DBReadWriter.ReadSeatPricesBySeatIDs(ctx, seatsIDs)
	if err != nil {
		s.logger.Error("Failed to read seat prices by seat ids", zap.Error(err), zap.Any("seatIds", seatsIDs))
		return seatMapResponse, dto.ErrSeatMapUnavailable
	}

	seatRowsDTO := make([]dto.SeatRow, 0)
	for _, seatRow := range seatRows {
		// Get seat codes for this row from seats data
		rowSeatCodes := make([]string, 0)
		codesMap := make(map[string]bool)

		// Get seats for this row
		rowSeats := make([]dto.Seat, 0)
		for _, seat := range seats {
			if seat.SeatRowID == seatRow.ID {
				if seat.StorefrontSlotCode != "" && !codesMap[seat.StorefrontSlotCode] {
					rowSeatCodes = append(rowSeatCodes, seat.StorefrontSlotCode)
					codesMap[seat.StorefrontSlotCode] = true
				}

				// Get seat characteristics for this seat
				seatChars := make([]string, 0)
				for _, char := range seatCharacteristics {
					if char.SeatID == seat.ID {
						seatChars = append(seatChars, char.Characteristic)
					}
				}

				// Get raw seat characteristics for this seat
				rawSeatChars := make([]string, 0)
				for _, rawChar := range rawSeatCharacteristics {
					if rawChar.SeatID == seat.ID {
						rawSeatChars = append(rawSeatChars, rawChar.RawCharacteristic)
					}
				}

				// Get seat prices for this sea
				seatPricesDTO := make([][]dto.Price, 0)
				seatTaxesDTO := make([][]dto.Price, 0)
				seatTotalDTO := make([][]dto.Price, 0)
				for _, price := range seatPrices {
					if price.SeatID == seat.ID {
						if price.PriceType == "price" {
							seatPricesDTO = append(seatPricesDTO, []dto.Price{
								{
									Amount:   float64(price.Amount),
									Currency: string(price.Currency),
								},
							})
						}
						if price.PriceType == "tax" {
							seatTaxesDTO = append(seatTaxesDTO, []dto.Price{
								{
									Amount:   float64(price.Amount),
									Currency: string(price.Currency),
								},
							})
						}
						if price.PriceType == "total" {
							seatTotalDTO = append(seatTotalDTO, []dto.Price{
								{
									Amount:   float64(price.Amount),
									Currency: string(price.Currency),
								},
							})
						}
					}
				}

				rowSeats = append(rowSeats, dto.Seat{
					StorefrontSlotCode:     seat.StorefrontSlotCode,
					Available:              seat.Available,
					Code:                   seat.Code,
					Entitled:               seat.Entitled,
					FeeWaived:              seat.FeeWaived,
					SeatCharacteristics:    seatChars,
					RefundIndicator:        seat.RefundIndicator,
					FreeOfCharge:           seat.FreeOfCharge,
					OriginallySelected:     seat.OriginallySelected,
					RawSeatCharacteristics: rawSeatChars,
					Prices: &dto.PriceInfo{
						Alternatives: seatPricesDTO,
					},
					Taxes: &dto.PriceInfo{
						Alternatives: seatTaxesDTO,
					},
					Total: &dto.PriceInfo{
						Alternatives: seatTotalDTO,
					},
				})
			}
		}

		seatRowsDTO = append(seatRowsDTO, dto.SeatRow{
			RowNumber: seatRow.RowNumber,
			SeatCodes: rowSeatCodes,
			Seats:     rowSeats,
		})
	}

	// Convert To DTO
	cabinsDTO := make([]dto.Cabin, 0)
	for _, cabin := range cabins {
		cabinsDTO = append(cabinsDTO, dto.Cabin{
			Deck:        cabin.Deck,
			SeatColumns: seatColumnsDTO,
			SeatRows:    seatRowsDTO,
			FirstRow:    cabin.FirstRow,
			LastRow:     cabin.LastRow,
		})
	}

	// Read Passenger
	// TODO: Get passenger id from request / context / token
	passenger, err := s.repo.DBReadWriter.ReadPassengerByID(ctx, 1)
	if err != nil {
		s.logger.Error("Failed to read passenger by id", zap.Error(err), zap.Int64("passengerId", 1))
		return seatMapResponse, dto.ErrPassengerNotFound
	}

	emails, err := s.repo.DBReadWriter.ReadPassengerEmail(ctx, passenger.ID)
	if err != nil {
		s.logger.Error("Failed to read passenger email", zap.Error(err), zap.Int64("passengerId", passenger.ID))
		return seatMapResponse, dto.ErrPassengerNotFound
	}

	phones, err := s.repo.DBReadWriter.ReadPassengerPhone(ctx, passenger.ID)
	if err != nil {
		s.logger.Error("Failed to read passenger phone", zap.Error(err), zap.Int64("passengerId", passenger.ID))
		return seatMapResponse, dto.ErrPassengerNotFound
	}

	passengerDTO := dto.Passenger{
		PassengerIndex:      int(passenger.PassengerIndex),
		PassengerNameNumber: passenger.PassengerNameNumber,
		PassengerDetails: dto.PassengerDetails{
			FirstName: passenger.FirstName,
			LastName:  passenger.LastName,
		},
		PassengerInfo: dto.PassengerInfo{
			DateOfBirth: passenger.DateOfBirth,
			Gender:      passenger.Gender,
			Type:        passenger.Type,
			Address: dto.Address{
				Street1:     passenger.Street1,
				Street2:     passenger.Street2,
				Postcode:    passenger.Postcode,
				State:       passenger.State,
				City:        passenger.City,
				Country:     passenger.Country,
				AddressType: passenger.AddressType,
			},
			Emails: emails,
			Phones: phones,
		},
	}
	seatMapResponse = dto.SeatMapResponse{
		SeatsItineraryParts: []dto.SeatsItineraryPart{
			{
				SegmentSeatMaps: []dto.SegmentSeatMap{
					{
						PassengerSeatMaps: []dto.PassengerSeatMap{
							{
								SeatSelectionEnabledForPax: true,
								SeatMap: dto.SeatMap{
									RowsDisabledCauses: []string{},
									AirCraft:           aircraft.Code,
									Cabins:             cabinsDTO,
								},
								Passenger: passengerDTO,
							},
						},
						Segment: dto.Segment{},
					},
				},
			},
		},
	}

	return seatMapResponse, nil
}

func (s *BookCabinService) SelectSeat(ctx context.Context, req dto.SeatSelectionRequest) (dto.SeatSelectionResponse, error) {
	s.logger.Info("SelectSeat",
		zap.Int64("flightId", req.FlightID),
		zap.String("seatCode", req.SeatCode),
		zap.String("passengerName", fmt.Sprintf("%s %s", req.PassengerInfo.FirstName, req.PassengerInfo.LastName)))

	response := dto.SeatSelectionResponse{}

	if req.SeatCode == "" {
		s.logger.Warn("Invalid seat code provided", zap.String("seatCode", req.SeatCode))
		return response, dto.ErrInvalidSeatCode
	}

	// Read Seat
	tx, err := s.repo.DBReadWriter.BeginTx(ctx, nil)
	if err != nil {
		s.logger.Error("Failed to begin transaction", zap.Error(err))
		return response, dto.ErrSeatNotAvailable
	}
	defer tx.Rollback()

	seat, err := s.repo.DBReadWriter.ReadSeatsByCode(ctx, tx, req.SeatCode)
	if err != nil {
		s.logger.Error("Failed to read seat by code", zap.Error(err), zap.String("seatCode", req.SeatCode))
		return response, dto.ErrSeatNotAvailable
	}

	s.logger.Info("READ SEAT FROM DB",
		zap.String("seatCode", seat.Code),
		zap.Int64("seatID", seat.ID),
		zap.Bool("available", seat.Available))

	if !seat.Available {
		s.logger.Warn("Seat is not available", zap.String("seatCode", req.SeatCode))
		return response, dto.ErrSeatNotAvailable
	}

	// Update Seat
	seat.Available = false
	s.logger.Info("UPDATING SEAT TO UNAVAILABLE",
		zap.String("seatCode", seat.Code),
		zap.Int64("seatID", seat.ID),
		zap.Bool("newAvailable", seat.Available))

	err = s.repo.DBReadWriter.UpdateSeat(ctx, tx, seat)
	if err != nil {
		s.logger.Error("Failed to update seat", zap.Error(err), zap.Int64("seatId", seat.ID))
		return response, dto.ErrSeatNotAvailable
	}

	// Commit Transaction
	err = tx.Commit()
	if err != nil {
		s.logger.Error("Failed to commit transaction", zap.Error(err))
		return response, dto.ErrSeatNotAvailable
	}

	s.logger.Info("TRANSACTION COMMITTED SUCCESSFULLY", zap.String("seatCode", req.SeatCode))

	selectedSeat := &dto.SelectedSeat{
		FlightID:      req.FlightID,
		SeatCode:      req.SeatCode,
		PassengerID:   1, // TODO: Get from auth context
		Status:        "selected",
		SelectionTime: time.Now().Format(time.RFC3339),
		PassengerInfo: req.PassengerInfo,
	}

	response.Success = true
	response.Message = fmt.Sprintf("Seat %s has been selected successfully", req.SeatCode)
	response.SelectedSeat = selectedSeat

	s.logger.Info("Seat selected successfully",
		zap.String("seatCode", req.SeatCode),
		zap.Int64("flightId", req.FlightID))

	return response, nil
}
