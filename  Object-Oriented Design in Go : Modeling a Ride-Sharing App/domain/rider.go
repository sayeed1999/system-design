package domain

type Rider struct {
	Person // Rider IS A Person (inheritance via embedding)

	// rider specific fields
	VehicleRegistrationNo string
	DrivingLicenseNo      string
	Rating                float32
	RiderRatings          []RiderRating
}

func NewRider(id string, name string, phone string, email string, currentLocation Location, vehicleRegNo string, licenseNo string) *Rider {
	return &Rider{
		Person: Person{
			ID:              id,
			Name:            name,
			Phone:           phone,
			Email:           email,
			CurrentLocation: currentLocation,
		},
		VehicleRegistrationNo: vehicleRegNo,
		DrivingLicenseNo:      licenseNo,
		Rating:                0.0,
		RiderRatings:          []RiderRating{},
	}
}

// From DDD perspective, this is a domain behavior of Rider aggregate root!
// Rider will receive rating from Passenger after trip completion and modify its own state and its collection of RiderRatings
func (r *Rider) ReceiveRating(tripID string, passengerID string, ratingValue float32, comment string) {
	// first create the rating object
	rating := NewRiderRating(tripID, r.ID, passengerID, ratingValue, comment)

	// append to Rider's ratings
	r.RiderRatings = append(r.RiderRatings, *rating)

	// Re-calculate overall rating
	var total float32
	for _, r := range r.RiderRatings {
		total += r.Rating
	}
	r.Rating = total / float32(len(r.RiderRatings))
}
