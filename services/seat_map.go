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
	cabins, err := s.repo.DBReadWriter.ReadCabinsByAircraftID(ctx, int32(aircraft.ID))
	if err != nil {
		s.logger.Error("Failed to read cabins by aircraft id", zap.Error(err))
		return seatMapResponse, err
	}
	cabinsIDs := make([]int32, 0)
	for _, cabin := range cabins {
		cabinsIDs = append(cabinsIDs, int32(cabin.ID))
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
	seatRowsIDs := make([]int32, 0)
	seatRowsDTO := make([]dto.SeatRow, 0)
	for _, seatRow := range seatRows {
		seatRowsIDs = append(seatRowsIDs, int32(seatRow.ID))
		seatRowsDTO = append(seatRowsDTO, dto.SeatRow{
			RowNumber: seatRow.RowNumber,
		})
	}

	// Read Seats
	seats, err := s.repo.DBReadWriter.ReadSeatsBySeatRowIDs(ctx, seatRowsIDs)
	if err != nil {
		s.logger.Error("Failed to read seats by seat row ids", zap.Error(err))
		return seatMapResponse, err
	}
	seatsIDs := make([]int32, 0)
	for _, seat := range seats {
		seatsIDs = append(seatsIDs, int32(seat.ID))
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

	// Read Seat Codes
	seatCodes, err := s.repo.DBReadWriter.ReadSeatCodesBySeatRowIDs(ctx, seatRowsIDs)
	if err != nil {
		s.logger.Error("Failed to read seat codes by seat row ids", zap.Error(err))
		return seatMapResponse, err
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
