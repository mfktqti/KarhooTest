package common

//ResQuotes 询价响应结构体
type ResQuotes struct {
	Id string `json:"id"`
	//Validity of this quote request, in seconds (-1 means it has expired)
	Validity int64 `json:"validity"`
	//
	Status string `json:"status"`
	//uotes[].service_level_agreements

	Quotes []Quote `json:"quotes"`

	//Fleet        Quotefleet `json:"fleet"`
	Availability `json:"availability"`
}

type Availability struct {
	Vehicles AvailabilityVehicle `json:"vehicles"`
}

//Quote Quote
type Quote struct {
	Id         string        `json:"id"`
	QuoteType  string        `json:"quote_type"`
	Source     string        `json:"source"`
	PickUpType string        `json:"pick_up_type"`
	Price      QuotePrice    `json:"price"`
	Fleet      Quotefleet    `json:"fleet"`
	Vehicle    QuotesVehicle `json:"vehicle"`

	//用于验证结果返回值
	ServiceLevelAgreements `json:"service_level_agreements"`
}
type ServiceLevelAgreements struct {
	FreeCancellation `json:"free_cancellation"`
	freeWaitingTime  `json:"free_waiting_time"`
}
type FreeCancellation struct {
	Minutes int64  `json:"minutes"`
	Type    string `json:"type"`
}
type freeWaitingTime struct {
	Minutes int64 `json:"minutes"`
}
type QuotePrice struct {
	CurrencyCode string `json:"currency_code"`
	High         int64  `json:"high"`
	Low          int64  `json:"low"`
	//Object	quotes[].price.net
}
type Quotefleet struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LogoUrl     string `json:"logo_url"`
	//Object	quotes[].fleet.rating
	TermsConditionsUrl string `json:"terms_conditions_url"`

	PhoneNumber string `json:"phone_number"`
	/*A set of capabilities supported by a fleet:

	gps_tracking
	flight_tracking
	train_tracking
	driver_details
	vehicle_details
	The list might be extended in future.
	*/
	Capabilities []string `json:"capabilities"`
}
type AvailabilityVehicle struct {
	Class []string `json:"classes"`
	Type  []string `json:"types"`

	/*
		A set of vehicle's attributes.

		electric - a fully battery electric vehicle,
		hybrid - a mixed hybrid combustion vehicle,
		wheelchair - equipped to handle wheelchairs,
		child-seat - provides a child seat,
		taxi - provides a regulated metered fare,
		executive - describes a premium model of vehicle.
		The list might be extended in future.
	*/
	Tags []string `json:"tags"`
}

type QuotesVehicle struct {
	//Object	quotes[].vehicle.qta
	Class             string `json:"class"`
	Type              string `json:"type"`
	PassengerCapacity int64  `json:"passenger_capacity"`
	LuggageCapacity   int64  `json:"luggage_capacity"`
	/*
		A set of vehicle's attributes.

		electric - a fully battery electric vehicle,
		hybrid - a mixed hybrid combustion vehicle,
		wheelchair - equipped to handle wheelchairs,
		child-seat - provides a child seat,
		taxi - provides a regulated metered fare,
		executive - describes a premium model of vehicle.
		The list might be extended in future.
	*/
	Tags []string   `json:"tags"`
	Qta  VehicleQta `json:"qta"`
}
type VehicleQta struct {
	HighMinutes int64 `json:"high_minutes"`
	LowMinutes  int64 `json:"low_minutes"`
}

//Resbookings 订单结果
type Resbookings struct {
	Id            string `json:"id"`
	BookingsAgent `json:"agent"`
	DateBooked    string `json:"date_booked"`
	DateScheduled string `json:"date_scheduled"`
	DisplayTripId string `json:"display_trip_id"`
	Status        string `json:"status"`

	BookingsDestination `json:"destination"`
	BookingsQuote       `json:"quote"`
}
type BookingsDestination struct {
	DisplayAddress string `json:"display_address"`
	PlaceId        string `json:"place_id"`
	PoiType        string `json:"poi_type"`
	//position
	Timezone string `json:"timezone"`
}
type BookingsAgent struct {
	OrganisationId   string `json:"organisation_id"`
	OrganisationName string `json:"organisation_name"`
	UserId           string `json:"user_id"`
	UserName         string `json:"user_name"`
}
type BookingsQuote struct {
	Currency     string `json:"currency"`
	Source       string `json:"source"`
	Type         string `json:"type"`
	VehicleClass string `json:"vehicle_class"`
	HighPrice    int64  `json:"high_price"`
	LowPrice     int64  `json:"low_price"`
	Total        int64  `json:"total"`
	//vehicle_attributes
}
