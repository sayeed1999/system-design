package domain

// PassengerRating represents a rating given by a Rider to a Passenger
type PassengerRating struct {
	BaseRating // PassengerRating IS A BaseRating (inheritance via embedding)
}

func NewPassengerRating(trip Trip, ratingValue float32, comment string) *PassengerRating {
	return &PassengerRating{
		BaseRating: BaseRating{
			ID:          "", // generate unique ID
			RiderID:     trip.Rider.ID,
			PassengerID: trip.Passenger.ID,
			Rating:      ratingValue,
			Comment:     comment,
		},
	}
}
