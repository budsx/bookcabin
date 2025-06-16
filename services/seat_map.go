package services

import (
	"context"

	"github.com/budsx/bookcabin/dto"
	"go.uber.org/zap"
)

func (s *BookCabinService) GetSeatMap(ctx context.Context, seatMapRequest dto.SeatMapRequest) (dto.SeatMapResponse, error) {
	s.logger.Info("GetSeatMap", zap.String("aircraftCode", seatMapRequest.AircraftCode))
	seatMapResponse := dto.SeatMapResponse{}

	aircraft, err := s.repo.DBReadWriter.ReadAircraftsByCode(ctx, seatMapRequest.AircraftCode)
	if err != nil {
		s.logger.Error("Failed to read aircrafts by code", zap.Error(err))
		return seatMapResponse, err
	}

	cabins, err := s.repo.DBReadWriter.ReadCabinsByAircraftID(ctx, int32(aircraft.ID))
	if err != nil {
		s.logger.Error("Failed to read cabins by aircraft id", zap.Error(err))
		return seatMapResponse, err
	}

	cabinsDTO := make([]dto.Cabin, 0)
	for _, cabin := range cabins {
		seatColumns, err := s.repo.DBReadWriter.ReadSeatColumnsByCabinID(ctx, int32(cabin.ID))
		if err != nil {
			s.logger.Error("Failed to read seat columns by cabin id", zap.Error(err))
			return seatMapResponse, err
		}

		seatColumnsDTO := make([]string, 0)
		for _, seatColumn := range seatColumns {
			seatColumnsDTO = append(seatColumnsDTO, seatColumn.ColumnCode)
		}

		seatRows, err := s.repo.DBReadWriter.ReadSeatRowsByCabinID(ctx, int32(cabin.ID))
		if err != nil {
			s.logger.Error("Failed to read seat rows by cabin id", zap.Error(err))
			return seatMapResponse, err
		}

		seatRowsDTO := make([]dto.SeatRow, 0)
		for _, seatRow := range seatRows {
			seatCodes, err := s.repo.DBReadWriter.ReadSeatCodesBySeatRowID(ctx, int32(seatRow.ID))
			if err != nil {
				s.logger.Error("Failed to read seat codes by seat row id", zap.Error(err))
				return seatMapResponse, err
			}

			seats, err := s.repo.DBReadWriter.ReadSeatsBySeatRowID(ctx, int32(seatRow.ID))
			if err != nil {
				s.logger.Error("Failed to read seats by seat row id", zap.Error(err))
				return seatMapResponse, err
			}

			seatsDTO := make([]dto.Seat, 0)
			for _, seat := range seats {
				seatDTO := dto.Seat{}

				if seat.Code != "" {
					seatDTO.Code = seat.Code
				}

				if seat.StorefrontSlotCode != "" {
					seatDTO.StorefrontSlotCode = seat.StorefrontSlotCode
				}

				if seat.RefundIndicator != "" {
					seatDTO.RefundIndicator = seat.RefundIndicator
				}

				if seat.FreeOfCharge {
					seatDTO.FreeOfCharge = seat.FreeOfCharge
				}

				seatCharacteristics, err := s.repo.DBReadWriter.ReadSeatCharacteristicsBySeatID(ctx, int32(seat.ID))
				if err != nil {
					s.logger.Error("Failed to read seat characteristics by seat id", zap.Error(err))
					return seatMapResponse, err
				}

				for _, seatCharacteristic := range seatCharacteristics {
					seatDTO.SeatCharacteristics = append(seatDTO.SeatCharacteristics, seatCharacteristic.Characteristic)
				}

				rawSeatCharacteristics, err := s.repo.DBReadWriter.ReadRawSeatCharacteristicsBySeatID(ctx, int32(seat.ID))
				if err != nil {
					s.logger.Error("Failed to read raw seat characteristics by seat id", zap.Error(err))
					return seatMapResponse, err
				}

				for _, rawSeatCharacteristic := range rawSeatCharacteristics {
					seatDTO.RawSeatCharacteristics = append(seatDTO.RawSeatCharacteristics, rawSeatCharacteristic.RawCharacteristic)
				}

				seatsDTO = append(seatsDTO, seatDTO)
			}

			seatRowsDTO = append(seatRowsDTO, dto.SeatRow{
				RowNumber: seatRow.RowNumber,
				SeatCodes: seatCodes,
				Seats:     seatsDTO,
			})
		}

		cabinsDTO = append(cabinsDTO, dto.Cabin{
			Deck:        cabin.Deck,
			SeatColumns: seatColumnsDTO,
			SeatRows:    seatRowsDTO,
			FirstRow:    cabin.FirstRow,
			LastRow:     cabin.LastRow,
		})
	}

	// TODO: get passenger id from request or Context or JWT
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
					},
				},
			},
		},
	}

	s.logger.Info("[RESPONSE] GetSeatMap", zap.Any("seatMapResponse", seatMapResponse))
	return seatMapResponse, nil
}
