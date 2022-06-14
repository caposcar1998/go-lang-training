package main

type City struct {
	latitude  float64
	longitude float64
}

func CreateCity(lat float64, lon float64) *City {

	if lat < 0.0 || lon < 0.0 {
		return nil
	}

	return &City{
		latitude:  lat,
		longitude: lon,
	}

}

func (c *City) RetrieveLatitudeCity() float64 {
	return c.latitude
}

func (c *City) RetrieveLongitudeCity() float64 {
	return c.longitude
}
