package domain

type Passenger struct {
	Person // Passenger IS A Person (inheritance via embedding)

	// passenger specific fields
	FavoritePaymentMethod string
	Rating                float32
	PassengerRatings      []PassengerRating
}

func NewPassenger(id string, name string, phone string, email string, currentLocation Location, favoritePaymentMethod string) *Passenger {
	return &Passenger{
		Person: Person{
			ID:              id,
			Name:            name,
			Phone:           phone,
			Email:           email,
			CurrentLocation: currentLocation,
		},
		FavoritePaymentMethod: favoritePaymentMethod,
		Rating:                0.0,
		PassengerRatings:      []PassengerRating{},
	}
}

// From DDD perspective, this is a domain behavior of Passenger aggregate root!
// Passenger will receive rating from Rider after trip completion and modify its own state and its collection of PassengerRatings
func (p *Passenger) ReceiveRating(tripID string, riderID string, ratingValue float32, comment string) {
	// first create the rating object
	rating := NewPassengerRating(tripID, p.ID, riderID, ratingValue, comment)
	// append to Passenger's ratings
	p.PassengerRatings = append(p.PassengerRatings, *rating)

	// Re-calculate overall rating
	var total float32
	for _, r := range p.PassengerRatings {
		total += r.Rating
	}
	p.Rating = total / float32(len(p.PassengerRatings))
}
