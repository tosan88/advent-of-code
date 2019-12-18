package advent_of_code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//readInputAsIntPerLines reads the input file with the given fileName as integers one per line
func readInputAsIntPerLines(fileName string) ([]int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("reading input file: %v\n", err)
	}
	defer func() {
		_ = f.Close()
	}()

	var ints []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("converting number in input file: %v\n", err)
		}
		ints = append(ints, num)
	}

	return ints, nil
}

//readInputAsCsi reads the input file with the given fileName as comma-separated integers
func readInputAsCsi(fileName string) ([]int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("reading input file: %v\n", err)
	}
	defer func() {
		_ = f.Close()
	}()

	var nums []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		numsAsStrings := strings.Split(line, ",")
		for _, n := range numsAsStrings {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, fmt.Errorf("converting number in input file: %v\n", err)
			}
			nums = append(nums, num)
		}
	}

	return nums, nil
}

//readInputAsCSWP reads the input file with the given fileName as comma-separated wire path
func readInputAsCSWP(fileName string) (map[wire][]wireDirection, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("reading input file: %v\n", err)
	}
	defer func() {
		_ = f.Close()
	}()

	path := make(map[wire][]wireDirection)

	scanner := bufio.NewScanner(f)
	wireNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		directionsAsStrings := strings.Split(line, ",")
		var wdirs []wireDirection
		for _, dir := range directionsAsStrings {
			times, err := strconv.Atoi(dir[1:])
			if err != nil {
				return nil, fmt.Errorf("converting number in input file: %v\n", err)
			}
			wdirs = append(wdirs, wireDirection{value: string(dir[0]), times: times})
		}
		path[wire(wireNumber)] = wdirs
		wireNumber++
	}

	return path, nil
}
