package advent_of_code

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConstructPath(t *testing.T) {
	tests := []struct {
		name     string
		wirePath map[wire][]wireDirection
		grid     map[point]pathEntry
	}{
		{
			name: "firstWire: R8,U5,L5,D3",
			wirePath: map[wire][]wireDirection{
				firstWire: {
					{value: "R", times: 8},
					{value: "U", times: 5},
					{value: "L", times: 5},
					{value: "D", times: 3},
				},
			},
			grid: map[point]pathEntry{

				point{1, 0}:  {wire: firstWire},
				point{2, 0}:  {wire: firstWire},
				point{3, 0}:  {wire: firstWire},
				point{4, 0}:  {wire: firstWire},
				point{5, 0}:  {wire: firstWire},
				point{6, 0}:  {wire: firstWire},
				point{7, 0}:  {wire: firstWire},
				point{8, 0}:  {wire: firstWire},
				point{8, -1}: {wire: firstWire},
				point{8, -2}: {wire: firstWire},
				point{8, -3}: {wire: firstWire},
				point{8, -4}: {wire: firstWire},
				point{8, -5}: {wire: firstWire},
				point{7, -5}: {wire: firstWire},
				point{6, -5}: {wire: firstWire},
				point{5, -5}: {wire: firstWire},
				point{4, -5}: {wire: firstWire},
				point{3, -5}: {wire: firstWire},
				point{3, -4}: {wire: firstWire},
				point{3, -3}: {wire: firstWire},
				point{3, -2}: {wire: firstWire},
			},
		},
		{
			name: "secondWire: U7,R6,D4,L4",
			wirePath: map[wire][]wireDirection{
				secondWire: {
					{value: "U", times: 7},
					{value: "R", times: 6},
					{value: "D", times: 4},
					{value: "L", times: 4},
				},
			},
			grid: map[point]pathEntry{

				point{0, -1}: {wire: secondWire},
				point{0, -2}: {wire: secondWire},
				point{0, -3}: {wire: secondWire},
				point{0, -4}: {wire: secondWire},
				point{0, -5}: {wire: secondWire},
				point{0, -6}: {wire: secondWire},
				point{0, -7}: {wire: secondWire},
				point{1, -7}: {wire: secondWire},
				point{2, -7}: {wire: secondWire},
				point{3, -7}: {wire: secondWire},
				point{4, -7}: {wire: secondWire},
				point{5, -7}: {wire: secondWire},
				point{6, -7}: {wire: secondWire},
				point{6, -6}: {wire: secondWire},
				point{6, -5}: {wire: secondWire},
				point{6, -4}: {wire: secondWire},
				point{6, -3}: {wire: secondWire},
				point{5, -3}: {wire: secondWire},
				point{4, -3}: {wire: secondWire},
				point{3, -3}: {wire: secondWire},
				point{2, -3}: {wire: secondWire},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGrid()
			for w, dirs := range test.wirePath {
				g.constructPath(w, dirs)
			}
			assertNonNullMapsEqual(t, test.grid, g.path)
		})
	}
}

func TestConstructAllPath(t *testing.T) {
	tests := []struct {
		name     string
		wirePath map[wire][]wireDirection
		grid     map[point]pathEntry
	}{
		{
			name: "firstWire: R8,U5,L5,D3",
			wirePath: map[wire][]wireDirection{
				firstWire: {
					{value: "R", times: 8},
					{value: "U", times: 5},
					{value: "L", times: 5},
					{value: "D", times: 3},
				},
			},
			grid: map[point]pathEntry{

				point{1, 0}:  {wire: firstWire},
				point{2, 0}:  {wire: firstWire},
				point{3, 0}:  {wire: firstWire},
				point{4, 0}:  {wire: firstWire},
				point{5, 0}:  {wire: firstWire},
				point{6, 0}:  {wire: firstWire},
				point{7, 0}:  {wire: firstWire},
				point{8, 0}:  {wire: firstWire},
				point{8, -1}: {wire: firstWire},
				point{8, -2}: {wire: firstWire},
				point{8, -3}: {wire: firstWire},
				point{8, -4}: {wire: firstWire},
				point{8, -5}: {wire: firstWire},
				point{7, -5}: {wire: firstWire},
				point{6, -5}: {wire: firstWire},
				point{5, -5}: {wire: firstWire},
				point{4, -5}: {wire: firstWire},
				point{3, -5}: {wire: firstWire},
				point{3, -4}: {wire: firstWire},
				point{3, -3}: {wire: firstWire},
				point{3, -2}: {wire: firstWire},
			},
		},
		{
			name: "firstWire: R8,U5,L5,D3; secondWire: U7,R6,D4,L4",
			wirePath: map[wire][]wireDirection{
				firstWire: {
					{value: "R", times: 8},
					{value: "U", times: 5},
					{value: "L", times: 5},
					{value: "D", times: 3},
				},
				secondWire: {
					{value: "U", times: 7},
					{value: "R", times: 6},
					{value: "D", times: 4},
					{value: "L", times: 4},
				},
			},
			grid: map[point]pathEntry{
				point{1, 0}:  {wire: firstWire},
				point{2, 0}:  {wire: firstWire},
				point{3, 0}:  {wire: firstWire},
				point{4, 0}:  {wire: firstWire},
				point{5, 0}:  {wire: firstWire},
				point{6, 0}:  {wire: firstWire},
				point{7, 0}:  {wire: firstWire},
				point{8, 0}:  {wire: firstWire},
				point{8, -1}: {wire: firstWire},
				point{8, -2}: {wire: firstWire},
				point{8, -3}: {wire: firstWire},
				point{8, -4}: {wire: firstWire},
				point{8, -5}: {wire: firstWire},
				point{7, -5}: {wire: firstWire},
				point{5, -5}: {wire: firstWire},
				point{4, -5}: {wire: firstWire},
				point{3, -5}: {wire: firstWire},
				point{3, -4}: {wire: firstWire},
				point{3, -2}: {wire: firstWire},

				point{0, -1}: {wire: secondWire},
				point{0, -2}: {wire: secondWire},
				point{0, -3}: {wire: secondWire},
				point{0, -4}: {wire: secondWire},
				point{0, -5}: {wire: secondWire},
				point{0, -6}: {wire: secondWire},
				point{0, -7}: {wire: secondWire},
				point{1, -7}: {wire: secondWire},
				point{2, -7}: {wire: secondWire},
				point{3, -7}: {wire: secondWire},
				point{4, -7}: {wire: secondWire},
				point{5, -7}: {wire: secondWire},
				point{6, -7}: {wire: secondWire},
				point{6, -6}: {wire: secondWire},
				point{6, -5}: {wire: both},
				point{6, -4}: {wire: secondWire},
				point{6, -3}: {wire: secondWire},
				point{5, -3}: {wire: secondWire},
				point{4, -3}: {wire: secondWire},
				point{3, -3}: {wire: both},
				point{2, -3}: {wire: secondWire},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGrid()
			g.constructAllPath(test.wirePath)
			assertNonNullMapsEqual(t, test.grid, g.path)
		})
	}
}

func TestGetClosestDistance(t *testing.T) {
	tests := []struct {
		name     string
		wirePath map[wire][]wireDirection
		distance int
	}{
		{
			name: "R8,U5,L5,D3 + U7,R6,D4,L4 = 6",
			wirePath: map[wire][]wireDirection{
				firstWire: {
					{value: "R", times: 8},
					{value: "U", times: 5},
					{value: "L", times: 5},
					{value: "D", times: 3},
				},
				secondWire: {
					{value: "U", times: 7},
					{value: "R", times: 6},
					{value: "D", times: 4},
					{value: "L", times: 4},
				},
			},
			distance: 6,
		},
		{
			name: "R75,D30,R83,U83,L12,D49,R71,U7,L72 + U62,R66,U55,R34,D71,R55,D58,R83 = 159",
			wirePath: map[wire][]wireDirection{
				firstWire: {
					{value: "R", times: 75},
					{value: "D", times: 30},
					{value: "R", times: 83},
					{value: "U", times: 83},
					{value: "L", times: 12},
					{value: "D", times: 49},
					{value: "R", times: 71},
					{value: "U", times: 7},
					{value: "L", times: 72},
				},
				secondWire: {
					{value: "U", times: 62},
					{value: "R", times: 66},
					{value: "U", times: 55},
					{value: "R", times: 34},
					{value: "D", times: 71},
					{value: "R", times: 55},
					{value: "D", times: 58},
					{value: "R", times: 83},
				},
			},
			distance: 159,
		},
		{
			name: "aoc_input03.txt",
			wirePath: func() map[wire][]wireDirection {
				path, err := readInputAsCSWP("aoc_input03.txt")
				require.NoError(t, err)
				return path
			}(),
			distance: 1674,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGrid()
			g.constructAllPath(test.wirePath)
			assert.Equal(t, test.distance, g.getClosestDistance())
		})
	}
}

func TestGetShortestDistance(t *testing.T) {
	tests := []struct {
		name     string
		wirePath map[wire][]wireDirection
		steps    int
	}{
		{
			name: "R8,U5,L5,D3 + U7,R6,D4,L4 = 30",
			wirePath: map[wire][]wireDirection{
				firstWire: {
					{value: "R", times: 8},
					{value: "U", times: 5},
					{value: "L", times: 5},
					{value: "D", times: 3},
				},
				secondWire: {
					{value: "U", times: 7},
					{value: "R", times: 6},
					{value: "D", times: 4},
					{value: "L", times: 4},
				},
			},
			steps: 30,
		},
		{
			name: "R75,D30,R83,U83,L12,D49,R71,U7,L72 + U62,R66,U55,R34,D71,R55,D58,R83 = 610",
			wirePath: map[wire][]wireDirection{
				firstWire: {
					{value: "R", times: 75},
					{value: "D", times: 30},
					{value: "R", times: 83},
					{value: "U", times: 83},
					{value: "L", times: 12},
					{value: "D", times: 49},
					{value: "R", times: 71},
					{value: "U", times: 7},
					{value: "L", times: 72},
				},
				secondWire: {
					{value: "U", times: 62},
					{value: "R", times: 66},
					{value: "U", times: 55},
					{value: "R", times: 34},
					{value: "D", times: 71},
					{value: "R", times: 55},
					{value: "D", times: 58},
					{value: "R", times: 83},
				},
			},
			steps: 610,
		},
		{
			name: "aoc_input03.txt",
			wirePath: func() map[wire][]wireDirection {
				path, err := readInputAsCSWP("aoc_input03.txt")
				require.NoError(t, err)
				return path
			}(),
			steps: 14012,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGrid()
			g.constructAllPath(test.wirePath)
			assert.Equal(t, test.steps, g.getMinSteps())
		})
	}
}

func assertNonNullMapsEqual(t *testing.T, first map[point]pathEntry, second map[point]pathEntry) {
	if first == nil || second == nil || len(first) != len(second) {
		t.Fail()
	}
	for keyFirst, valueFirst := range first {
		for keySecond, valueSecond := range second {
			assert.Equal(t, valueFirst.wire, second[keyFirst].wire)
			assert.Equal(t, valueSecond.wire, first[keySecond].wire)
		}
	}

}
