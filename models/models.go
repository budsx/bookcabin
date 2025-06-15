package models

import "time"

type Aircraft struct {
	ID        int64
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Cabin struct {
	ID         int64
	AircraftID int32
	Deck       string
	FirstRow   int32
	LastRow    int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type SeatColumn struct {
	ID         int64
	CabinID    int32
	ColumnCode string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type SeatRow struct {
	ID        int64
	CabinID   int32
	RowNumber int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Seat struct {
	ID                 int64
	SeatRowID          int32
	Code               string
	StorefrontSlotCode string
	RefundIndicator    string
	FreeOfCharge       bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type SeatCharacteristic struct {
	ID             int64
	SeatID         int32
	Characteristic string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type RawSeatCharacteristic struct {
	ID                int64
	SeatID            int32
	RawCharacteristic string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Passenger struct {
	ID                  int64
	PassengerIndex      int32
	PassengerNameNumber string
	FirstName           string
	LastName            string
	DateOfBirth         string
	Gender              string
	Type                string
	Street1             string
	Street2             string
	Postcode            string
	State               string
	City                string
	Country             string
	AddressType         string
	IssuingCountry      string
	CountryOfBirth      string
	DocumentType        string
	Nationality         string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type PassengerEmail struct {
	ID          int64
	PassengerID int32
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PassengerPhone struct {
	ID          int64
	PassengerID int32
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type FrequentFlyer struct {
	ID          int64
	PassengerID int32
	Airline     string
	Number      string
	TierNumber  int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SpecialPreference struct {
	ID             int64
	PassengerID    int32
	MealPreference string
	SeatPreference string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type SpecialRequest struct {
	ID           int64
	PreferenceID int32
	Request      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SpecialServiceRequestRemark struct {
	ID           int64
	PreferenceID int32
	Remark       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Booking struct {
	ID                          int64
	BookingReference            string
	Origin                      string
	Destination                 string
	Departure                   string
	Arrival                     string
	Equipment                   string
	FareBasis                   string
	BookingClass                string
	CabinClass                  string
	Duration                    int32
	LayoverDuration             int32
	SegmentRef                  string
	SubjectToGovernmentApproval bool
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
}

type BookingFlight struct {
	ID                    int64
	BookingID             int32
	FlightNumber          int32
	OperatingFlightNumber int32
	AirlineCode           string
	OperatingAirlineCode  string
	DepartureTerminal     string
	ArrivalTerminal       string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type BookingFlightStopAirport struct {
	ID          int64
	FlightID    int32
	AirportCode string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookingSeat struct {
	ID                 int64
	BookingID          int32
	PassengerID        int32
	SeatID             int32
	Available          bool
	Entitled           bool
	FeeWaived          bool
	EntitledRuleID     string
	FeeWaivedRuleID    string
	PriceAmount        float64
	PriceCurrency      string
	TaxAmount          float64
	TaxCurrency        string
	TotalAmount        float64
	TotalCurrency      string
	OriginallySelected bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
