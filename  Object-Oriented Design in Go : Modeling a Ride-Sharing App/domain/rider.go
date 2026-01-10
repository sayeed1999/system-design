package domain

type Rider struct {
	Person // Rider IS A Person

	// rider specific fields
	VehicleRegistrationNo string
	DrivingLicenseNo string
}

func NewRider(id string, name string, phone string, email string, currentLocation Location, vehicleRegNo string, licenseNo string) *Rider {
	return &Rider{
		Person: Person{
			ID: id,
			Name: name,
			Phone: phone,
			Email: email,
			CurrentLocation: currentLocation,
		},
		VehicleRegistrationNo: vehicleRegNo,
		DrivingLicenseNo: licenseNo,
	}
}
