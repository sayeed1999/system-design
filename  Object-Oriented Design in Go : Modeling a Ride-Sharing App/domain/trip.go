package domain

import "time"

type Trip struct {
	ID              string
	RiderID         string
	Rider           Rider // Trip HAS-A Rider (composition)
	PassengerID     string
	Passenger       Passenger // Trip HAS-A Passenger (composition)
	PickupLocation  Location
	DropoffLocation Location
	StartTime       time.Time
	EndTime         time.Time
	FareAmount      int
}

func NewTrip(TripRequest TripRequest, Rider Rider) *Trip {
	return &Trip{
		ID: "", // generate a unique ID here
		// take driver info from the assigned driver
		RiderID: Rider.ID,
		Rider:   Rider,
		// take rest of the info from the trip request
		PassengerID:     TripRequest.PassengerID,
		Passenger:       TripRequest.Passenger,
		PickupLocation:  TripRequest.PickupLocation,
		DropoffLocation: TripRequest.DropoffLocation,
		FareAmount:      TripRequest.FareAmount,
		StartTime:       time.Now().UTC(),
		EndTime:         time.Time{},
	}
}
