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
