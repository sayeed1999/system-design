package domain

type Payment struct {
	ID            string
	TripID        string
	Amount        int
	PaymentMethod string
	Status        string
}

func NewPayment(Trip Trip, PaymentMethod string) *Payment {
	return &Payment{
		TripID:        Trip.ID,
		Amount:        Trip.FareAmount,
		PaymentMethod: PaymentMethod,
		Status:        "Pending",
	}
}
