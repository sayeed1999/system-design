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
		domain.Location{LatitudeX: 37.7749, LongitudeY: -122.4194},
		"ABC123",
		"DL1234567890")

	fmt.Printf("Rider: %+v\n\n", rider)

	passenger := domain.NewPassenger(
		"passenger1",
		"Jane Smith",
		"+0987654321",
		"jane.smith@example.com",
		domain.Location{LatitudeX: 37.7749, LongitudeY: -122.4194},
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
}
