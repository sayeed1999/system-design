package domain

// PassengerRating represents a rating given by a Rider to a Passenger
type PassengerRating struct {
	ID         string
	DriverID   string
	CustomerID string
	Rating     float32
	Comment    string
}
