package day07

import (
	"github.com/stretchr/testify/assert"
	"github.com/tosan88/advent-of-code/aoc_io"
	"testing"
	"time"
)

//flaky tests
func TestAmplify(t *testing.T) {
	if testing.Short() {
		t.Skip("Flaky test skipped")
	}
	tests := []struct {
		name   string
		amp    Amplifier
		phases []int
		output int
	}{
		{
			name:   "Output: 43210",
			amp:    Amplifier{Code: []int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0}, Timeout: 2 * time.Second},
			phases: []int{4,3,2,1,0},
			output: 43210,
		},
		{
			name: "Output: 54321",
			amp: Amplifier{Code: []int{3,23,3,24,1002,24,10,24,1002,23,-1,23,
				101,5,23,23,1,24,23,23,4,23,99,0,0}, Timeout: 2 * time.Second},
			phases: []int{0,1,2,3,4},
			output: 54321,
		},
		{
			name: "Output: 65210",
			amp: Amplifier{Code: []int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,
				1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}, Timeout: 2 * time.Second},
			phases: []int{1,0,4,3,2},
			output: 65210,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualOutput, err := test.amp.Amplify(test.phases, 0)
			assert.NoError(t, err)
			assert.Equal(t, test.output, actualOutput)
		})
	}
}

func TestFindMaxThrusterSignal(t *testing.T) {
	if testing.Short() {
		t.Skip("Flaky test skipped")
	}
	tests := []struct {
		name   string
		amp    Amplifier
		output int
	}{
		{
			name:   "Output: 43210",
			amp:    Amplifier{Code: []int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0}, Timeout: 5 * time.Second},
			output: 43210,
		},
		{
			name: "Output: 54321",
			amp: Amplifier{Code: []int{3,23,3,24,1002,24,10,24,1002,23,-1,23,
				101,5,23,23,1,24,23,23,4,23,99,0,0}, Timeout: 5 * time.Second},
			output: 54321,
		},
		{
			name: "Output: 65210",
			amp: Amplifier{Code: []int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,
				1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}, Timeout: 5 * time.Second},
			output: 65210,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualOutput, err := test.amp.FindMaxThrusterSignal(false)
			assert.NoError(t, err)
			assert.Equal(t, test.output, actualOutput)

		})
	}
}

func TestAoCDay07Part1(t *testing.T) {
	if testing.Short() {
		t.Skip("Flaky test skipped")
	}
	ints, err := aoc_io.ReadInputAsCsi("aoc_input07.txt")
	if err != nil {
		t.Fatalf("ERROR: %v\n", err)
	}

	amp := Amplifier{Code: ints, Timeout: 10 * time.Second}
	maxThrusterSignal, _ := amp.FindMaxThrusterSignal(false)

	assert.Equal(t, 18812, maxThrusterSignal)
}


func TestAmplifyWithLoop(t *testing.T) {
	if testing.Short() {
		t.Skip("Flaky test skipped")
	}
	tests := []struct {
		name   string
		amp    Amplifier
		phases []int
		output int
	}{
		{
			name: "Output: 139629729",
			amp: Amplifier{Code: []int{3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,
				27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5}, Timeout: 2 * time.Second},
			phases: []int{9,8,7,6,5},
			output: 139629729,
		},
		{
			name: "Output: 18216",
			amp: Amplifier{Code: []int{3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,
				-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,
				53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10}, Timeout: 2 * time.Second},
			phases: []int{9,7,8,5,6},
			output: 18216,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualOutput, err := test.amp.AmplifyWithLoop(test.phases)
			assert.NoError(t, err)
			assert.Equal(t, test.output, actualOutput)
		})
	}
}


func TestAoCDay07Part2(t *testing.T) {
	if testing.Short() {
		t.Skip("Flaky test skipped")
	}
	ints, err := aoc_io.ReadInputAsCsi("aoc_input07.txt")
	if err != nil {
		t.Fatalf("ERROR: %v\n", err)
	}

	amp := Amplifier{Code: ints, Timeout: 10 * time.Second}
	maxThrusterSignal, err := amp.FindMaxThrusterSignal(true)

	assert.NoError(t, err)
	assert.Equal(t, 25534964, maxThrusterSignal)

}