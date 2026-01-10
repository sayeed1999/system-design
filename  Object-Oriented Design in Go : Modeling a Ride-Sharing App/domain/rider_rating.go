package domain

// RiderRating represents a rating given by a Passenger to a Rider
type RiderRating struct {
	BaseRating // RiderRating IS A BaseRating (inheritance via embedding)
}
