package mysql

import (
	"bookcabin/models"
	"context"
)

const (
	selectAircraftsByCode = `
		SELECT id, code, created_at, updated_at FROM aircrafts WHERE code = ?
	`
	selectCabinsByAircraftID = `
		SELECT id, aircraft_id, deck, first_row, last_row, created_at, updated_at FROM cabins WHERE aircraft_id = ?
	`
	selectSeatColumnsByCabinID = `
		SELECT id, cabin_id, column_code, created_at, updated_at FROM seat_columns WHERE cabin_id = ?
	`
	selectSeatRowsByCabinID = "SELECT id, cabin_id, `row_number`, created_at, updated_at FROM seat_rows WHERE cabin_id = ?"
	
	selectSeatsBySeatRowID = `
		SELECT id, seat_row_id, code, storefront_slot_code, refund_indicator, free_of_charge, created_at, updated_at FROM seats WHERE seat_row_id = ?
	`
	selectSeatCharacteristicsBySeatID = `
		SELECT id, seat_id, characteristic, created_at, updated_at FROM seat_characteristics WHERE seat_id = ?
	`
	selectRawSeatCharacteristicsBySeatID = `
		SELECT id, seat_id, raw_characteristic, created_at, updated_at FROM raw_seat_characteristics WHERE seat_id = ?
	`
)

func (d *dbReadWriter) ReadAircraftsByCode(ctx context.Context, code string) (models.Aircraft, error) {
	row := d.db.QueryRowContext(ctx, selectAircraftsByCode, code)

	var aircraft models.Aircraft
	err := row.Scan(&aircraft.ID, &aircraft.Code, &aircraft.CreatedAt, &aircraft.UpdatedAt)
	if err != nil {
		return models.Aircraft{}, err
	}
	return aircraft, nil
}

func (d *dbReadWriter) ReadCabinsByAircraftID(ctx context.Context, aircraftID int32) ([]models.Cabin, error) {
	rows, err := d.db.QueryContext(ctx, selectCabinsByAircraftID, aircraftID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cabins []models.Cabin
	for rows.Next() {
		var cabin models.Cabin
		err := rows.Scan(&cabin.ID, &cabin.AircraftID, &cabin.Deck, &cabin.FirstRow, &cabin.LastRow, &cabin.CreatedAt, &cabin.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cabins = append(cabins, cabin)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cabins, nil
}

func (d *dbReadWriter) ReadSeatColumnsByCabinID(ctx context.Context, cabinID int32) ([]models.SeatColumn, error) {
	rows, err := d.db.QueryContext(ctx, selectSeatColumnsByCabinID, cabinID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seatColumns []models.SeatColumn
	for rows.Next() {
		var seatColumn models.SeatColumn
		err := rows.Scan(&seatColumn.ID, &seatColumn.CabinID, &seatColumn.ColumnCode, &seatColumn.CreatedAt, &seatColumn.UpdatedAt)
		if err != nil {
			return nil, err
		}
		seatColumns = append(seatColumns, seatColumn)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return seatColumns, nil
}

func (d *dbReadWriter) ReadSeatRowsByCabinID(ctx context.Context, cabinID int32) ([]models.SeatRow, error) {
	rows, err := d.db.QueryContext(ctx, selectSeatRowsByCabinID, cabinID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seatRows []models.SeatRow
	for rows.Next() {
		var seatRow models.SeatRow
		err := rows.Scan(
			&seatRow.ID, 
			&seatRow.CabinID, 
			&seatRow.RowNumber, 
			&seatRow.CreatedAt, 
			&seatRow.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		seatRows = append(seatRows, seatRow)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return seatRows, nil
}

func (d *dbReadWriter) ReadSeatsBySeatRowID(ctx context.Context, seatRowID int32) ([]models.Seat, error) {
	rows, err := d.db.QueryContext(ctx, selectSeatsBySeatRowID, seatRowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []models.Seat
	for rows.Next() {
		var seat models.Seat
		err := rows.Scan(&seat.ID, &seat.SeatRowID, &seat.Code, &seat.StorefrontSlotCode, &seat.RefundIndicator, &seat.FreeOfCharge, &seat.CreatedAt, &seat.UpdatedAt)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}

func (d *dbReadWriter) ReadSeatCharacteristicsBySeatID(ctx context.Context, seatID int32) ([]models.SeatCharacteristic, error) {
	rows, err := d.db.QueryContext(ctx, selectSeatCharacteristicsBySeatID, seatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seatCharacteristics []models.SeatCharacteristic
	for rows.Next() {
		var seatCharacteristic models.SeatCharacteristic
		err := rows.Scan(&seatCharacteristic.ID, &seatCharacteristic.SeatID, &seatCharacteristic.Characteristic, &seatCharacteristic.CreatedAt, &seatCharacteristic.UpdatedAt)
		if err != nil {
			return nil, err
		}
		seatCharacteristics = append(seatCharacteristics, seatCharacteristic)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return seatCharacteristics, nil
}

func (d *dbReadWriter) ReadRawSeatCharacteristicsBySeatID(ctx context.Context, seatID int32) ([]models.RawSeatCharacteristic, error) {
	rows, err := d.db.QueryContext(ctx, selectRawSeatCharacteristicsBySeatID, seatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rawSeatCharacteristics []models.RawSeatCharacteristic
	for rows.Next() {
		var rawSeatCharacteristic models.RawSeatCharacteristic
		err := rows.Scan(&rawSeatCharacteristic.ID, &rawSeatCharacteristic.SeatID, &rawSeatCharacteristic.RawCharacteristic, &rawSeatCharacteristic.CreatedAt, &rawSeatCharacteristic.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rawSeatCharacteristics = append(rawSeatCharacteristics, rawSeatCharacteristic)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rawSeatCharacteristics, nil
}
