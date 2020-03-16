package day01

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tosan88/advent-of-code/aoc_io"
	"testing"
)

func TestCalculateFuel(t *testing.T) {
	tests := []struct {
		name string
		mass int
		fuel int
	}{
		{
			name: "12 mass should return 2",
			mass: 12,
			fuel: 2,
		},
		{
			name: "14 mass should return 2",
			mass: 14,
			fuel: 2,
		},
		{
			name: "1969 mass should return 654",
			mass: 1969,
			fuel: 654,
		},
		{
			name: "100756 mass should return 33583",
			mass: 100756,
			fuel: 33583,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.fuel, calculateFuel(test.mass))
		})
	}
}

func TestCalculateFuelWithCorrection(t *testing.T) {
	tests := []struct {
		name string
		mass int
		fuel int
	}{
		{
			name: "14 mass should return 2",
			mass: 14,
			fuel: 2,
		},
		{
			name: "1969 mass should return 966",
			mass: 1969,
			fuel: 966,
		},
		{
			name: "100756 mass should return 50346",
			mass: 100756,
			fuel: 50346,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.fuel, calculateFuelWithCorrection(test.mass))
		})
	}
}

func TestRunCodeAoCDay01(t *testing.T) {
	masses, err := aoc_io.ReadInputAsIntPerLines("aoc_input01.txt")
	if err != nil {
		errMsg := fmt.Sprintf("Error: %v\n", err)
		t.Fatalf(errMsg)
	}

	allFuel := calculateAllFuel(masses)
	t.Logf("All fuel is: %v\n", allFuel)
	assert.Equal(t, 3239890, allFuel)
	allFuel = calculateAllFuelWithCorrection(masses)
	t.Logf("All fuel with correction is: %v\n", allFuel)
	assert.Equal(t, 4856963, allFuel)
}
