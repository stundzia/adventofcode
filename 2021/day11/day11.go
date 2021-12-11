package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

type octoGrid struct {
	flashCount int
	octopi map[string]*octopus
}

type octopus struct {
	octoGrid *octoGrid
	x int
	y int
	energy int
	adjacent map[string]*octopus
	flashed bool
}

func (og *octoGrid) assignNeighbours() {
	for coords, octo := range og.octopi {
		x, y := utils.CoordsStrToInts(coords)
		if nei, exists := og.octopi[fmt.Sprintf("%d,%d", x+1, y)]; exists {
			octo.adjacent[fmt.Sprintf("%d,%d", x+1,y)] = nei
		}
		if nei, exists := og.octopi[fmt.Sprintf("%d,%d", x+1, y+1)]; exists {
			octo.adjacent[fmt.Sprintf("%d,%d", x+1,y+1)] = nei
		}
		if nei, exists := og.octopi[fmt.Sprintf("%d,%d", x+1, y-1)]; exists {
			octo.adjacent[fmt.Sprintf("%d,%d", x+1,y-1)] = nei
		}
		if nei, exists := og.octopi[fmt.Sprintf("%d,%d", x-1, y)]; exists {
			octo.adjacent[fmt.Sprintf("%d,%d", x-1,y)] = nei
		}
		if nei, exists := og.octopi[fmt.Sprintf("%d,%d", x-1, y+1)]; exists {
			octo.adjacent[fmt.Sprintf("%d,%d", x-1,y+1)] = nei
		}
		if nei, exists := og.octopi[fmt.Sprintf("%d,%d", x-1, y-1)]; exists {
			octo.adjacent[fmt.Sprintf("%d,%d", x-1,y-1)] = nei
		}
		if nei, exists := og.octopi[fmt.Sprintf("%d,%d", x, y+1)]; exists {
			octo.adjacent[fmt.Sprintf("%d,%d", x,y+1)] = nei
		}
		if nei, exists := og.octopi[fmt.Sprintf("%d,%d", x, y-1)]; exists {
			octo.adjacent[fmt.Sprintf("%d,%d", x,y-1)] = nei
		}
	}
}

// doStep - handles the logic of a step and returns true if all octopi flashed
func (og *octoGrid) doStep() bool {
	for _, octo := range og.octopi {
		octo.flashed = false
		octo.energy++
	}
	for _, octo := range og.octopi {
		octo.flashIfNeeded()
	}
	for _, octo := range og.octopi {
		if octo.flashed {
			octo.energy = 0
		}
	}
	for _, octo := range og.octopi {
		if octo.flashed {
			continue
		} else {
			return false
		}
	}
	return true
}

func (o *octopus) flashIfNeeded() {
	if o.energy > 9 && o.flashed == false {
		o.flashed = true
		o.octoGrid.flashCount++
		for _, nei := range o.adjacent {
			nei.energy += 1
			nei.flashIfNeeded()
		}
	}
}

func newOctoGrid(nums []string) *octoGrid {
	og := &octoGrid{
		flashCount: 0,
		octopi:     map[string]*octopus{},
	}
	for y, line := range nums {
		valsStr := strings.Split(line, "")
		for x, val := range valsStr {
			valInt, _ := strconv.Atoi(val)
			og.octopi[fmt.Sprintf("%d,%d", x,y)] = &octopus{
				octoGrid: og,
				x:        x,
				y:        y,
				energy: valInt,
				adjacent: map[string]*octopus{},
			}
		}
	}
	og.assignNeighbours()
	return og
}

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 11, "\n")
	og := newOctoGrid(nums)
	for i := 0; i < 100; i++ {
		og.doStep()
	}
	return fmt.Sprintf("Solution: %d",og.flashCount)
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 11, "\n")
	og := newOctoGrid(nums)
	res := 0
	for i := 0; ; i++ {
		if og.doStep() {
			res = i + 1
			break
		}
	}
	return fmt.Sprintf("Solution: %d", res)
}
