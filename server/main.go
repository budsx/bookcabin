package main

import (
	"bookcabin/config"
	"bookcabin/dto"
	"bookcabin/models"
	"bookcabin/repository"
	"bookcabin/util/logger"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.WithError(err).Fatal("Failed to load config")
		return
	}

	repo, err := repository.NewBookCabinRepository(&repository.RepoConfig{
		DBConfig: repository.DBConfig{
			User:     cfg.DBUser,
			Password: cfg.DBPassword,
			Host:     cfg.DBHost,
			Port:     cfg.DBPort,
			DBName:   cfg.DBName,
		},
	})
	if err != nil {
		logger.WithError(err).Fatal("Failed to create repository")
		return
	}
	defer repo.Close()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		repo.HealthCheck()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Get("/seat-map", func(w http.ResponseWriter, r *http.Request) {
		// Logic Get Seat Map
		// Read Master Data
		// Read Passenger Data
		// Read Segment Data

		aircraft, err := repo.DBReadWriter.ReadAircraftsByCode(r.Context(), "738")
		if err != nil {
			logger.WithError(err).Error("Failed to read aircrafts by code")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger.Info("Aircraft retrieved successfully", zap.Any("aircraft", aircraft))

		cabins, err := repo.DBReadWriter.ReadCabinsByAircraftID(r.Context(), int32(aircraft.ID))
		if err != nil {
			logger.WithError(err).Error("Failed to read cabins by aircraft id")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger.Info("Cabins retrieved successfully", zap.Any("cabins", cabins))

		seatColumns := make([]models.SeatColumn, 0)
		for _, cabin := range cabins {
			seatColumns, err = repo.DBReadWriter.ReadSeatColumnsByCabinID(r.Context(), int32(cabin.ID))
			if err != nil {
				logger.WithError(err).Error("Failed to read seat columns by cabin id")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			logger.Info("Seat columns retrieved successfully", zap.Any("seatColumns", seatColumns))
		}

		seatRows := make([]models.SeatRow, 0)
		for _, cabin := range cabins {
			seatRows, err = repo.DBReadWriter.ReadSeatRowsByCabinID(r.Context(), int32(cabin.ID))
			if err != nil {
				logger.WithError(err).Error("Failed to read seat rows by cabin id")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			logger.Info("Seat rows retrieved successfully", zap.Any("seatRows", seatRows))
		}

		seats := make([]models.Seat, 0)
		for _, seatRow := range seatRows {
			seats, err = repo.DBReadWriter.ReadSeatsBySeatRowID(r.Context(), int32(seatRow.ID))
			if err != nil {
				logger.WithError(err).Error("Failed to read seats by seat row id")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			logger.Info("Seats retrieved successfully", zap.Any("seats", seats))
		}

		seatCharacteristics := make([]models.SeatCharacteristic, 0)
		for _, seat := range seats {
			seatCharacteristics, err = repo.DBReadWriter.ReadSeatCharacteristicsBySeatID(r.Context(), int32(seat.ID))
			if err != nil {
				logger.WithError(err).Error("Failed to read seat characteristics by seat id")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			logger.Info("Seat characteristics retrieved successfully", zap.Any("seatCharacteristics", seatCharacteristics))
		}
		for _, seat := range seats {
			rawSeatCharacteristics, err := repo.DBReadWriter.ReadRawSeatCharacteristicsBySeatID(r.Context(), int32(seat.ID))
			if err != nil {
				logger.WithError(err).Error("Failed to read raw seat characteristics by seat id")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			logger.Info("Raw seat characteristics retrieved successfully", zap.Any("rawSeatCharacteristics", rawSeatCharacteristics))
		}

		// Convert data to DTO Response
		seatColumnsDTO := make([]dto.SeatColumn, 0)
		for _, seatColumn := range seatColumns {
			seatColumnsDTO = append(seatColumnsDTO, dto.SeatColumn{
				ColumnCode: seatColumn.ColumnCode,
			})
		}

		seatRowsDTO := make([]dto.SeatRow, 0)
		for _, seatRow := range seatRows {
			seatRowsDTO = append(seatRowsDTO, dto.SeatRow{
				RowNumber: seatRow.RowNumber,
				SeatCodes: seatRow.SeatCodes,
			})
		}

		seatsDTO := make([]dto.Seat, 0)
		for _, seat := range seats {
			seatsDTO = append(seatsDTO, dto.Seat{
				Code: seat.Code,
			})
		}

		seatCharacteristicsDTO := make([]dto.SeatCharacteristic, 0)
		for _, seatCharacteristic := range seatCharacteristics {
			seatCharacteristicsDTO = append(seatCharacteristicsDTO, dto.SeatCharacteristic{
				Characteristic: seatCharacteristic.Characteristic,
			})
		}

		rawSeatCharacteristics := make([]models.RawSeatCharacteristic, 0)
		rawSeatCharacteristicsDTO := make([]dto.RawSeatCharacteristic, 0)
		for _, rawSeatCharacteristic := range rawSeatCharacteristics {
			rawSeatCharacteristicsDTO = append(rawSeatCharacteristicsDTO, dto.RawSeatCharacteristic{
				RawCharacteristic: rawSeatCharacteristic.RawCharacteristic,
			})
		}
		
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

		seatMapResponse := dto.SeatMapResponse{
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
								},
							},
						},
					},
				},
			},
		}

		// Response JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(seatMapResponse)
	})

	logger.Info("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
