package day11

import (
	"fmt"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

type universe struct {
	raw       [][]string
	galaxyMap map[string]galaxy
}

type galaxy struct {
	id string
	x  int
	y  int
}

func (u *universe) expand() {
	lenY := len(u.raw)
	for y := 0; y < lenY; y++ {
		if utils.SliceStrAllEqual(u.raw[y], ".") {
			u.raw = insertY(u.raw, y, u.raw[y])
			y++
			lenY = len(u.raw)
		}
	}

	lenX := len(u.raw[0])
colLoop:
	for x := 0; x < lenX; x++ {
		for _, n := range u.raw {
			if n[x] != "." {
				continue colLoop
			}
		}
		for y, _ := range u.raw {
			u.raw[y] = insertX(u.raw[y], x, ".")
		}
		lenX++
		x++
	}
}

func (u *universe) expandN(expansions int) {
	lenY := len(u.raw)
	for y := 0; y < lenY; y++ {
		if utils.SliceStrAllEqual(u.raw[y], ".") {
			for t := 0; t < expansions; t++ {
				u.raw = insertY(u.raw, y, u.raw[y])
				y++
				lenY = len(u.raw)
			}
		}
	}

	lenX := len(u.raw[0])
colLoop:
	for x := 0; x < lenX; x++ {
		for _, n := range u.raw {
			if n[x] != "." {
				continue colLoop
			}
		}
		for t := 0; t < expansions; t++ {
			for y, _ := range u.raw {
				u.raw[y] = insertX(u.raw[y], x, ".")
			}
			lenX++
			x++
		}
	}
}

func (u *universe) print() {
	fmt.Println("lenx: ", len(u.raw[0]))
	fmt.Println("leny: ", len(u.raw))
	for _, l := range u.raw {
		fmt.Println(strings.Join(l, ""))
	}
}

func (u *universe) parseGalaxyMap() {
	for y, l := range u.raw {
		for x, n := range l {
			if n == "#" {
				u.galaxyMap[utils.CoordsIntsToStr(x, y)] = galaxy{
					id: utils.CoordsIntsToStr(x, y),
					x:  x,
					y:  y,
				}
			}
		}
	}
}

func (u *universe) getGalaxyPairs() [][2]galaxy {
	galaxies := []galaxy{}
	pairs := [][2]galaxy{}
	for _, v := range u.galaxyMap {
		galaxies = append(galaxies, v)
	}
	for i, g := range galaxies {
		for _, g2 := range galaxies[i:] {
			pairs = append(pairs, [2]galaxy{g, g2})
		}
	}

	return pairs
}

func distance(galA, galB galaxy) int {
	return utils.AbsInt(galA.x-galB.x) + utils.AbsInt(galB.y-galA.y)
}

func insertX(row []string, index int, value string) []string {
	if len(row) == index {
		return append(row, value)
	}
	row = append(row[:index+1], row[index:]...)
	row[index] = value
	return row
}

func insertY(grid [][]string, index int, value []string) [][]string {
	if len(grid) == index {
		return append(grid, value)
	}
	grid = append(grid[:index+1], grid[index:]...)
	grid[index] = value
	return grid
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringGrid(2023, 11)
	u := &universe{raw: input, galaxyMap: map[string]galaxy{}}
	u.expand()
	u.parseGalaxyMap()
	pairs := u.getGalaxyPairs()

	res := 0

	for _, pair := range pairs {
		res += distance(pair[0], pair[1])
	}

	return fmt.Sprintf("Solution: %d", res)
}

func getResWithNExpands(expands int) int {
	input, _ := utils.ReadInputFileContentsAsStringGrid(2023, 11)
	u := &universe{raw: input, galaxyMap: map[string]galaxy{}}
	u.expandN(expands)
	u.parseGalaxyMap()
	pairs := u.getGalaxyPairs()

	res := 0
	for _, pair := range pairs {
		res += distance(pair[0], pair[1])
	}
	return res
}

func DoGold() string {
	diff := getResWithNExpands(1) - getResWithNExpands(0)

	return fmt.Sprintf("Solution: %d", getResWithNExpands(0)+(diff*999999))
}
