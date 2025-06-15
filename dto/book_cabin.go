package dto

type HealthCheckResponse struct {
	Status string `json:"status"`
}

type SeatMapResponse struct {
	SeatsItineraryParts []SeatsItineraryPart `json:"seatsItineraryParts"`
	SelectedSeats       []SelectedSeat       `json:"selectedSeats,omitempty"`
}

type SeatsItineraryPart struct {
	SegmentSeatMaps []SegmentSeatMap `json:"segmentSeatMaps"`
}

type SegmentSeatMap struct {
	PassengerSeatMaps []PassengerSeatMap `json:"passengerSeatMaps"`
	Segment           Segment            `json:"segment"`
}

type PassengerSeatMap struct {
	SeatSelectionEnabledForPax bool      `json:"seatSelectionEnabledForPax"`
	SeatMap                    SeatMap   `json:"seatMap"`
	Passenger                  Passenger `json:"passenger"`
}

type SeatMap struct {
	RowsDisabledCauses []string `json:"rowsDisabledCauses"`
	AirCraft           string   `json:"aircraft"`
	Cabins             []Cabin  `json:"cabins"`
}

type Cabin struct {
	Deck        string    `json:"deck"`
	SeatColumns []string  `json:"seatColumns"`
	SeatRows    []SeatRow `json:"seatRows"`
	FirstRow    int32     `json:"firstRow"`
	LastRow     int32     `json:"lastRow"`
}

type SeatRow struct {
	RowNumber int32    `json:"rowNumber"`
	SeatCodes []string `json:"seatCodes"`
	Seats     []Seat   `json:"seats"`
}

type Seat struct {
	StorefrontSlotCode     string    `json:"storefrontSlotCode"`
	Available              bool      `json:"available"`
	Code                   string    `json:"code,omitempty"`
	Designations           []string  `json:"designations,omitempty"`
	Entitled               bool      `json:"entitled"`
	FeeWaived              bool      `json:"feeWaived"`
	EntitledRuleID         string    `json:"entitledRuleId,omitempty"`
	FeeWaivedRuleID        string    `json:"feeWaivedRuleId,omitempty"`
	SeatCharacteristics    []string  `json:"seatCharacteristics,omitempty"`
	Limitations            []string  `json:"limitations,omitempty"`
	RefundIndicator        string    `json:"refundIndicator,omitempty"`
	FreeOfCharge           bool      `json:"freeOfCharge"`
	Prices                 PriceInfo `json:"prices,omitempty"`
	Taxes                  PriceInfo `json:"taxes,omitempty"`
	Total                  PriceInfo `json:"total,omitempty"`
	OriginallySelected     bool      `json:"originallySelected"`
	RawSeatCharacteristics []string  `json:"rawSeatCharacteristics,omitempty"`
}

type PriceInfo struct {
	Alternatives [][]Price `json:"alternatives"`
}

type Price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type Passenger struct {
	PassengerIndex      int              `json:"passengerIndex"`
	PassengerNameNumber string           `json:"passengerNameNumber"`
	PassengerDetails    PassengerDetails `json:"passengerDetails"`
	PassengerInfo       PassengerInfo    `json:"passengerInfo"`
	Preferences         Preferences      `json:"preferences"`
	DocumentInfo        DocumentInfo     `json:"documentInfo"`
}

type PassengerDetails struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type PassengerInfo struct {
	DateOfBirth string   `json:"dateOfBirth"`
	Gender      string   `json:"gender"`
	Type        string   `json:"type"`
	Emails      []string `json:"emails"`
	Phones      []string `json:"phones"`
	Address     Address  `json:"address"`
}

type Address struct {
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	Postcode    string `json:"postcode"`
	State       string `json:"state"`
	City        string `json:"city"`
	Country     string `json:"country"`
	AddressType string `json:"addressType"`
}

type Preferences struct {
	SpecialPreferences SpecialPreferences `json:"specialPreferences"`
	FrequentFlyer      []FrequentFlyer    `json:"frequentFlyer"`
}

type SpecialPreferences struct {
	MealPreference               string   `json:"mealPreference"`
	SeatPreference               string   `json:"seatPreference"`
	SpecialRequests              []string `json:"specialRequests"`
	SpecialServiceRequestRemarks []string `json:"specialServiceRequestRemarks"`
}

type FrequentFlyer struct {
	Airline    string `json:"airline"`
	Number     string `json:"number"`
	TierNumber int    `json:"tierNumber"`
}

type DocumentInfo struct {
	IssuingCountry string `json:"issuingCountry"`
	CountryOfBirth string `json:"countryOfBirth"`
	DocumentType   string `json:"documentType"`
	Nationality    string `json:"nationality"`
}

type Segment struct {
	Type                        string                  `json:"@type"`
	SegmentOfferInformation     SegmentOfferInformation `json:"segmentOfferInformation"`
	Duration                    int                     `json:"duration"`
	CabinClass                  string                  `json:"cabinClass"`
	Equipment                   string                  `json:"equipment"`
	Flight                      Flight                  `json:"flight"`
	Origin                      string                  `json:"origin"`
	Destination                 string                  `json:"destination"`
	Departure                   string                  `json:"departure"`
	Arrival                     string                  `json:"arrival"`
	BookingClass                string                  `json:"bookingClass"`
	LayoverDuration             int                     `json:"layoverDuration"`
	FareBasis                   string                  `json:"fareBasis"`
	SubjectToGovernmentApproval bool                    `json:"subjectToGovernmentApproval"`
	SegmentRef                  string                  `json:"segmentRef"`
}

type SegmentOfferInformation struct {
	FlightsMiles int  `json:"flightsMiles"`
	AwardFare    bool `json:"awardFare"`
}

type Flight struct {
	FlightNumber          int      `json:"flightNumber"`
	OperatingFlightNumber int      `json:"operatingFlightNumber"`
	AirlineCode           string   `json:"airlineCode"`
	OperatingAirlineCode  string   `json:"operatingAirlineCode"`
	StopAirports          []string `json:"stopAirports"`
	DepartureTerminal     string   `json:"departureTerminal"`
	ArrivalTerminal       string   `json:"arrivalTerminal"`
}

type SelectedSeat struct {
}
