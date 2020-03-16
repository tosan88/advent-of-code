package day08

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tosan88/advent-of-code/aoc_io"
	"reflect"
	"testing"
)

func TestImage_Checksum(t *testing.T) {
	type fields struct {
		Wide   int
		Tall   int
		Layers []Layer
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:"Want 1",
			fields:fields{Wide:3, Tall: 2, Layers: []Layer{
				{1,2,3,4,5,6},
				{7,8,9,0,1,2},

			}},
			want: 1,
		},
		{
			name:"Want 6",
			fields:fields{Wide:3, Tall: 2, Layers: []Layer{
				{1,2,1,1,2,6},
				{7,8,9,0,1,2},

			}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Image{
				Wide:   tt.fields.Wide,
				Tall:   tt.fields.Tall,
				Layers: tt.fields.Layers,
			}
			if got := i.Checksum(); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_findLayerWithFewestDigits(t *testing.T) {
	type fields struct {
		Wide   int
		Tall   int
		Layers []Layer
	}
	type args struct {
		digit int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Layer
	}{
		{
			name:"No occurrence",
			fields:fields{Wide:3, Tall: 2, Layers: []Layer{
				{1,2,3,4,5,6},
				{7,8,9,0,1,2},

			}},
			args:args{digit:0},
			want: Layer{1,2,3,4,5,6},
		},
		{
			name:"Same occurrence",
			fields:fields{Wide:3, Tall: 2, Layers: []Layer{
				{1,2,3,4,5,6},
				{7,8,9,0,1,2},

			}},
			args:args{digit:1},
			want: Layer{1,2,3,4,5,6},
		},
		{
			name:"More occurrence",
			fields:fields{Wide:3, Tall: 2, Layers: []Layer{
				{1,2,2,4,5,6},
				{7,8,9,0,1,2},

			}},
			args:args{digit:2},
			want: Layer{7,8,9,0,1,2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Image{
				Wide:   tt.fields.Wide,
				Tall:   tt.fields.Tall,
				Layers: tt.fields.Layers,
			}
			if got := i.findLayerWithFewestDigits(tt.args.digit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLayerWithFewestDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLayer_sumNumberOfDigits(t *testing.T) {
	type args struct {
		digit int
	}
	tests := []struct {
		name string
		l    Layer
		args args
		want int
	}{
		{
			name: "0s",
			l: Layer{1,2,3,2,1,2},
			args:args{digit:0},
			want: 0,
		},
		{
			name: "1s",
			l: Layer{1,2,3,2,1,2},
			args:args{digit:1},
			want: 2,
		},
		{
			name: "2s",
			l: Layer{1,2,3,2,1,2},
			args:args{digit:2},
			want: 3,
		},
		{
			name: "3s",
			l: Layer{1,2,3,2,1,2},
			args:args{digit:3},
			want: 1,
		},
		{
			name: "4s",
			l: Layer{1,2,3,2,1,2},
			args:args{digit:4},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.sumNumberOfDigits(tt.args.digit); got != tt.want {
				t.Errorf("sumNumberOfDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAoCDay08Part1(t *testing.T) {
	wide := 25
	tall := 6
	layersAsInts, err := aoc_io.ReadInputAsIntLayers("aoc_input08.txt", wide*tall)
	if err != nil {
		t.Fatalf("ERROR: %v\n", err)
	}

	var layer Layer
	var layers []Layer

	for _, l := range layersAsInts {
		layer := append(layer, l...)
		layers = append(layers, layer)
	}
	image := Image{Layers: layers}
	assert.Equal(t, 2375,image.Checksum())
}

func TestAoCDay08Part2(t *testing.T) {
	wide := 25
	tall := 6
	layersAsInts, err := aoc_io.ReadInputAsIntLayers("aoc_input08.txt", wide*tall)
	if err != nil {
		t.Fatalf("ERROR: %v\n", err)
	}

	var layer Layer
	var layers []Layer

	for _, l := range layersAsInts {
		layer := append(layer, l...)
		layers = append(layers, layer)
	}
	image := Image{Layers: layers, Wide:wide, Tall:tall}
	renderedLayer := image.Render()
	assert.Equal(t, Layer{1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1,
		0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0,
		1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0,
		1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0,
		0, 1, 0, 0, 0, 1, 0, 0}, renderedLayer)

	//should be printed out RKHRY shaped image
	for i := 0; i < len(renderedLayer); i++ {
		if i % wide == 0 {
			fmt.Print("\n")
		}
		renderedBit := fmt.Sprint(renderedLayer[i])
		if renderedBit == "0" {
			renderedBit = " "
		} else {
			renderedBit = "x"
		}
		fmt.Print(renderedBit)
	}
	fmt.Print("\n")
}