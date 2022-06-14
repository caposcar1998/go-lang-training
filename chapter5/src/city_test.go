package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var res City

func TestCity_CreateCity(t *testing.T) {
	city := CreateCity(51.45, 1.15)
	if city == nil {
		t.Errorf("City is not created correctly")
	}
}

func TestCity_RetrieveLatitude(t *testing.T) {
	city := CreateCity(51.45, 1.15)
	assert.Equal(t, city.RetrieveLatitudeCity(), 51.45)
}

func TestCity_RetrieveLongitude(t *testing.T) {
	city := CreateCity(51.45, 1.15)
	assert.Equal(t, city.RetrieveLongitudeCity(), 1.15)
}

func Benchmark_CreateCity(b *testing.B) {

	var city City

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CreateCity(51.45, 1.15)
	}
	res = city
}
