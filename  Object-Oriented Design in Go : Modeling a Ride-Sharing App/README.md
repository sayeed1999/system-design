# Object-Oriented Design in Go: Modeling a Ride Sharing App

## Define Core Entities: Rider, Passenger

First we define the core entities.

We start with Rider and Passenger: the two must entity for a ride sharing app like Pathao/Uber.

```go
package domain

type Rider struct {
	ID   string
	Name string
	Phone string
	Email string
	CurrentLocation string // TODO: split lat, lon later

	// rider specific fields
	VehicleRegistrationNo string
	DrivingLicenseNo string
}

type Passenger struct {
	ID              string
	Name            string
	Phone           string
	Email           string
	CurrentLocation string // TODO: split lat, lon later

	// passenger specific fields
	FavoritePaymentMethod string
}
```

Now our primary goal is to follow DRY principle: - DON'T REPEAT YOURSELF!
We want to centralize the common behavior within the two entities.

So we make a base entity for both rider & passenger, say `Person`.

```go
package domain

type Person struct {
	ID              string
	Name            string
	Phone           string
	Email           string
	CurrentLocation string // TODO: split lat, lon later
}
```

Here is the truth: Go doesn't support classical inheritance like other OOP languages.
So to mimic inheritance, we embed the base struct in child structs like composition.

```go
package domain

type Rider struct {
	Person // Rider IS A Person

	// rider specific fields
	VehicleRegistrationNo string
	DrivingLicenseNo string
}

type Passenger struct {
	Person // Passenger IS A Person

	// passenger specific fields
	FavoritePaymentMethod string
}
```

Since this is not identical to classical inheritance, we cannot assign `rider.Name` directly, instead in golang we do `rider.Person.Name`. In constructor, we do like this -

```go
package domain

func NewRider(id string, name string, phone string, email string, currentLocation string, vehicleRegNo string, licenseNo string) *Rider {
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

func NewPassenger(id string, name string, phone string, email string, currentLocation string, favoritePaymentMethod string) *Rider {
	return &Passenger{
		Person: Person{
			ID: id,
			Name: name,
			Phone: phone,
			Email: email,
			CurrentLocation: currentLocation,
		},
		FavoritePaymentMethod: favoritePaymentMethod,
	}
}
```

**What we learnt -**
In golang, inheritance is done in a composition way.

Now the next enhancement is the `CurrentLocation`. Previously we kept it as a string, but actually location is real life is a combination of two points: `Latitude (X)` and `Longitude (Y)`.

We can either create two fields LatX and LonY inside Person struct, or we can encapsulate them into a separate struct `Location` and **embed** in `Person` struct.

```go
package domain

type Location struct {
	LatitudeX float64
	LongitudeY float64
}
```

Now the constructor `NewRider` and `NewPassenger` remains as before except `currentLocation` changes from `string` -> `Location` struct.

And the main function becomes: -

```go
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
```

## Extend Core Entities: Rider, Passenger -> TripRequest -> Trip -> Payment

Now that we are done with the two pillars Rider & Passenger, we will move to the business logic.

- When a passenger needs a ride, he creates a request for a ride, say `TripRequest`
- Once a rider accepts the request, then it becomes a trip (the journery from A -> B), say `Trip`
- Once the ride is finished, we need a payment for the ride, say `Payment`.

So this is the chain of entities needed: trip_request -> trip -> payment.

We assume a `trip_request` will start from the passenger's current location. So we didn't take `pickup_location` in `NewTripRequest(...)` params. <i>Later, you can ofcourse enhance a passenger's ability to choose a different pickup location.</i>

```go
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
```

You should notice - there is no rider info in a trip_request. Because there can be a trip_request, but no driver found.

But once a driver wants to accept a request, it shall convert to a trip which requires both parties presence!

```go
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
```

This time you notice a design decision. We could have taken 5-8 params in constructor, instead we took the whole `TripRequest` and `Rider` entity in constructor. It helps in better maintainability.


