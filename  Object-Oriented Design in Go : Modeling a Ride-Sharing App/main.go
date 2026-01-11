package main

import "fmt"
import "ood-tutorial-ride-sharing-app/domain"

func main() {
	fmt.Println("Object-Oriented Design in Go: Modeling a Ride-Sharing App")

	rider := domain.NewRider(
		"rider1",
		"John Doe",
		"+1234567890",
		"john.doe@example.com",
		*domain.NewLocation(37.7749, -122.4194),
		"ABC123",
		"DL1234567890")

	fmt.Printf("Rider: %+v\n\n", rider)

	passenger := domain.NewPassenger(
		"passenger1",
		"Jane Smith",
		"+0987654321",
		"jane.smith@example.com",
		*domain.NewLocation(37.7749, -122.4194),
		"Credit Card")

	fmt.Printf("Passenger: %+v\n\n", passenger)

	tripRequest := domain.NewTripRequest(
		*passenger,
		domain.Location{LatitudeX: 37.7849, LongitudeY: -122.4094},
		200,
	)

	fmt.Printf("Trip Request: %+v\n\n", tripRequest)

	trip := domain.NewTrip(*tripRequest, *rider)
	fmt.Printf("Trip: %+v\n\n", trip)

	payment := domain.NewPayment(*trip, "Cash On Delivery")
	fmt.Printf("Payment: %+v\n\n", payment)

	// rider receives rating from passenger after trip completion
	rider.ReceiveRating(trip.ID, trip.PassengerID, 4.5, "Great ride!")
	fmt.Printf("Updated Rider after Rating: %+v\n\n", rider)

	// passenger receives rating from rider after trip completion
	passenger.ReceiveRating(trip.ID, trip.RiderID, 5.0, "Excellent passenger!")
	fmt.Printf("Updated Passenger after Rating: %+v\n\n", passenger)

	// TODO: Use composition to extract common behaviors of Receive Rating between Rider and Passenger
}
