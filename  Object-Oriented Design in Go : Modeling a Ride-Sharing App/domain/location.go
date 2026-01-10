package domain

type Location struct {
	LatitudeX  float64
	LongitudeY float64
}

func NewLocation(lat, long float64) *Location {
	return &Location{
		LatitudeX:  lat,
		LongitudeY: long,
	}
}
