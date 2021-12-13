package day13

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

type manual struct {
	dots map[string]int
	instructions []string
}

func (m *manual) fold(foldLine string, coord int) {
	switch foldLine {
	case "x":
		for c, v := range m.dots {
			if v > 0 {
				x, y := utils.CoordsStrToInts(c)
				if x < coord {
					continue
				}
				m.dots[c] = 0
				x = coord - (x - coord)
				m.dots[fmt.Sprintf("%d,%d", x,y)] = v
			}
		}
	case "y":
		for c, v := range m.dots {
			if v > 0 {
				x, y := utils.CoordsStrToInts(c)
				if y < coord {
					continue
				}
				m.dots[c] = 0
				y = coord - (y - coord)
				m.dots[fmt.Sprintf("%d,%d", x,y)] = v
			}
		}
	}
}

func (m *manual) doTheInstructions() {
	for _, inst := range m.instructions {
		parts := strings.Split(inst, "=")
		foldType := parts[0]
		coord, _ := strconv.Atoi(parts[1])
		m.fold(foldType, coord)
	}
}

func (m *manual) printView() {
	maxX := 0
	minX := math.MaxInt
	maxY := 0
	minY := math.MaxInt
	for coords, val := range m.dots {
		if val > 0 {
			x, y := utils.CoordsStrToInts(coords)
			if x > maxX {
				maxX = x
			}
			if x < minX {
				minX = x
			}
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}
		}
	}
	lines := []string{}
	for y := minY; y <= maxY; y++ {
		line := ""
		for x := minX; x <= maxX; x++ {
			if val, exists := m.dots[fmt.Sprintf("%d,%d", x, y)]; exists {
				if val > 0 {
					line = line + "#"
					continue
				}
			}
			line = line + "."
		}
		lines = append(lines, line)
	}
	for _, l := range lines {
		fmt.Println(l)
	}
}

func newManual(lines []string) *manual {
	handleCoords := true
	coords := []string{}
	instructions := []string{}
	for _, line := range lines {
		if line == "" {
			handleCoords = false
			continue
		}
		if handleCoords {
			coords = append(coords, line)
		} else {
			instructions = append(instructions, strings.Replace(line, "fold along ", "", -1))
		}
	}
	m := &manual{
		dots: map[string]int{},
		instructions: instructions,
	}
	for _, c := range coords {
		m.dots[c] = 1
	}
	return m
}

func (m *manual) visibleDotCount() int {
	count := 0
	for _, val := range m.dots {
		if val > 0 {
			count++
		}
	}
	return count
}

func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 13, "\n")
	m := newManual(lines)
	m.fold("x", 655)
	return fmt.Sprintf("Solution: %d", m.visibleDotCount())
}

func DoGold() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 13, "\n")
	m := newManual(lines)
	m.doTheInstructions()
	m.printView()
	return fmt.Sprintf("Solution: ^^^")
}
