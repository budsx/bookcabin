package services

import (
	"context"

	"github.com/budsx/bookcabin/dto"
	"go.uber.org/zap"
)

func (s *BookCabinService) GetSeatMap(ctx context.Context, seatMapRequest dto.SeatMapRequest) (dto.SeatMapResponse, error) {
	s.logger.Info("GetSeatMap", zap.String("aircraftCode", seatMapRequest.AircraftCode))
	seatMapResponse := dto.SeatMapResponse{}

	// Read Aircraft
	aircraft, err := s.repo.DBReadWriter.ReadAircraftsByCode(ctx, seatMapRequest.AircraftCode)
	if err != nil {
		s.logger.Error("Failed to read aircrafts by code", zap.Error(err))
		return seatMapResponse, err
	}

	// Read Cabins
	cabins, err := s.repo.DBReadWriter.ReadCabinsByAircraftID(ctx, aircraft.ID)
	if err != nil {
		s.logger.Error("Failed to read cabins by aircraft id", zap.Error(err))
		return seatMapResponse, err
	}
	cabinsIDs := make([]int64, 0)
	for _, cabin := range cabins {
		cabinsIDs = append(cabinsIDs, cabin.ID)
	}

	// Read Seat Columns
	seatColumns, err := s.repo.DBReadWriter.ReadSeatColumnsByCabinIDs(ctx, cabinsIDs)
	if err != nil {
		s.logger.Error("Failed to read seat columns by cabin ids", zap.Error(err))
		return seatMapResponse, err
	}
	seatColumnsDTO := make([]string, 0)
	for _, seatColumn := range seatColumns {
		seatColumnsDTO = append(seatColumnsDTO, seatColumn.ColumnCode)
	}

	// Read Seat Rows
	seatRows, err := s.repo.DBReadWriter.ReadSeatRowsByCabinIDs(ctx, cabinsIDs)
	if err != nil {
		s.logger.Error("Failed to read seat rows by cabin ids", zap.Error(err))
		return seatMapResponse, err
	}
	seatRowsIDs := make([]int64, 0)
	for _, seatRow := range seatRows {
		seatRowsIDs = append(seatRowsIDs, seatRow.ID)
	}

	// Read Seats
	seats, err := s.repo.DBReadWriter.ReadSeatsBySeatRowIDs(ctx, seatRowsIDs)
	if err != nil {
		s.logger.Error("Failed to read seats by seat row ids", zap.Error(err))
		return seatMapResponse, err
	}
	seatsIDs := make([]int64, 0)
	for _, seat := range seats {
		seatsIDs = append(seatsIDs, seat.ID)
	}

	// Read Seat Characteristics
	seatCharacteristics, err := s.repo.DBReadWriter.ReadSeatCharacteristicsBySeatIDs(ctx, seatsIDs)
	if err != nil {
		s.logger.Error("Failed to read seat characteristics by seat ids", zap.Error(err))
		return seatMapResponse, err
	}

	// Read Raw Seat Characteristics
	rawSeatCharacteristics, err := s.repo.DBReadWriter.ReadRawSeatCharacteristicsBySeatIDs(ctx, seatsIDs)
	if err != nil {
		s.logger.Error("Failed to read raw seat characteristics by seat ids", zap.Error(err))
		return seatMapResponse, err
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
		s.logger.Error("Failed to read passenger by id", zap.Error(err))
		return seatMapResponse, err
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
			// Emails: []string{
			// 	passenger.Email,
			// },
			// Phones: []string{
			// 	passenger.Phone,
			// },
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
