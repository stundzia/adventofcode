package day17

import "fmt"

type PocketDimension struct {
	Cubes map[[3]int]*Cube
}

type Cube struct {
	dimension *PocketDimension
	coords [3]int
	active bool
}

func NewPockedDimensionFromInitialStateSlice(stateString []string) *PocketDimension {
	pd := &PocketDimension{Cubes: map[[3]int]*Cube{}}
	for y, line := range stateString {
		for x, val := range line {
			if string(val) == "#" {
				pd.CreateCube([3]int{x,y,0}, true)
			} else {
				pd.CreateCube([3]int{x,y,0}, false)
			}
		}
	}
	return pd
}


func (pd *PocketDimension) DoCycle() {
	for _, cube := range pd.Cubes {
		cube.GetNeighbours()
	}
	cubes := []*Cube{}
	makeActive := []int{}
	makeInactive := []int{}
	for _, cube := range pd.Cubes {
		cubes = append(cubes, cube)
	}

	for i, cube := range cubes {
		nactive := cube.GetActiveNeighbourCount()
		if cube.active && nactive != 2 && nactive != 3 {
			makeInactive = append(makeInactive, i)
			continue
		}
		if !cube.active && nactive == 3 {
			makeActive = append(makeActive, i)
		}
	}
	for _, activeIndex := range makeActive {
		cubes[activeIndex].active = true
	}
	for _, inactiveIndex := range makeInactive {
		cubes[inactiveIndex].active = false
	}
}

func (pd *PocketDimension) CreateCube(coords [3]int, active bool) *Cube {
	cube := &Cube{
		dimension: pd,
		coords:    coords,
		active:    active,
	}
	pd.Cubes[coords] = cube
	return cube
}


func (pd *PocketDimension) GetActiveCount() (activeCount int) {
	for _, c := range pd.Cubes {
		if c.active {
			activeCount++
		}
	}
	fmt.Println(activeCount)
	return activeCount
}


func (pd *PocketDimension) GetOrCreateCube(coords [3]int) *Cube {
	if cube, ok := pd.Cubes[coords]; !ok {
		return pd.CreateCube(coords, false)
	} else {
		return cube
	}
}

func (c *Cube) GetNeighbours() []*Cube {
	neighbours := []*Cube{}
	coords := c.coords
	deltas := [3]int{-1, 0, 1}
	for _, dx := range deltas {
		for _, dy := range deltas {
			for _, dz := range deltas {
				dcoords := [3]int{coords[0] + dx, coords[1] + dy, coords[2] + dz}
				if dcoords != coords {
					neighbours = append(neighbours, c.dimension.GetOrCreateCube(dcoords))
				}
			}
		}
	}
	return neighbours
}

func (c *Cube) GetActiveNeighbourCount() (count int) {
	neighbours := c.GetNeighbours()
	for _, neighbour := range neighbours {
		if neighbour.active {
			count++
		}
	}
	return count
}

