package day10

import (
	"fmt"
	"log"

	"github.com/stundzia/adventofcode/utils"
)

type pipeLand struct {
	input    [][]string
	pipesMap map[string]*pipePart
}

type pipePart struct {
	pipeType   string
	x          int
	y          int
	xy         string
	partOfLoop bool
	n          *pipePart
	s          *pipePart
	w          *pipePart
	e          *pipePart
}

func parsePipeland(input [][]string) *pipeLand {
	pl := &pipeLand{
		input:    input,
		pipesMap: map[string]*pipePart{},
	}
	for y, l := range input {
		for x, pp := range l {
			pl.pipesMap[utils.CoordsIntsToStr(x, y)] = &pipePart{
				pipeType:   pp,
				x:          x,
				y:          y,
				partOfLoop: pp == "S",
				xy:         utils.CoordsIntsToStr(x, y),
			}
		}
	}
	for _, pp := range pl.pipesMap {
		switch pp.pipeType {
		case "|":
			pp.n = pl.pipesMap[utils.CoordsIntsToStr(pp.x, pp.y-1)]
			pp.s = pl.pipesMap[utils.CoordsIntsToStr(pp.x, pp.y+1)]
		case "-":
			pp.e = pl.pipesMap[utils.CoordsIntsToStr(pp.x+1, pp.y)]
			pp.w = pl.pipesMap[utils.CoordsIntsToStr(pp.x-1, pp.y)]
		case "L":
			pp.n = pl.pipesMap[utils.CoordsIntsToStr(pp.x, pp.y-1)]
			pp.e = pl.pipesMap[utils.CoordsIntsToStr(pp.x+1, pp.y)]
		case "J":
			pp.n = pl.pipesMap[utils.CoordsIntsToStr(pp.x, pp.y-1)]
			pp.w = pl.pipesMap[utils.CoordsIntsToStr(pp.x-1, pp.y)]
		case "7":
			pp.s = pl.pipesMap[utils.CoordsIntsToStr(pp.x, pp.y+1)]
			pp.w = pl.pipesMap[utils.CoordsIntsToStr(pp.x-1, pp.y)]
		case "F":
			pp.s = pl.pipesMap[utils.CoordsIntsToStr(pp.x, pp.y+1)]
			pp.e = pl.pipesMap[utils.CoordsIntsToStr(pp.x+1, pp.y)]
		}
		if pp.n != nil && pp.n.pipeType == "S" {
			pp.partOfLoop = true
			pp.n.s = pp
		}
		if pp.s != nil && pp.s.pipeType == "S" {
			pp.partOfLoop = true
			pp.s.n = pp
		}
		if pp.e != nil && pp.e.pipeType == "S" {
			pp.partOfLoop = true
			pp.e.w = pp
		}
		if pp.w != nil && pp.w.pipeType == "S" {
			pp.partOfLoop = true
			pp.w.e = pp
		}

	}

	return pl
}

func (pl *pipeLand) moveThrougPipeLoop() int {
	var start *pipePart
	for _, pp := range pl.pipesMap {
		if pp.pipeType == "S" {
			start = pp
			break
		}
	}

	next := start.getNext(start)

	stepsTaken := 1
	last := start
	current := next

	for current != start {
		stepsTaken++
		previous := current
		current = current.getNext(last)
		last = previous
	}

	return stepsTaken
}

func (pp *pipePart) getNext(last *pipePart) *pipePart {
	if pp.n != nil && pp.n != last {
		return pp.n
	} else if pp.s != nil && pp.s != last {
		return pp.s
	} else if pp.e != nil && pp.e != last {
		return pp.e
	} else if pp.w != nil && pp.w != last {
		return pp.w
	}
	log.Fatal("well wtf")
	return nil
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringGrid(2023, 10)

	pl := parsePipeland(input)
	steps := pl.moveThrougPipeLoop()
	fmt.Println(steps)

	res := steps / 2

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {

	res := 0

	return fmt.Sprintf("Solution: %d", res)
}
