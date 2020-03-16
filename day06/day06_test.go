package day06

import (
	"github.com/stretchr/testify/assert"
	"github.com/tosan88/advent-of-code/aoc_io"
	"testing"
)

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
			name:   "A)B",
			fields: fields{orbits: map[string]string{"B": "A"}},
			want:   1,
		},
		{
			name:   "A)B B)C",
			fields: fields{orbits: map[string]string{"B": "A", "C": "B"}},
			want:   3,
		},
		{
			name:   "A)B B)C B)D",
			fields: fields{orbits: map[string]string{"B": "A", "C": "B", "D": "B"}},
			want:   5,
		},
		{
			name:   "A)B B)C C)D",
			fields: fields{orbits: map[string]string{"B": "A", "C": "B", "D": "C"}},
			want:   6,
		},
		{
			name:   "COM)B B)C C)D D)E E)F B)G G)H D)I E)J J)K K)L",
			fields: fields{orbits: map[string]string{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K"}},
			want:   42,
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
	orbitingSystem, err := aoc_io.ReadInputAsOM("aoc_input06.txt")
	if err != nil {
		t.Fatalf("ERROR: %v\n", err)
	}
	o := orbits{orbitingSystem}
	count := o.countOrbits()

	assert.Equal(t, 453028, count)
}

func TestCountTransfers(t *testing.T) {
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
		want   map[string]int
	}{
		{
			name:   "From COM",
			fields: fields{orbits: map[string]string{"A": "COM", "C": "A", "D": "C", "E": "D", "F": "E", "G": "A", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "YOU": "K", "SAN": "I"}},
			args:   args{planet: "COM"},
			want:   map[string]int{},
		},
		{
			name:   "From A",
			fields: fields{orbits: map[string]string{"A": "COM", "C": "A", "D": "C", "E": "D", "F": "E", "G": "A", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "YOU": "K", "SAN": "I"}},
			args:   args{planet: "A"},
			want:   map[string]int{"COM": 1},
		},
		{
			name:   "From YOU",
			fields: fields{orbits: map[string]string{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "YOU": "K", "SAN": "I"}},
			args:   args{planet: "YOU"},
			want:   map[string]int{"COM": 7, "K": 1, "J": 2, "E": 3, "D": 4, "C": 5, "B": 6},
		},
		{
			name:   "From SAN",
			fields: fields{orbits: map[string]string{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "YOU": "K", "SAN": "I"}},
			args:   args{planet: "SAN"},
			want:   map[string]int{"COM": 5, "I": 1, "D": 2, "C": 3, "B": 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orbits{
				orbits: tt.fields.orbits,
			}

			assert.Equal(t, tt.want, o.countTransfers(tt.args.planet, make(map[string]int)))
		})
	}

}

func Test_orbits_minTransfer(t *testing.T) {
	type fields struct {
		orbits map[string]string
	}
	type args struct {
		fromPlanet string
		toPlanet   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "B to B -> 0",
			fields: fields{orbits: map[string]string{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "YOU": "K", "SAN": "I"}},
			args:   args{fromPlanet: "B", toPlanet: "B"},
			want:   0,
		},
		{
			name:   "B to COM -> 1",
			fields: fields{orbits: map[string]string{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "YOU": "K", "SAN": "I"}},
			args:   args{fromPlanet: "B", toPlanet: "COM"},
			want:   1,
		},
		{
			name:   "COM to B -> 1",
			fields: fields{orbits: map[string]string{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "YOU": "K", "SAN": "I"}},
			args:   args{fromPlanet: "COM", toPlanet: "B"},
			want:   1,
		},
		{
			name:   "D to COM -> 3",
			fields: fields{orbits: map[string]string{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "YOU": "K", "SAN": "I"}},
			args:   args{fromPlanet: "D", toPlanet: "COM"},
			want:   3,
		},
		{
			name:   "COM to D -> 3",
			fields: fields{orbits: map[string]string{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "YOU": "K", "SAN": "I"}},
			args:   args{fromPlanet: "COM", toPlanet: "D"},
			want:   3,
		},
		{
			name:   "K to I -> 4",
			fields: fields{orbits: map[string]string{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "YOU": "K", "SAN": "I"}},
			args:   args{fromPlanet: "K", toPlanet: "I"},
			want:   4,
		},
		{
			name:   "I to K -> 4",
			fields: fields{orbits: map[string]string{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K", "YOU": "K", "SAN": "I"}},
			args:   args{fromPlanet: "I", toPlanet: "K"},
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orbits{
				orbits: tt.fields.orbits,
			}
			got := o.minTransfer(tt.args.fromPlanet, tt.args.toPlanet)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAoCDay06Part2(t *testing.T) {
	orbitingSystem, err := aoc_io.ReadInputAsOM("aoc_input06.txt")
	if err != nil {
		t.Fatalf("ERROR: %v\n", err)
	}
	o := orbits{orbitingSystem}
	count := o.minTransfer(orbitingSystem["YOU"], orbitingSystem["SAN"])

	assert.Equal(t, 562, count)
}
