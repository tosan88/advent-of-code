package advent_of_code

import "math"

type orbits struct {
	orbits map[string]string
}

func (o *orbits) countOrbits() int {
	count := make(map[string]int)

	for planet, mass := range o.orbits {
		count[planet] = 1 + o.countOrbit(mass)
	}
	sum := 0
	for _, c := range count {
		sum += c
	}

	return sum
}

func (o *orbits) countOrbit(planet string) int {
	mass, found := o.orbits[planet]
	if !found {
		return 0
	}
	return 1 + o.countOrbit(mass)
}

func (o *orbits) countTransfers(planet string, visited map[string]int) map[string]int {
	mass, found := o.orbits[planet]
	if !found {
		return visited
	}
	visited[mass] = 1 + visited[planet]
	return o.countTransfers(mass, visited)
}

func (o *orbits) minTransfer(fromPlanet, toPlanet string) int {
	if fromPlanet == toPlanet {
		return 0
	}
	visitedFrom := o.countTransfers(fromPlanet, map[string]int{fromPlanet: 0})
	visitedTo := o.countTransfers(toPlanet, map[string]int{toPlanet: 0})

	minTransfer := math.MaxInt32
	for from, stepsFrom := range visitedFrom {
		if stepsTo, found := visitedTo[from]; found {
			if stepsFrom+stepsTo < minTransfer {
				minTransfer = stepsFrom + stepsTo
			}
		}
	}
	return minTransfer

}
