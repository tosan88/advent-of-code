package advent_of_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_orbitSystem_countOrbit(t *testing.T) {
	type fields struct {
		orbits map[string]string
	}
	type args struct {
		planet string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orbits{
				orbits: tt.fields.orbits,
			}
			if got := o.countOrbit(tt.args.planet); got != tt.want {
				t.Errorf("countOrbit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orbitSystem_countOrbits(t *testing.T) {
	type fields struct {
		orbits map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "A)B",
			fields:fields{orbits: map[string]string{"B" : "A"}},
			want: 1,
		},
		{
			name: "A)B B)C",
			fields:fields{orbits: map[string]string{"B" : "A", "C" : "B"}},
			want: 3,
		},
		{
			name: "A)B B)C B)D",
			fields:fields{orbits: map[string]string{"B" : "A", "C" : "B", "D" : "B"}},
			want: 5,
		},
		{
			name: "A)B B)C C)D",
			fields:fields{orbits: map[string]string{"B" : "A", "C" : "B", "D" : "C"}},
			want: 6,
		},
		{
			name: "COM)B B)C C)D D)E E)F B)G G)H D)I E)J J)K K)L",
			fields:fields{orbits: map[string]string{"B" : "COM", "C" : "B", "D" : "C", "E" : "D", "F" : "E", "G" : "B", "H" : "G", "I": "D", "J": "E", "K": "J", "L": "K"}},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orbits{
				orbits: tt.fields.orbits,
			}
			if got := o.countOrbits(); got != tt.want {
				t.Errorf("countOrbits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAoCDay06Part1(t *testing.T) {
	orbitingSystem, err := readInputAsOM("aoc_input06.txt")
	if err != nil {
		t.Fatalf("ERROR: %v\n", err)
	}
	o := orbits{orbitingSystem}
	count := o.countOrbits()

	assert.Equal(t, 453028, count)
}