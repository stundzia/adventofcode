package day6

import (
	"strings"
)

type system struct {
	bodies map[string]*body
}

type body struct {
	inSystem *system
	name string
	orbits *body
	satellites map[string]*body
	distanceToCenter int
}

func NewSystemFromInput(input []string) *system {
	ss := &system{bodies: map[string]*body{}}
	for _, line := range input {
		bodies := strings.Split(line, ")")
		ss.createOrbitBodyPair(bodies[0], bodies[1])
	}
	return ss
}


func (ss *system) getOrCreateBody(name string) (cosmicBody *body) {
	if cosmicBody, ok := ss.bodies[name]; !ok { // TODO: see below, this seems to assign nil to cosmicBody and true to ok
		cosmicBody = &body{
			inSystem:   ss,
			name:       name,
			orbits:     nil,
			satellites: map[string]*body{},
			distanceToCenter: -1,
		}
		ss.bodies[name] = cosmicBody
		return cosmicBody
	}

	cosmicBody = ss.bodies[name] // TODO: without this it doesn't work, need to figure out why

	return cosmicBody
}


func (ss *system) createOrbitBodyPair(bigger, smaller string) {
	bigBody := ss.getOrCreateBody(bigger)
	smallBody := ss.getOrCreateBody(smaller)
	smallBody.orbits = bigBody
	bigBody.satellites[smallBody.name] = smallBody
}

func (ss *system) getTotalOrbits() int {
	orbitCount := 0
	for _, body := range ss.bodies {
		orbitCount += body.getDistanceToCenter()
	}
	return orbitCount
}

func (b *body) getDistanceToCenter() int {
	if b.distanceToCenter != -1 {
		return b.distanceToCenter
	}
	if b.orbits == nil {
		return 0
	}
	return 1 + b.orbits.getDistanceToCenter()
}
