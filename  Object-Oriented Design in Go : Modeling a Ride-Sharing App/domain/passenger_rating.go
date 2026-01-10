package domain

// PassengerRating represents a rating given by a Rider to a Passenger
type PassengerRating struct {
	BaseRating // PassengerRating IS A BaseRating (inheritance via embedding)
}
