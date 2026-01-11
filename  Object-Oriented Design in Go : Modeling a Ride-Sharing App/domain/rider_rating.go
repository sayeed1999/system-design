package domain

// RiderRating represents a rating given by a Passenger to a Rider
type RiderRating struct {
	BaseRating // RiderRating IS A BaseRating (inheritance via embedding)
}

func NewRiderRating(tripID string, riderID string, passengerID string, ratingValue float32, comment string) *RiderRating {
	return &RiderRating{
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
