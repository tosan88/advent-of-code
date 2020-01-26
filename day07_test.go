package advent_of_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAmplify(t *testing.T) {
	tests := []struct {
		name string
		amp amplifier
		phases []int
		output int
	}{
		{
			name: "Output: 43210",
			amp: amplifier{[]int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0}},
			phases: []int{4,3,2,1,0},
			output: 43210,
		},
		{
			name: "Output: 54321",
			amp: amplifier{[]int{3,23,3,24,1002,24,10,24,1002,23,-1,23,
				101,5,23,23,1,24,23,23,4,23,99,0,0}},
			phases: []int{0,1,2,3,4},
			output: 54321,
		},
		{
			name: "Output: 65210",
			amp: amplifier{[]int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,
				1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}},
			phases: []int{1,0,4,3,2},
			output: 65210,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualOutput := test.amp.amplify(test.phases, 0)
			assert.Equal(t, test.output, actualOutput)
		})
	}
}

func TestFindMaxThrusterSignal(t *testing.T) {
	tests := []struct {
		name string
		amp amplifier
		output int
	}{
		{
			name: "Output: 43210",
			amp: amplifier{[]int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0}},
			output: 43210,
		},
		{
			name: "Output: 54321",
			amp: amplifier{[]int{3,23,3,24,1002,24,10,24,1002,23,-1,23,
				101,5,23,23,1,24,23,23,4,23,99,0,0}},
			output: 54321,
		},
		{
			name: "Output: 65210",
			amp: amplifier{[]int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,
				1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}},
			output: 65210,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualOutput := test.amp.findMaxThrusterSignal()
			assert.Equal(t, test.output, actualOutput)

		})
	}
}

func TestAoCDay07Part1(t *testing.T) {
	ints, err := readInputAsCsi("aoc_input07.txt")
	if err != nil {
		t.Fatalf("ERROR: %v\n", err)
	}

	amp := amplifier{ints}

	maxThrusterSignal := amp.findMaxThrusterSignal()

	assert.Equal(t, 18812, maxThrusterSignal)

}

//func TestAoCDay07Part2(t *testing.T) {
//	ints, err := readInputAsCsi("aoc_input07.txt")
//	if err != nil {
//		t.Fatalf("ERROR: %v\n", err)
//	}
//
//	amp := amplifier{ints}
//
//	maxThrusterSignal := amp.findMaxThrusterSignalPart2()
//
//	assert.Equal(t, 18812, maxThrusterSignal)
//
//}