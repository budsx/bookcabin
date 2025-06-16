package mysql

import (
	"context"

	"github.com/budsx/bookcabin/models"
)

func (r *dbReadWriter) ReadBookingFlightByID(ctx context.Context, flightNumber int64) (models.BookingFlight, error) {
	bookingFlight := models.BookingFlight{}

	query := `
		SELECT id, booking_id, flight_number, operating_flight_number, airline_code, operating_airline_code, departure_terminal, arrival_terminal, created_at, updated_at
		FROM booking_flights
		WHERE flight_number = ?;
	`

	err := r.db.QueryRowContext(ctx, query, flightNumber).Scan(
		&bookingFlight.ID,
		&bookingFlight.BookingID,
		&bookingFlight.FlightNumber,
		&bookingFlight.OperatingFlightNumber,
		&bookingFlight.AirlineCode,
		&bookingFlight.OperatingAirlineCode,
		&bookingFlight.DepartureTerminal,
		&bookingFlight.ArrivalTerminal,
		&bookingFlight.CreatedAt,
		&bookingFlight.UpdatedAt,
	)
	if err != nil {
		return bookingFlight, err
	}

	return bookingFlight, nil
}

func (r *dbReadWriter) ReadBookingByID(ctx context.Context, bookingID int64) (models.Booking, error) {
	booking := models.Booking{}

	query := `
		SELECT id, booking_reference, origin, destination, departure, arrival, equipment, fare_basis, booking_class, cabin_class, duration, layover_duration, segment_ref, subject_to_government_approval, created_at, updated_at
		FROM bookings
		WHERE id = ?;
	`

	err := r.db.QueryRowContext(ctx, query, bookingID).Scan(
		&booking.ID,
		&booking.BookingReference,
		&booking.Origin,
		&booking.Destination,
		&booking.Departure,
		&booking.Arrival,
		&booking.Equipment,
		&booking.FareBasis,
		&booking.BookingClass,
		&booking.CabinClass,
		&booking.Duration,
		&booking.LayoverDuration,
		&booking.SegmentRef,
		&booking.SubjectToGovernmentApproval,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)
	if err != nil {
		return booking, err
	}

	return booking, nil
}
