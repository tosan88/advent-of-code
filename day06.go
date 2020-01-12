package advent_of_code

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
