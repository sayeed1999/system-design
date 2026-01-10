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
		"DL1234567890");

	fmt.Printf("Rider: %+v\n", rider)

	passenger := domain.NewPassenger(
		"passenger1",
		"Jane Smith",
		"+0987654321",
		"jane.smith@example.com",
		domain.Location{LatitudeX: 37.7749, LongitudeY: -122.4194},
		"Credit Card");

	fmt.Printf("Passenger: %+v\n", passenger)
}