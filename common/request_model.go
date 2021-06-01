package common

type ReqBookings struct {
	QuoteId    string `json:"quote_id"`
	Passengers `json:"passengers"`
}
type Passengers struct {
	PassengerDetails     []PassengerDetail `json:"passenger_details"`
	Luggage              `json:"luggage"`
	AdditionalPassengers int64 `json:"additional_passengers"`
}
type Luggage struct {
	Total int64 `json:"total"`
}
type PassengerDetail struct {
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}
