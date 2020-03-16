package day08

import (
	"math"
)

type Layer []int

type Image struct {
	Wide int
	Tall int
	Layers []Layer
}

type Condition func(layer Layer) bool

func (image *Image) findLayerWithFewestDigits(digit int) Layer {
	minDigits := math.MaxInt32
	var layer Layer
	for _, l := range image.Layers {
		sumDigits := l.sumNumberOfDigits(digit)
		if sumDigits < minDigits {
			minDigits = sumDigits
			layer = l
		}
	}
	return layer
}

func (l Layer) sumNumberOfDigits(digit int) int {
	sum := 0
	for _, dig := range l {
		if dig == digit {
			sum++
		}
	}
	return sum
}

func (image *Image) Checksum() int {
	layer := image.findLayerWithFewestDigits(0)
	sumOnes := layer.sumNumberOfDigits(1)
	sumTwos := layer.sumNumberOfDigits(2)
	return sumOnes * sumTwos
}

func (image *Image) Render() Layer {
	var result Layer
	transparentDigit := 2
	for _, layer := range image.Layers {
		if result == nil {
			result = layer
			continue
		}
		for i, digit := range layer {
			if result[i] == transparentDigit {
				result[i] = digit
			}
		}
	}
	return result
}