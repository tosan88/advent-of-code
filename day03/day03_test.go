package day03

import (
	"github.com/tosan88/advent-of-code/aoc_io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConstructPath(t *testing.T) {
	tests := []struct {
		name     string
		wirePath map[Wire][]WireDirection
		grid     map[Point]PathEntry
	}{
		{
			name: "firstWire: R8,U5,L5,D3",
			wirePath: map[Wire][]WireDirection{
				firstWire: {
					{Value: "R", Times: 8},
					{Value: "U", Times: 5},
					{Value: "L", Times: 5},
					{Value: "D", Times: 3},
				},
			},
			grid: map[Point]PathEntry{

				Point{1, 0}:  {wire: firstWire},
				Point{2, 0}:  {wire: firstWire},
				Point{3, 0}:  {wire: firstWire},
				Point{4, 0}:  {wire: firstWire},
				Point{5, 0}:  {wire: firstWire},
				Point{6, 0}:  {wire: firstWire},
				Point{7, 0}:  {wire: firstWire},
				Point{8, 0}:  {wire: firstWire},
				Point{8, -1}: {wire: firstWire},
				Point{8, -2}: {wire: firstWire},
				Point{8, -3}: {wire: firstWire},
				Point{8, -4}: {wire: firstWire},
				Point{8, -5}: {wire: firstWire},
				Point{7, -5}: {wire: firstWire},
				Point{6, -5}: {wire: firstWire},
				Point{5, -5}: {wire: firstWire},
				Point{4, -5}: {wire: firstWire},
				Point{3, -5}: {wire: firstWire},
				Point{3, -4}: {wire: firstWire},
				Point{3, -3}: {wire: firstWire},
				Point{3, -2}: {wire: firstWire},
			},
		},
		{
			name: "secondWire: U7,R6,D4,L4",
			wirePath: map[Wire][]WireDirection{
				secondWire: {
					{Value: "U", Times: 7},
					{Value: "R", Times: 6},
					{Value: "D", Times: 4},
					{Value: "L", Times: 4},
				},
			},
			grid: map[Point]PathEntry{

				Point{0, -1}: {wire: secondWire},
				Point{0, -2}: {wire: secondWire},
				Point{0, -3}: {wire: secondWire},
				Point{0, -4}: {wire: secondWire},
				Point{0, -5}: {wire: secondWire},
				Point{0, -6}: {wire: secondWire},
				Point{0, -7}: {wire: secondWire},
				Point{1, -7}: {wire: secondWire},
				Point{2, -7}: {wire: secondWire},
				Point{3, -7}: {wire: secondWire},
				Point{4, -7}: {wire: secondWire},
				Point{5, -7}: {wire: secondWire},
				Point{6, -7}: {wire: secondWire},
				Point{6, -6}: {wire: secondWire},
				Point{6, -5}: {wire: secondWire},
				Point{6, -4}: {wire: secondWire},
				Point{6, -3}: {wire: secondWire},
				Point{5, -3}: {wire: secondWire},
				Point{4, -3}: {wire: secondWire},
				Point{3, -3}: {wire: secondWire},
				Point{2, -3}: {wire: secondWire},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGrid()
			for w, dirs := range test.wirePath {
				g.constructPath(w, dirs)
			}
			assertNonNullMapsEqual(t, test.grid, g.Path)
		})
	}
}

func TestConstructAllPath(t *testing.T) {
	tests := []struct {
		name     string
		wirePath map[Wire][]WireDirection
		grid     map[Point]PathEntry
	}{
		{
			name: "firstWire: R8,U5,L5,D3",
			wirePath: map[Wire][]WireDirection{
				firstWire: {
					{Value: "R", Times: 8},
					{Value: "U", Times: 5},
					{Value: "L", Times: 5},
					{Value: "D", Times: 3},
				},
			},
			grid: map[Point]PathEntry{

				Point{1, 0}:  {wire: firstWire},
				Point{2, 0}:  {wire: firstWire},
				Point{3, 0}:  {wire: firstWire},
				Point{4, 0}:  {wire: firstWire},
				Point{5, 0}:  {wire: firstWire},
				Point{6, 0}:  {wire: firstWire},
				Point{7, 0}:  {wire: firstWire},
				Point{8, 0}:  {wire: firstWire},
				Point{8, -1}: {wire: firstWire},
				Point{8, -2}: {wire: firstWire},
				Point{8, -3}: {wire: firstWire},
				Point{8, -4}: {wire: firstWire},
				Point{8, -5}: {wire: firstWire},
				Point{7, -5}: {wire: firstWire},
				Point{6, -5}: {wire: firstWire},
				Point{5, -5}: {wire: firstWire},
				Point{4, -5}: {wire: firstWire},
				Point{3, -5}: {wire: firstWire},
				Point{3, -4}: {wire: firstWire},
				Point{3, -3}: {wire: firstWire},
				Point{3, -2}: {wire: firstWire},
			},
		},
		{
			name: "firstWire: R8,U5,L5,D3; secondWire: U7,R6,D4,L4",
			wirePath: map[Wire][]WireDirection{
				firstWire: {
					{Value: "R", Times: 8},
					{Value: "U", Times: 5},
					{Value: "L", Times: 5},
					{Value: "D", Times: 3},
				},
				secondWire: {
					{Value: "U", Times: 7},
					{Value: "R", Times: 6},
					{Value: "D", Times: 4},
					{Value: "L", Times: 4},
				},
			},
			grid: map[Point]PathEntry{
				Point{1, 0}:  {wire: firstWire},
				Point{2, 0}:  {wire: firstWire},
				Point{3, 0}:  {wire: firstWire},
				Point{4, 0}:  {wire: firstWire},
				Point{5, 0}:  {wire: firstWire},
				Point{6, 0}:  {wire: firstWire},
				Point{7, 0}:  {wire: firstWire},
				Point{8, 0}:  {wire: firstWire},
				Point{8, -1}: {wire: firstWire},
				Point{8, -2}: {wire: firstWire},
				Point{8, -3}: {wire: firstWire},
				Point{8, -4}: {wire: firstWire},
				Point{8, -5}: {wire: firstWire},
				Point{7, -5}: {wire: firstWire},
				Point{5, -5}: {wire: firstWire},
				Point{4, -5}: {wire: firstWire},
				Point{3, -5}: {wire: firstWire},
				Point{3, -4}: {wire: firstWire},
				Point{3, -2}: {wire: firstWire},

				Point{0, -1}: {wire: secondWire},
				Point{0, -2}: {wire: secondWire},
				Point{0, -3}: {wire: secondWire},
				Point{0, -4}: {wire: secondWire},
				Point{0, -5}: {wire: secondWire},
				Point{0, -6}: {wire: secondWire},
				Point{0, -7}: {wire: secondWire},
				Point{1, -7}: {wire: secondWire},
				Point{2, -7}: {wire: secondWire},
				Point{3, -7}: {wire: secondWire},
				Point{4, -7}: {wire: secondWire},
				Point{5, -7}: {wire: secondWire},
				Point{6, -7}: {wire: secondWire},
				Point{6, -6}: {wire: secondWire},
				Point{6, -5}: {wire: both},
				Point{6, -4}: {wire: secondWire},
				Point{6, -3}: {wire: secondWire},
				Point{5, -3}: {wire: secondWire},
				Point{4, -3}: {wire: secondWire},
				Point{3, -3}: {wire: both},
				Point{2, -3}: {wire: secondWire},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGrid()
			g.constructAllPath(test.wirePath)
			assertNonNullMapsEqual(t, test.grid, g.Path)
		})
	}
}

func TestGetClosestDistance(t *testing.T) {
	tests := []struct {
		name     string
		wirePath map[Wire][]WireDirection
		distance int
	}{
		{
			name: "R8,U5,L5,D3 + U7,R6,D4,L4 = 6",
			wirePath: map[Wire][]WireDirection{
				firstWire: {
					{Value: "R", Times: 8},
					{Value: "U", Times: 5},
					{Value: "L", Times: 5},
					{Value: "D", Times: 3},
				},
				secondWire: {
					{Value: "U", Times: 7},
					{Value: "R", Times: 6},
					{Value: "D", Times: 4},
					{Value: "L", Times: 4},
				},
			},
			distance: 6,
		},
		{
			name: "R75,D30,R83,U83,L12,D49,R71,U7,L72 + U62,R66,U55,R34,D71,R55,D58,R83 = 159",
			wirePath: map[Wire][]WireDirection{
				firstWire: {
					{Value: "R", Times: 75},
					{Value: "D", Times: 30},
					{Value: "R", Times: 83},
					{Value: "U", Times: 83},
					{Value: "L", Times: 12},
					{Value: "D", Times: 49},
					{Value: "R", Times: 71},
					{Value: "U", Times: 7},
					{Value: "L", Times: 72},
				},
				secondWire: {
					{Value: "U", Times: 62},
					{Value: "R", Times: 66},
					{Value: "U", Times: 55},
					{Value: "R", Times: 34},
					{Value: "D", Times: 71},
					{Value: "R", Times: 55},
					{Value: "D", Times: 58},
					{Value: "R", Times: 83},
				},
			},
			distance: 159,
		},
		{
			name: "aoc_input03.txt",
			wirePath: func() map[Wire][]WireDirection {
				path, err := aoc_io.ReadInputAsCSWP("aoc_input03.txt")
				require.NoError(t, err)

				return mapInput(path)
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

func mapInput(path map[aoc_io.Wire][]aoc_io.WireDirection) map[Wire][]WireDirection {
	result := make(map[Wire][]WireDirection)
	for wire, directions := range path {
		var dirs []WireDirection
		for _, dir := range directions {
			dirs = append(dirs, WireDirection{
				Value: dir.Value,
				Times: dir.Times,
			})
		}
		result[Wire(wire)] = dirs
	}
	return result
}

func TestGetShortestDistance(t *testing.T) {
	tests := []struct {
		name     string
		wirePath map[Wire][]WireDirection
		steps    int
	}{
		{
			name: "R8,U5,L5,D3 + U7,R6,D4,L4 = 30",
			wirePath: map[Wire][]WireDirection{
				firstWire: {
					{Value: "R", Times: 8},
					{Value: "U", Times: 5},
					{Value: "L", Times: 5},
					{Value: "D", Times: 3},
				},
				secondWire: {
					{Value: "U", Times: 7},
					{Value: "R", Times: 6},
					{Value: "D", Times: 4},
					{Value: "L", Times: 4},
				},
			},
			steps: 30,
		},
		{
			name: "R75,D30,R83,U83,L12,D49,R71,U7,L72 + U62,R66,U55,R34,D71,R55,D58,R83 = 610",
			wirePath: map[Wire][]WireDirection{
				firstWire: {
					{Value: "R", Times: 75},
					{Value: "D", Times: 30},
					{Value: "R", Times: 83},
					{Value: "U", Times: 83},
					{Value: "L", Times: 12},
					{Value: "D", Times: 49},
					{Value: "R", Times: 71},
					{Value: "U", Times: 7},
					{Value: "L", Times: 72},
				},
				secondWire: {
					{Value: "U", Times: 62},
					{Value: "R", Times: 66},
					{Value: "U", Times: 55},
					{Value: "R", Times: 34},
					{Value: "D", Times: 71},
					{Value: "R", Times: 55},
					{Value: "D", Times: 58},
					{Value: "R", Times: 83},
				},
			},
			steps: 610,
		},
		{
			name: "aoc_input03.txt",
			wirePath: func() map[Wire][]WireDirection {
				path, err := aoc_io.ReadInputAsCSWP("aoc_input03.txt")
				require.NoError(t, err)
				return mapInput(path)
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

func assertNonNullMapsEqual(t *testing.T, first map[Point]PathEntry, second map[Point]PathEntry) {
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
