package advent_of_code

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDigits(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output []int
	}{
		{
			name:   "Simple test",
			input:  123,
			output: []int{1, 2, 3},
		},
		{
			name:   "Negative test",
			input:  -123,
			output: []int{1, 2, 3},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, getDigits(test.input))
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output bool
	}{
		{
			name:   "Valid simple input",
			input:  111111,
			output: true,
		},
		{
			name:   "Valid complex input",
			input:  122345,
			output: true,
		},
		{
			name:   "Invalid lacking double",
			input:  123789,
			output: false,
		},
		{
			name:   "Invalid decreasing order",
			input:  223450,
			output: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			 assert.Equal(t, test.output, isValid(test.input))
		})
	}
}

func TestIsValidPart2(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output bool
	}{
		{
			name:   "Invalid triple input, no double",
			input:  111111,
			output: false,
		},
		{
			name:   "Valid complex input",
			input:  122345,
			output: true,
		},
		{
			name:   "Valid input with double and a larger group at the beginning",
			input:  111122,
			output: true,
		},
		{
			name:   "Valid input with double and a larger group at the end",
			input:  112222,
			output: true,
		},
		{
			name:   "Valid input with double and a larger group at the middle",
			input:  122233,
			output: true,
		},
		{
			name:   "Invalid lacking double",
			input:  123789,
			output: false,
		},
		{
			name:   "Invalid decreasing order",
			input:  223450,
			output: false,
		},
		{
			name:   "Invalid larger group",
			input:  123444,
			output: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			 assert.Equal(t, test.output, isValidPart2(test.input))
		})
	}
}

func TestAoCDay04Part1(t *testing.T) {
	counter := 0
	for i := 235741; i < 706948; i++ {
		if isValid(i) {
			counter++
		}
	}
	fmt.Println(counter)
}

func TestAoCDay04Part2(t *testing.T) {
	counter := 0
	for i := 235741; i <= 706948; i++ {
		if isValidPart2(i) {
			counter++
		}
	}
	fmt.Println(counter)
}

