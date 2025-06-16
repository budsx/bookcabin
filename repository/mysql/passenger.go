package mysql

import (
	"context"

	"github.com/budsx/bookcabin/models"
)

const (
	selectPassengerByID = `SELECT id, passenger_index, passenger_name_number, first_name, last_name, 
	date_of_birth, gender, "type", street1, street2, postcode, state, city, country, address_type, 
	issuing_country, country_of_birth, document_type, nationality, created_at, updated_at 
	FROM passengers WHERE id=?`

)

func (d *dbReadWriter) ReadPassengerByID(ctx context.Context, id int32) (models.Passenger, error) {
	passenger := models.Passenger{}

	err := d.db.QueryRowContext(ctx, selectPassengerByID, id).Scan(
		&passenger.ID,
		&passenger.PassengerIndex,
		&passenger.PassengerNameNumber,
		&passenger.FirstName,
		&passenger.LastName,
		&passenger.DateOfBirth,
		&passenger.Gender,
		&passenger.Type,
		&passenger.Street1,
		&passenger.Street2,
		&passenger.Postcode,
		&passenger.State,
		&passenger.City,
		&passenger.Country,
		&passenger.AddressType,
		&passenger.IssuingCountry,
		&passenger.CountryOfBirth,
		&passenger.DocumentType,
		&passenger.Nationality,
		&passenger.CreatedAt,
		&passenger.UpdatedAt,
	)

	return passenger, err
}

