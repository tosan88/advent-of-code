package advent_of_code

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNext(t *testing.T) {
	tests := []struct {
		name   string
		output bool
		cursor int
		input  []int
		code   []int
	}{
		{
			name:   "No input code",
			output: false,
			cursor: 0,
			input:  nil,
			code:   nil,
		},
		{
			name:   "Halt program",
			output: false,
			cursor: 0,
			input:  []int{99},
			code:   []int{99},
		},
		{
			name:   "Addition replacing next position",
			output: true,
			cursor: 4,
			input:  []int{1, 0, 3, 3},
			code:   []int{1, 0, 3, 4},
		},
		{
			name:   "Multiplication replacing next position",
			output: true,
			cursor: 4,
			input:  []int{2, 3, 0, 3},
			code:   []int{2, 3, 0, 6},
		},
		{
			name:   "Addition replacing first position",
			output: true,
			cursor: 4,
			input:  []int{1, 2, 1, 0},
			code:   []int{3, 2, 1, 0},
		},
		{
			name:   "Multiplication replacing first position",
			output: true,
			cursor: 4,
			input:  []int{2, 2, 2, 0},
			code:   []int{4, 2, 2, 0},
		},
		{
			name:   "Invalid operation code",
			output: false,
			cursor: 0,
			input:  []int{8},
			code:   []int{8},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := NewProgram(test.input)
			assert.Equal(t, test.output, p.next())
			assert.Equal(t, test.cursor, p.cursor)
			assert.Equal(t, test.code, p.code)
		})
	}
}

func TestRunCode(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "1,0,0,0,99 -> 2,0,0,0,99",
			input:  []int{1, 0, 0, 0, 99},
			output: []int{2, 0, 0, 0, 99},
		},
		{
			name:   "2,3,0,3,99 -> 2,3,0,6,99",
			input:  []int{2, 3, 0, 3, 99},
			output: []int{2, 3, 0, 6, 99},
		},
		{
			name:   "2,4,4,5,99,0 -> 2,4,4,5,99,9801",
			input:  []int{2, 4, 4, 5, 99, 0},
			output: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			name:   "1,1,1,4,99,5,6,0,99 -> 30,1,1,4,2,5,6,0,99",
			input:  []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			output: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, test := range tests {
		p := NewProgram(test.input)
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, p.RunCode(test.input[1], test.input[2]))
		})
	}
}

func TestRunCodeAoCPart1(t *testing.T) {
	ints, err := readInputAsCsi("aoc_input02.txt")
	if err != nil {
		t.Fatalf("ERROR: %v\n", err)
	}
	require.Equal(t, []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 9, 1, 19, 1, 19, 6, 23, 2, 6, 23, 27, 2,
		27, 9, 31, 1, 5, 31, 35, 1, 35, 10, 39, 2, 39, 9, 43, 1, 5, 43, 47, 2, 47, 10, 51, 1, 51, 6, 55, 1, 5, 55, 59,
		2, 6, 59, 63, 2, 63, 6, 67, 1, 5, 67, 71, 1, 71, 9, 75, 2, 75, 10, 79, 1, 79, 5, 83, 1, 10, 83, 87, 1, 5, 87,
		91, 2, 13, 91, 95, 1, 95, 10, 99, 2, 99, 13, 103, 1, 103, 5, 107, 1, 107, 13, 111, 2, 111, 9, 115, 1, 6, 115,
		119, 2, 119, 6, 123, 1, 123, 6, 127, 1, 127, 9, 131, 1, 6, 131, 135, 1, 135, 2, 139, 1, 139, 10, 0, 99,
		2, 0, 14, 0}, ints)

	p := NewProgram(ints)
	result := p.RunCode(12, 2)

	assert.Equal(t, []int{6730673, 12, 2, 2, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 9, 1, 36, 1, 19, 6, 38, 2, 6, 23,
		76, 2, 27, 9, 228, 1, 5, 31, 229, 1, 35, 10, 233, 2, 39, 9, 699, 1, 5, 43, 700, 2, 47, 10, 2800, 1, 51, 6, 2802,
		1, 5, 55, 2803, 2, 6, 59, 5606, 2, 63, 6, 11212, 1, 5, 67, 11213, 1, 71, 9, 11216, 2, 75, 10, 44864, 1,
		79, 5, 44865, 1, 10, 83, 44869, 1, 5, 87, 44870, 2, 13, 91, 224350, 1, 95, 10, 224354, 2, 99, 13, 1121770,
		1, 103, 5, 1121771, 1, 107, 13, 1121776, 2, 111, 9, 3365328, 1, 6, 115, 3365330, 2, 119, 6, 6730660, 1, 123,
		6, 6730662, 1, 127, 9, 6730665, 1, 6, 131, 6730667, 1, 135, 2, 6730669, 1, 139, 10, 0, 99, 2, 0, 14, 0}, result)
	//fmt.Printf("%v\n", result[0])
}

func TestRunCodeAoCPart2(t *testing.T) {
	ints, err := readInputAsCsi("aoc_input02.txt")
	if err != nil {
		t.Fatalf("ERROR: %v\n", err)
	}
	require.Equal(t, []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 9, 1, 19, 1, 19, 6, 23, 2, 6, 23, 27, 2,
		27, 9, 31, 1, 5, 31, 35, 1, 35, 10, 39, 2, 39, 9, 43, 1, 5, 43, 47, 2, 47, 10, 51, 1, 51, 6, 55, 1, 5, 55, 59,
		2, 6, 59, 63, 2, 63, 6, 67, 1, 5, 67, 71, 1, 71, 9, 75, 2, 75, 10, 79, 1, 79, 5, 83, 1, 10, 83, 87, 1, 5, 87,
		91, 2, 13, 91, 95, 1, 95, 10, 99, 2, 99, 13, 103, 1, 103, 5, 107, 1, 107, 13, 111, 2, 111, 9, 115, 1, 6, 115,
		119, 2, 119, 6, 123, 1, 123, 6, 127, 1, 127, 9, 131, 1, 6, 131, 135, 1, 135, 2, 139, 1, 139, 10, 0, 99,
		2, 0, 14, 0}, ints)

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			//need to make a copy of the slice, as program internally modifies the same slice
			input := make([]int, len(ints))
			copy(input, ints)
			
			p := NewProgram(input)
			result := p.RunCode(noun, verb)
			if result[0] == 19690720 {
				fmt.Printf("Noun and verb: %d% d\n", noun, verb)
				return
			}
		}
	}
	t.Fail()
}
