package day3

import (
	"github.com/stundzia/adventofcode/utils"
	"strconv"
	"strings"
)

func getWireMoves(wires []string) (movesFirstWire, movesSecondWire []string) {
	movesFirstWire = strings.Split(wires[0], ",")
	movesSecondWire = strings.Split(wires[1], ",")
	return
}

func DoSilver() string {
	wires, _ := utils.ReadInputFileContentsAsStringSlice(2019, 3, "\n")
	movesA, movesB := getWireMoves(wires)
	grid := NewGrid()
	wireA := &Wire{
		signature: 1,
		grid:      grid,
		x:         0,
		y:         0,
	}
	wireB := &Wire{
		signature: 2,
		grid:      grid,
		x:         0,
		y:         0,
	}
	for _, mov := range movesA {
		wireA.handleCommandInGridMap(mov)
	}
	for _, mov := range movesB {
		wireB.handleCommandInGridMap(mov)
	}
	lowestDist := -1
	for coords, mark := range grid.gridMap {
		if mark == 3 {
			dist := grid.getDistanceToCenter(coords[0], coords[1])
			if lowestDist == -1 {
				lowestDist = dist
			}
			if dist < lowestDist {
				lowestDist = dist
			}
		}
	}
	return strconv.Itoa(lowestDist)
}

func DoGold() string {
	wires, _ := utils.ReadInputFileContentsAsStringSlice(2019, 3, "\n")
	movesA, movesB := getWireMoves(wires)
	grid := NewGrid()
	wireA := &Wire{
		signature: 1,
		grid:      grid,
		x:         0,
		y:         0,
	}
	wireB := &Wire{
		signature: 2,
		grid:      grid,
		x:         0,
		y:         0,
	}
	for _, mov := range movesA {
		wireA.handleCommandInGridMapWithSteps(mov)
	}
	for _, mov := range movesB {
		wireB.handleCommandInGridMapWithSteps(mov)
	}
	lowestSteps := -1
	for coords, mark := range grid.gridMap {
		if mark == 3 {
			steps := grid.gridMapIntersectionSteps[[2]int{coords[0], coords[1]}]
			if lowestSteps == -1 {
				lowestSteps = steps
			}
			if steps < lowestSteps {
				lowestSteps = steps
			}
		}
	}
	return strconv.Itoa(lowestSteps)
}