package repoiface

import (
	"context"
	"io"

	"bookcabin/models"
)

type DBReadWriter interface {
	ReadAircraftsByCode(ctx context.Context, code string) (models.Aircraft, error)
	ReadCabinsByAircraftID(ctx context.Context, aircraftID int32) ([]models.Cabin, error)
	ReadSeatColumnsByCabinID(ctx context.Context, cabinID int32) ([]models.SeatColumn, error)
	ReadSeatRowsByCabinID(ctx context.Context, cabinID int32) ([]models.SeatRow, error)
	ReadSeatsBySeatRowID(ctx context.Context, seatRowID int32) ([]models.Seat, error)
	ReadSeatCharacteristicsBySeatID(ctx context.Context, seatID int32) ([]models.SeatCharacteristic, error)
	ReadRawSeatCharacteristicsBySeatID(ctx context.Context, seatID int32) ([]models.RawSeatCharacteristic, error)
	Ping(ctx context.Context) error
	io.Closer
}
