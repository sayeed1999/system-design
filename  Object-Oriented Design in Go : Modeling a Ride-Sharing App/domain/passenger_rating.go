package domain

// PassengerRating represents a rating given by a Rider to a Passenger
type PassengerRating struct {
	BaseRating // PassengerRating IS A BaseRating (inheritance via embedding)
}

func NewPassengerRating(tripID string, passengerID string, riderID string, ratingValue float32, comment string) *PassengerRating {
	return &PassengerRating{
		BaseRating: BaseRating{
			ID:          "", // generate unique ID
			TripID:      tripID,
			RiderID:     riderID,
			PassengerID: passengerID,
			Rating:      ratingValue,
			Comment:     comment,
		},
	}
}
