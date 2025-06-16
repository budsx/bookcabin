package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/budsx/bookcabin/models"
)

const (
	selectAircraftsByCode = `
		SELECT id, code, created_at, updated_at FROM aircrafts WHERE code = ?
	`
	selectCabinsByAircraftID = `
		SELECT id, aircraft_id, deck, first_row, last_row, created_at, updated_at FROM cabins WHERE aircraft_id = ?
	`
	selectSeatColumnsByCabinIDs = `
		SELECT id, cabin_id, column_code, created_at, updated_at FROM seat_columns WHERE cabin_id IN (%s)
	`
	selectSeatRowsByCabinIDs = "SELECT id, cabin_id, `row_number`, created_at, updated_at FROM seat_rows WHERE cabin_id IN (%s)"

	selectSeatsBySeatRowIDs = `
		SELECT id, seat_row_id, code, storefront_slot_code, refund_indicator, free_of_charge, available, designations, entitled, 
		fee_waived, entitled_rule_id, fee_waived_rule_id, limitations, originally_selected, created_at, updated_at 
		FROM seats WHERE seat_row_id IN (%s)
	`
	selectSeatCharacteristicsBySeatIDs = `
		SELECT id, seat_id, characteristic, created_at, updated_at FROM seat_characteristics WHERE seat_id IN (%s)
	`
	selectRawSeatCharacteristicsBySeatIDs = `
		SELECT id, seat_id, raw_characteristic, created_at, updated_at FROM raw_seat_characteristics WHERE seat_id IN (%s)
	`

	selectSeatCodesBySeatRowID = `
		SELECT DISTINCT storefront_slot_code FROM seats WHERE seat_row_id IN (%s)
	`
)

const (
	seatCode         = "SEAT"
	blankSlotCode    = "BLANK"
	bulkheadSlotCode = "BULKHEAD"
	windowSlotCode   = "WINDOW"
	aisleSlotCode    = "AISLE"
	centerSlotCode   = "CENTER"
	exitSlotCode     = "EXIT"
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

func (d *dbReadWriter) ReadCabinsByAircraftID(ctx context.Context, aircraftID int64) ([]models.Cabin, error) {
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

func (d *dbReadWriter) ReadSeatColumnsByCabinIDs(ctx context.Context, cabinIDs []int64) ([]models.SeatColumn, error) {
	inClause := buildInClause(len(cabinIDs))
	query := fmt.Sprintf(selectSeatColumnsByCabinIDs, inClause)

	args := make([]interface{}, len(cabinIDs))
	for i, id := range cabinIDs {
		args[i] = id
	}

	rows, err := d.db.QueryContext(ctx, query, args...)
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

func (d *dbReadWriter) ReadSeatRowsByCabinIDs(ctx context.Context, cabinIDs []int64) ([]models.SeatRow, error) {
	inClause := buildInClause(len(cabinIDs))
	query := fmt.Sprintf(selectSeatRowsByCabinIDs, inClause)

	args := make([]interface{}, len(cabinIDs))
	for i, id := range cabinIDs {
		args[i] = id
	}
	rows, err := d.db.QueryContext(ctx, query, args...)
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

func (d *dbReadWriter) ReadSeatsBySeatRowIDs(ctx context.Context, seatRowIDs []int64) ([]models.Seat, error) {
	inClause := buildInClause(len(seatRowIDs))
	query := fmt.Sprintf(selectSeatsBySeatRowIDs, inClause)

	args := make([]interface{}, len(seatRowIDs))
	for i, id := range seatRowIDs {
		args[i] = id
	}

	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []models.Seat
	for rows.Next() {
		var seat models.Seat
		err := rows.Scan(
			&seat.ID,
			&seat.SeatRowID,
			&seat.Code,
			&seat.StorefrontSlotCode,
			&seat.RefundIndicator,
			&seat.FreeOfCharge,
			&seat.Available,
			&seat.Designations,
			&seat.Entitled,
			&seat.FeeWaived,
			&seat.EntitledRuleID,
			&seat.FeeWaivedRuleID,
			&seat.Limitations,
			&seat.OriginallySelected,
			&seat.CreatedAt,
			&seat.UpdatedAt,
		)
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

func (d *dbReadWriter) ReadSeatCharacteristicsBySeatIDs(ctx context.Context, seatIDs []int64) ([]models.SeatCharacteristic, error) {
	inClause := buildInClause(len(seatIDs))
	query := fmt.Sprintf(selectSeatCharacteristicsBySeatIDs, inClause)

	args := make([]interface{}, len(seatIDs))
	for i, id := range seatIDs {
		args[i] = id
	}
	rows, err := d.db.QueryContext(ctx, query, args...)
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

func (d *dbReadWriter) ReadRawSeatCharacteristicsBySeatIDs(ctx context.Context, seatIDs []int64) ([]models.RawSeatCharacteristic, error) {
	inClause := buildInClause(len(seatIDs))
	query := fmt.Sprintf(selectRawSeatCharacteristicsBySeatIDs, inClause)

	args := make([]interface{}, len(seatIDs))
	for i, id := range seatIDs {
		args[i] = id
	}
	rows, err := d.db.QueryContext(ctx, query, args...)
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

func (d *dbReadWriter) ReadSeatCodesBySeatRowIDs(ctx context.Context, seatRowIDs []int64) ([]string, error) {
	inClause := buildInClause(len(seatRowIDs))
	query := fmt.Sprintf(selectSeatCodesBySeatRowID, inClause)

	args := make([]interface{}, len(seatRowIDs))
	for i, id := range seatRowIDs {
		args[i] = id
	}
	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seatCodes []string
	for rows.Next() {
		var seatCode string
		err := rows.Scan(&seatCode)
		if err != nil {
			return nil, err
		}
		seatCodes = append(seatCodes, seatCode)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return seatCodes, nil
}

func buildInClause(length int) string {
	placeholders := make([]string, length)
	for i := 0; i < length; i++ {
		placeholders[i] = "?"
	}
	return strings.Join(placeholders, ",")
}
