package aoc_io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//ReadInputAsIntPerLines reads the input file with the given fileName as integers one per line
func ReadInputAsIntPerLines(fileName string) ([]int, error) {
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
func ReadInputAsCsi(fileName string) ([]int, error) {
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

type Wire int

type WireDirection struct {
	Value string
	Times int
}

//readInputAsCSWP reads the input file with the given fileName as comma-separated wire path
func ReadInputAsCSWP(fileName string) (map[Wire][]WireDirection, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("reading input file: %v\n", err)
	}
	defer func() {
		_ = f.Close()
	}()

	path := make(map[Wire][]WireDirection)

	scanner := bufio.NewScanner(f)
	wireNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		directionsAsStrings := strings.Split(line, ",")
		var wdirs []WireDirection
		for _, dir := range directionsAsStrings {
			times, err := strconv.Atoi(dir[1:])
			if err != nil {
				return nil, fmt.Errorf("converting number in input file: %v\n", err)
			}
			wdirs = append(wdirs, WireDirection{Value: string(dir[0]), Times: times})
		}
		path[Wire(wireNumber)] = wdirs
		wireNumber++
	}

	return path, nil
}

//readInputAsOM reads the input file with the given fileName as an orbit map
func ReadInputAsOM(fileName string) (map[string]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("reading input file: %v\n", err)
	}
	defer func() {
		_ = f.Close()
	}()

	orbits := make(map[string]string)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		orbit := strings.Split(line, ")")
		orbits[orbit[1]] = orbit[0]
	}

	return orbits, nil
}

func ReadInputAsIntLayers(fileName string, layerLen int) ([][]int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("reading input file: %v\n", err)
	}
	defer func() {
		_ = f.Close()
	}()

	var result [][]int
	currentLen := 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)
	var layer []int

	for scanner.Scan() {
		text := scanner.Text()
		if text == "\n" {
			continue
		}
		digit, err := strconv.Atoi(text)
		if err != nil {
			return nil, fmt.Errorf("reading int: %v\n", err)
		}
		layer = append(layer, digit)
		currentLen++
		if currentLen == layerLen {
			result = append(result, layer)
			layer = []int{}
			currentLen = 0
		}
	}
	if layer != nil && len(layer) > 0 {
		result = append(result, layer)
	}
	return result, nil
}
