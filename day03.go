package advent_of_code

import "math"

type Direction interface {
}

type pathEntry struct {
	wire wire
	step map[wire]int
}

type wire int

const (
	firstWire  = wire(1) //01 in binary
	secondWire = wire(2) //10 in binary
	both       = wire(3) //11 in binary
)

type wireDirection struct {
	value string
	times int
}

type point struct {
	x int
	y int
}

type grid struct {
	path map[point]pathEntry
}

func NewGrid() *grid {
	return &grid{path: make(map[point]pathEntry)}
}

func (g *grid) constructAllPath(wirePath map[wire][]wireDirection) {
	for w, dir := range wirePath {
		g.constructPath(w, dir)
	}
}

func (g *grid) constructPath(currentWire wire, directions []wireDirection) {
	posX, posY := 0, 0
	stepCounter := 1
	for _, dir := range directions {
		switch dir.value {
		case "R":
			for i := posX+1; i <= posX+dir.times; i++ {
				entry := g.path[point{i, posY}]
				entry.wire = currentWire | entry.wire
				if entry.step == nil {
					entry.step = make(map[wire]int)
				}
				if _, found := entry.step[currentWire]; !found {
					entry.step[currentWire] = stepCounter
				}

				g.path[point{i, posY}] = entry
				stepCounter++
			}
			posX += dir.times
		case "L":
			for i := posX-1; i >= posX-dir.times; i-- {
				entry := g.path[point{i, posY}]
				entry.wire = currentWire | entry.wire
				if entry.step == nil {
					entry.step = make(map[wire]int)

				}
				if _, found := entry.step[currentWire]; !found {
					entry.step[currentWire] = stepCounter
				}
				g.path[point{i, posY}] = entry
				stepCounter++
			}
			posX -= dir.times
		case "U":
			for i := posY-1; i >= posY-dir.times; i-- {
				entry := g.path[point{posX, i}]
				entry.wire = currentWire | entry.wire
				if entry.step == nil {
					entry.step = make(map[wire]int)

				}
				if _, found := entry.step[currentWire]; !found {
					entry.step[currentWire] = stepCounter
				}
				g.path[point{posX, i}] = entry
				stepCounter++
			}
			posY -= dir.times
		case "D":
			for i := posY+1; i <= posY+dir.times; i++ {
				entry := g.path[point{posX, i}]
				entry.wire = currentWire | entry.wire
				if entry.step == nil {
					entry.step = make(map[wire]int)
				}
				if _, found := entry.step[currentWire]; !found {
					entry.step[currentWire] = stepCounter
				}
				g.path[point{posX, i}] = entry
				stepCounter++
			}
			posY += dir.times
		}
	}
}

func (g *grid) getClosestDistance() int {
	minDistance := math.MaxInt64 //a sufficiently large number
	for p, entry := range g.path {
		if entry.wire == both && !(p.x == 0 && p.y == 0) {
			dist := calculateDistance(p)
			if dist < minDistance {
				minDistance = dist
			}
		}
	}
	return minDistance
}

func calculateDistance(p point) int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func (g *grid) getMinSteps() int {
	minSteps := math.MaxInt64 //a sufficiently large number
	for p, entry := range g.path {
		if entry.wire == both && !(p.x == 0 && p.y == 0) {
			steps := calculateSteps(entry.step)
			if steps < minSteps {
				minSteps = steps
			}
		}
	}
	return minSteps
}

func calculateSteps(steps map[wire]int) int {
	sum := 0
	for _, step := range steps {
		sum += step
	}
	return sum
}
