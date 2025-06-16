package repoiface

import (
	"context"
	"io"

	"github.com/budsx/bookcabin/models"
)

type DBReadWriter interface {
	ReadAircraftsByCode(ctx context.Context, code string) (models.Aircraft, error)
	ReadCabinsByAircraftID(ctx context.Context, aircraftID int32) ([]models.Cabin, error)
	ReadSeatColumnsByCabinIDs(ctx context.Context, cabinIDs []int32) ([]models.SeatColumn, error)
	ReadSeatRowsByCabinIDs(ctx context.Context, cabinIDs []int32) ([]models.SeatRow, error)
	ReadSeatsBySeatRowIDs(ctx context.Context, seatRowIDs []int32) ([]models.Seat, error)
	ReadSeatCharacteristicsBySeatIDs(ctx context.Context, seatIDs []int32) ([]models.SeatCharacteristic, error)
	ReadRawSeatCharacteristicsBySeatIDs(ctx context.Context, seatIDs []int32) ([]models.RawSeatCharacteristic, error)
	ReadSeatCodesBySeatRowIDs(ctx context.Context, seatRowIDs []int32) ([]string, error)

	ReadPassengerByID(ctx context.Context, id int32) (models.Passenger, error)

	Ping(ctx context.Context) error
	io.Closer
}
