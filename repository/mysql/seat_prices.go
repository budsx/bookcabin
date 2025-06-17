package mysql

import (
	"context"
	"fmt"

	"github.com/budsx/bookcabin/models"
)

const (
	selectSeatPricesBySeatIDs = `
	SELECT id, seat_id, price_type, amount, currency, alternative_group
	FROM seat_prices
	WHERE seat_id IN (%s)
	`
)

func (d *dbReadWriter) ReadSeatPricesBySeatIDs(ctx context.Context, seatIDs []int64) ([]models.SeatPrice, error) {
	inClause := buildInClause(len(seatIDs))
	query := fmt.Sprintf(selectSeatPricesBySeatIDs, inClause)

	args := make([]interface{}, len(seatIDs))
	for i, id := range seatIDs {
		args[i] = id
	}
	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seatPrices []models.SeatPrice
	for rows.Next() {
		var seatPrice models.SeatPrice
		err := rows.Scan(
			&seatPrice.ID,
			&seatPrice.SeatID,
			&seatPrice.PriceType,
			&seatPrice.Amount,
			&seatPrice.Currency,
			&seatPrice.AlternativeGroup,
		)
		if err != nil {
			return nil, err
		}
		seatPrices = append(seatPrices, seatPrice)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return seatPrices, nil
}
