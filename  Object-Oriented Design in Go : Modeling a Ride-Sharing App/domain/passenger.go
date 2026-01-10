package domain

type Passenger struct {
	Person // Passenger IS A Person

	// passenger specific fields
	FavoritePaymentMethod string
}

func NewPassenger(id string, name string, phone string, email string, currentLocation Location, favoritePaymentMethod string) *Passenger {
	return &Passenger{
		Person: Person{
			ID: id,
			Name: name,
			Phone: phone,
			Email: email,
			CurrentLocation: currentLocation,
		},
		FavoritePaymentMethod: favoritePaymentMethod,
	}
}
