package domain

// RiderRating represents a rating given by a Passenger to a Rider
type RiderRating struct {
	ID         string
	DriverID   string
	CustomerID string
	Rating     float32
	Comment    string
}
