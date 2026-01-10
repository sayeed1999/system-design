package domain

// RiderRating represents a rating given by a Passenger to a Rider
type RiderRating struct {
	BaseRating // RiderRating IS A BaseRating (inheritance via embedding)
}

func NewRiderRating(trip Trip, ratingValue float32, comment string) *RiderRating {
	return &RiderRating{
		BaseRating: BaseRating{
			ID:          "", // generate unique ID
			RiderID:     trip.Rider.ID,
			PassengerID: trip.Passenger.ID,
			Rating:      ratingValue,
			Comment:     comment,
		},
	}

	// TODO: where to update Rider's overall rating???
}
