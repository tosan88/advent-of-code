package day03

import "math"

type Direction interface {
}

type PathEntry struct {
	wire Wire
	step map[Wire]int
}

type Wire int

const (
	firstWire  = Wire(1) //01 in binary
	secondWire = Wire(2) //10 in binary
	both       = Wire(3) //11 in binary
)

type WireDirection struct {
	Value string
	Times int
}

type Point struct {
	X int
	Y int
}

type Grid struct {
	Path map[Point]PathEntry
}

func NewGrid() *Grid {
	return &Grid{Path: make(map[Point]PathEntry)}
}

func (g *Grid) constructAllPath(wirePath map[Wire][]WireDirection) {
	for w, dir := range wirePath {
		g.constructPath(w, dir)
	}
}

func (g *Grid) constructPath(currentWire Wire, directions []WireDirection) {
	posX, posY := 0, 0
	stepCounter := 1
	for _, dir := range directions {
		switch dir.Value {
		case "R":
			for i := posX+1; i <= posX+dir.Times; i++ {
				entry := g.Path[Point{i, posY}]
				entry.wire = currentWire | entry.wire
				if entry.step == nil {
					entry.step = make(map[Wire]int)
				}
				if _, found := entry.step[currentWire]; !found {
					entry.step[currentWire] = stepCounter
				}

				g.Path[Point{i, posY}] = entry
				stepCounter++
			}
			posX += dir.Times
		case "L":
			for i := posX-1; i >= posX-dir.Times; i-- {
				entry := g.Path[Point{i, posY}]
				entry.wire = currentWire | entry.wire
				if entry.step == nil {
					entry.step = make(map[Wire]int)

				}
				if _, found := entry.step[currentWire]; !found {
					entry.step[currentWire] = stepCounter
				}
				g.Path[Point{i, posY}] = entry
				stepCounter++
			}
			posX -= dir.Times
		case "U":
			for i := posY-1; i >= posY-dir.Times; i-- {
				entry := g.Path[Point{posX, i}]
				entry.wire = currentWire | entry.wire
				if entry.step == nil {
					entry.step = make(map[Wire]int)

				}
				if _, found := entry.step[currentWire]; !found {
					entry.step[currentWire] = stepCounter
				}
				g.Path[Point{posX, i}] = entry
				stepCounter++
			}
			posY -= dir.Times
		case "D":
			for i := posY+1; i <= posY+dir.Times; i++ {
				entry := g.Path[Point{posX, i}]
				entry.wire = currentWire | entry.wire
				if entry.step == nil {
					entry.step = make(map[Wire]int)
				}
				if _, found := entry.step[currentWire]; !found {
					entry.step[currentWire] = stepCounter
				}
				g.Path[Point{posX, i}] = entry
				stepCounter++
			}
			posY += dir.Times
		}
	}
}

func (g *Grid) getClosestDistance() int {
	minDistance := math.MaxInt64 //a sufficiently large number
	for p, entry := range g.Path {
		if entry.wire == both && !(p.X == 0 && p.Y == 0) {
			dist := calculateDistance(p)
			if dist < minDistance {
				minDistance = dist
			}
		}
	}
	return minDistance
}

func calculateDistance(p Point) int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
}

func (g *Grid) getMinSteps() int {
	minSteps := math.MaxInt64 //a sufficiently large number
	for p, entry := range g.Path {
		if entry.wire == both && !(p.X == 0 && p.Y == 0) {
			steps := calculateSteps(entry.step)
			if steps < minSteps {
				minSteps = steps
			}
		}
	}
	return minSteps
}

func calculateSteps(steps map[Wire]int) int {
	sum := 0
	for _, step := range steps {
		sum += step
	}
	return sum
}
