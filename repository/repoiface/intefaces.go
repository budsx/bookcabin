package repoiface

import (
	"context"
	"io"

	"github.com/budsx/bookcabin/models"
)

type DBReadWriter interface {
	ReadAircraftsByCode(ctx context.Context, code string) (models.Aircraft, error)
	ReadCabinsByAircraftID(ctx context.Context, aircraftID int64) ([]models.Cabin, error)
	ReadSeatColumnsByCabinIDs(ctx context.Context, cabinIDs []int64) ([]models.SeatColumn, error)
	ReadSeatRowsByCabinIDs(ctx context.Context, cabinIDs []int64) ([]models.SeatRow, error)
	ReadSeatsBySeatRowIDs(ctx context.Context, seatRowIDs []int64) ([]models.Seat, error)
	ReadSeatCharacteristicsBySeatIDs(ctx context.Context, seatIDs []int64) ([]models.SeatCharacteristic, error)
	ReadRawSeatCharacteristicsBySeatIDs(ctx context.Context, seatIDs []int64) ([]models.RawSeatCharacteristic, error)
	ReadSeatCodesBySeatRowIDs(ctx context.Context, seatRowIDs []int64) ([]string, error)
	ReadPassengerByID(ctx context.Context, id int64) (models.Passenger, error)
	ReadBookingFlightByID(ctx context.Context, flightNumber int64) (models.BookingFlight, error)
	ReadBookingByID(ctx context.Context, bookingID int64) (models.Booking, error) 
	Ping(ctx context.Context) error
	io.Closer
}
