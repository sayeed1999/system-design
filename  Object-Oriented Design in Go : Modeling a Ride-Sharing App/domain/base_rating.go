package domain

type BaseRating struct {
	ID          string
	TripID      string
	RiderID     string
	PassengerID string
	Rating      float32
	Comment     string
}
