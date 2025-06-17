package mysql

import (
	"context"
	"database/sql"

	"github.com/budsx/bookcabin/models"
)

const (
	selectPassengerByID = `SELECT id, passenger_index, passenger_name_number, first_name, last_name, 
	date_of_birth, gender, "type", street1, street2, postcode, state, city, country, address_type, 
	issuing_country, country_of_birth, document_type, nationality, created_at, updated_at 
	FROM passengers WHERE id = ?`

	selectPassengerEmailByID = `SELECT email FROM passenger_emails WHERE passenger_id = ?`

	selectPassengerPhoneByID = `SELECT phone_number FROM passenger_phones WHERE passenger_id = ?`
)

func (d *dbReadWriter) ReadPassengerByID(ctx context.Context, id int64) (models.Passenger, error) {
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

func (d *dbReadWriter) ReadPassengerEmail(ctx context.Context, passengerID int64) ([]string, error) {
	var emails []string

	rows, err := d.db.QueryContext(ctx, selectPassengerEmailByID, passengerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		emails = append(emails, email)
	}

	return emails, nil
}

func (d *dbReadWriter) ReadPassengerPhone(ctx context.Context, passengerID int64) ([]string, error) {
	var phones []string

	rows, err := d.db.QueryContext(ctx, selectPassengerPhoneByID, passengerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var phone string
		err := rows.Scan(&phone)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		phones = append(phones, phone)
	}

	return phones, nil
}
