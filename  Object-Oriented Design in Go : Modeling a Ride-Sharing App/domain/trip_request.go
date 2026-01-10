package domain

import "time"

type TripRequest struct {
	ID              string
	PassengerID     string
	Passenger       Passenger // TripRequest HAS-A Passenger (composition)
	PickupLocation  Location
	DropoffLocation Location
	RequestTime     time.Time
	FareAmount      int
}

func NewTripRequest(Passenger Passenger, DropoffLocation Location, FareAmount int) *TripRequest {
	return &TripRequest{
		Passenger:       Passenger,
		PickupLocation:  Passenger.CurrentLocation, // a trip must start from the passenger's current location
		DropoffLocation: DropoffLocation,
		RequestTime:     time.Now().UTC(),
		FareAmount:      FareAmount,
	}
}
