package day10

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

type pipeLand struct {
	input    [][]string
	pipesMap map[string]*pipePart
}

type pipePart struct {
	pipeType   string
	sType      string
	x          int
	y          int
	xy         string
	partOfLoop bool
	inLoop     bool
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

func (pl *pipeLand) calcIfInLoop(pp *pipePart) bool {
	hLeft := 0
	for x := 0; x < pp.x; x++ {
		if n := pl.pipesMap[utils.CoordsIntsToStr(x, pp.y)]; n.partOfLoop {
			if n.pipeType == "|" || n.sType == "|" || n.pipeType == "L" || n.sType == "L" || n.pipeType == "J" || n.sType == "J" {
				if x < pp.x {
					hLeft++
				}
				if x > pp.x {
					hLeft++
				}
			}
		}
	}
	if hLeft%2 == 1 {
		pp.inLoop = true
	}

	return pp.inLoop
}

func (pl *pipeLand) moveThrougPipeLoopP2() int {
	var start *pipePart
	for _, pp := range pl.pipesMap {
		if pp.pipeType == "S" {
			start = pp
			break
		}
	}

	next := start.getNext(start)
	if start.n != nil && start.s != nil {
		start.sType = "|"
	}
	if start.e != nil && start.w != nil {
		start.sType = "-"
	}
	if start.e != nil && start.n != nil {
		start.sType = "L"
	}
	if start.w != nil && start.n != nil {
		start.sType = "J"
	}

	stepsTaken := 1
	last := start
	current := next

	for current != start {
		current.partOfLoop = true
		stepsTaken++
		previous := current
		current = current.getNext(last)
		last = previous
	}

	for coords, pp := range pl.pipesMap {
		x, y := utils.CoordsStrToInts(coords)
		if pp.pipeType == "S" {
			pl.input[y][x] = "S"
			continue
		}
		if pp.partOfLoop {
			pl.input[y][x] = "*"
		} else {
			if pl.calcIfInLoop(pp) {
				pl.input[y][x] = "I"
				continue
			}
			pl.input[y][x] = "."
		}
	}
	f, err := os.Create("visual_dump.txt")
	if err != nil {
		log.Fatal("oh bugger")
	}
	for _, l := range pl.input {
		f.WriteString(strings.Join(l, "") + "\n")
	}

	res := 0
	for _, pp := range pl.pipesMap {
		if pp.inLoop {
			res++
		}
	}

	return res
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
	input, _ := utils.ReadInputFileContentsAsStringGrid(2023, 10)

	pl := parsePipeland(input)
	res := pl.moveThrougPipeLoopP2()

	// 357 < res < 752
	return fmt.Sprintf("Solution: %d", res)
}
