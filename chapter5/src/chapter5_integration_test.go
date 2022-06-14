package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func calculate_distance(t *testing.T) {
	oxford := CreateCity(51.45, 1.15)
	turin := CreateCity(45.04, 7.42)
	ml, km := CalculateHarvesineDistance(oxford, turin)
	assert.Equal(t, 42.80364479912663, ml, "The miles where calculated correctly")
	assert.Equal(t, 712.759479791621, km, "The kilometers where calculated correctly")
}
