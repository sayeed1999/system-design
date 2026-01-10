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
