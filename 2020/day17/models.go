package day17

import "fmt"

type PocketDimension struct {
	Cubes   map[[3]int]*Cube
	Cubes4D map[[4]int]*Cube
}

type Cube struct {
	dimension *PocketDimension
	coords    [3]int
	coords4D  [4]int
	active    bool
}

func NewPockedDimensionFromInitialStateSlice(stateString []string, dimensions int) *PocketDimension {
	pd := &PocketDimension{Cubes: map[[3]int]*Cube{}}
	if dimensions == 4 {
		pd.Cubes4D = map[[4]int]*Cube{}
	}
	if dimensions == 4 {
		for y, line := range stateString {
			for x, val := range line {
				if string(val) == "#" {
					pd.CreateCube4D([4]int{x, y, 0, 0}, true)
				} else {
					pd.CreateCube4D([4]int{x, y, 0, 0}, false)
				}
			}
		}
		for _, cube := range pd.Cubes4D {
			cube.GetNeighbours4D()
		}
	} else {
		for y, line := range stateString {
			for x, val := range line {
				if string(val) == "#" {
					pd.CreateCube([3]int{x, y, 0}, true)
				} else {
					pd.CreateCube([3]int{x, y, 0}, false)
				}
			}
		}
		for _, cube := range pd.Cubes {
			cube.GetNeighbours()
		}
	}
	return pd
}

func (pd *PocketDimension) DoCycle() {
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

func (pd *PocketDimension) Do4DCycle() {
	cubes := []*Cube{}
	makeActive := []int{}
	makeInactive := []int{}
	for _, cube := range pd.Cubes4D {
		cubes = append(cubes, cube)
	}

	for i, cube := range cubes {
		nactive := cube.GetActive4DNeighbourCount()
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

func (pd *PocketDimension) CreateCube4D(coords [4]int, active bool) *Cube {
	cube := &Cube{
		dimension: pd,
		coords4D:  coords,
		active:    active,
	}
	pd.Cubes4D[coords] = cube
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

func (pd *PocketDimension) Get4DActiveCount() (activeCount int) {
	for _, c := range pd.Cubes4D {
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

func (pd *PocketDimension) GetOrCreateCube4D(coords [4]int) *Cube {
	if cube, ok := pd.Cubes4D[coords]; !ok {
		return pd.CreateCube4D(coords, false)
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

func (c *Cube) GetNeighbours4D() []*Cube {
	neighbours := []*Cube{}
	coords := c.coords4D
	deltas := [3]int{-1, 0, 1}
	for _, dx := range deltas {
		for _, dy := range deltas {
			for _, dz := range deltas {
				for _, dw := range deltas {
					dcoords := [4]int{coords[0] + dx, coords[1] + dy, coords[2] + dz, coords[3] + dw}
					if dcoords != coords {
						neighbours = append(neighbours, c.dimension.GetOrCreateCube4D(dcoords))
					}
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

func (c *Cube) GetActive4DNeighbourCount() (count int) {
	neighbours := c.GetNeighbours4D()
	for _, neighbour := range neighbours {
		if neighbour.active {
			count++
		}
	}
	return count
}
