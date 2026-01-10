# Object-Oriented Design in Go: Modeling a Ride Sharing App

## Define Core Entities

First we define the core entities.

We start with Driver and Passenger: the two must entity for a ride sharing app like Pathao/Uber.

```go
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
