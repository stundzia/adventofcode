package day3

import (
	"github.com/stundzia/adventofcode/utils"
	"regexp"
	"strconv"
)

type Grid struct {
	gridMap map[[2]int]uint8
	gridMapIntersectionSteps map[[2]int]int
}

type Wire struct {
	signature uint8
	grid *Grid
	x int
	y int
	stepsTaken int
}

func NewGrid() *Grid {
	grid := &Grid{
		gridMap: map[[2]int]uint8{},
		gridMapIntersectionSteps: map[[2]int]int{},
	}
	return grid
}

func (w *Wire) markPosition(position uint8) uint8 {
	return position | w.signature
}

func parseCommand(command string) (dX, dY, distance int) {
	re := regexp.MustCompile("([A-Z])([0-9]+)")
	res := re.FindAllStringSubmatch(command, -1)
	distance, _ = strconv.Atoi(res[0][2])
	direction := res[0][1]
	switch direction {
	case "U":
		dY = 1
		break
	case "D":
		dY = -1
		break
	case "R":
		dX = 1
		break
	case "L":
		dX = -1
		break
	default:
		break
	}

	return dX, dY, distance
}

func (w *Wire) handleCommandInGridMap(command string) {
	dX, dY, dist := parseCommand(command)
	for i := 0; i < dist; i++ {
		w.x += dX
		w.y += dY
		if val, ok := w.grid.gridMap[[2]int{w.x,w.y}]; ok {
			w.grid.gridMap[[2]int{w.x,w.y}] = w.markPosition(val)
		} else {
			w.grid.gridMap[[2]int{w.x,w.y}] = w.signature
		}
	}
}

func (w *Wire) handleCommandInGridMapWithSteps(command string) {
	dX, dY, dist := parseCommand(command)
	for i := 0; i < dist; i++ {
		w.x += dX
		w.y += dY
		w.stepsTaken++
		if val, ok := w.grid.gridMap[[2]int{w.x,w.y}]; ok {
			w.grid.gridMap[[2]int{w.x,w.y}] = w.markPosition(val)
			if val & w.signature != w.signature {
				w.grid.gridMapIntersectionSteps[[2]int{w.x,w.y}] += w.stepsTaken
			}
		} else {
			w.grid.gridMap[[2]int{w.x,w.y}] = w.signature
			w.grid.gridMapIntersectionSteps[[2]int{w.x,w.y}] = w.stepsTaken
		}
	}
}

func (g *Grid) getDistanceToCenter(x, y int) int {
	return utils.AbsInt(x) + utils.AbsInt(y)
}