package advent_of_code

import (
	"github.com/stretchr/stew/slice"
	"math"
)

func getDigits(password int) []int {
	q := int(math.Abs(float64(password)))
	var digits []int
	for q > 0 {
		r := q % 10
		q = q / 10
		digits = append(digits, r)
	}
	return reverse(digits)
}

func reverse(ints []int) []int {
	var reversed []int
	for i:= len(ints)-1; i>=0; i-- {
		reversed = append(reversed, ints[i])
	}
	return reversed
}

func isValid(password int) bool {
	digits := getDigits(password)
	if len(digits) != 6 {
		return false
	}
	var double bool
	var previous int
	for _, d := range digits {
		if previous == d {
			double = true
		}
		if previous > d {
			return false
		}
		previous = d
	}

	return double
}

func isValidPart2(password int) bool {
	digits := getDigits(password)
	if len(digits) != 6 {
		return false
	}
	var tripleNrs []int
	var doubleNrs []int
	prevPrevious := -2
	previous := -1
	for _, d := range digits {
		if previous == d {
			if prevPrevious == previous {
				tripleNrs = append(tripleNrs, d)
			}
			doubleNrs = append(doubleNrs, d)
		}

		if previous > d {
			return false
		}
		prevPrevious = previous
		previous = d
	}
	for _, d := range doubleNrs {
		if !slice.ContainsInt(tripleNrs, d) {
			return true
		}
	}
	return false
}
