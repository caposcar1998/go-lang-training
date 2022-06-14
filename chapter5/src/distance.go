package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/umahmood/haversine"
)

const radius = 6371

func CalculateHarvesineDistance(cityOne *City, cityTwo *City) (float64, float64) {
	cityO := haversine.Coord{Lat: cityOne.latitude, Lon: cityOne.longitude} // Oxford, UK
	cityT := haversine.Coord{Lat: cityTwo.latitude, Lon: cityOne.longitude} // Turin, Italy
	mi, km := haversine.Distance(cityO, cityT)
	fmt.Println("Miles:", mi, "Kilometers:", km)
	return mi, km

}

func degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func (origin *City) CalculateHarvesineDistanceManual(destination *City) (float64, error) {

	degreesLat := degrees2radians(destination.latitude - origin.latitude)
	degreesLong := degrees2radians(destination.longitude - origin.longitude)
	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(degrees2radians(origin.latitude))*
			math.Cos(degrees2radians(destination.latitude))*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := radius * c
	fmt.Println("km manual: ", d)
	if d > 0 {
		return d, nil
	} else {
		return d, errors.New("The value cant be zero")
	}

}
