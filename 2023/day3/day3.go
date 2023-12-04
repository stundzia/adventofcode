package day3

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"math"
	"strconv"
)

type partNum struct {
	val  int
	xMin int
	xMax int
	y    int
}

func (pn partNum) isAdjacent(x, y int) bool {
	if utils.AbsInt(x-pn.xMin) <= 1 || utils.AbsInt(x-pn.xMax) <= 1 {
		return utils.AbsInt(y-pn.y) <= 1
	}
	return false
}

func isConsidered(grid [][]string, x, y int) bool {
	toCheck := [][]int{}
	for xDelta := -1; xDelta < 2; xDelta++ {
		for yDelta := -1; yDelta < 2; yDelta++ {
			toCheck = append(toCheck, []int{x + xDelta, y + yDelta})
		}
	}

	for _, tc := range toCheck {
		if tc[1] > len(grid)-1 || tc[0] < 0 || tc[0] > len(grid[0])-1 || tc[1] < 0 {
			continue
		}
		if grid[tc[1]][tc[0]] == "." {
			continue
		}
		if _, err := strconv.Atoi(grid[tc[1]][tc[0]]); err == nil {
			continue
		}
		return true
	}
	return false
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringGrid(2023, 3)

	res := 0
	for y, l := range input {
		nums := []int{}
		considered := false
		for x, c := range l {
			cNum, err := strconv.Atoi(c)
			if err == nil {
				if isConsidered(input, x, y) {
					considered = true
				}
				nums = append(nums, cNum)
			} else {
				if len(nums) != 0 && considered {
					toAdd := 0
					for i, num := range nums {
						toAdd += int(math.Pow10(len(nums)-i-1)) * num
					}
					res += toAdd
				}
				nums = []int{}
				considered = false
			}
		}
		if len(nums) != 0 && considered {
			toAdd := 0
			for i, num := range nums {
				toAdd += int(math.Pow10(len(nums)-i-1)) * num
			}
			res += toAdd
		}
	}

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringGrid(2023, 3)
	for _, l := range input {
		fmt.Println(l)
	}

	res := 0
	parts := []partNum{}
	for y, l := range input {
		nums := []int{}
		considered := false
		for x, c := range l {
			cNum, err := strconv.Atoi(c)
			if err == nil {
				if isConsidered(input, x, y) {
					considered = true
				}
				nums = append(nums, cNum)
			} else {
				if len(nums) != 0 && considered {
					toAdd := 0
					for i, num := range nums {
						toAdd += int(math.Pow10(len(nums)-i-1)) * num
					}
					parts = append(parts, partNum{
						val:  toAdd,
						xMin: x - len(nums),
						xMax: x - 1,
						y:    y,
					})
					res += toAdd
				}
				nums = []int{}
				considered = false
			}
		}
		if len(nums) != 0 && considered {
			toAdd := 0
			for i, num := range nums {
				toAdd += int(math.Pow10(len(nums)-i-1)) * num
			}
			parts = append(parts, partNum{
				val:  toAdd,
				xMin: len(input[0]) - len(nums),
				xMax: len(input[0]),
				y:    y,
			})
			res += toAdd
		}
	}

	res2 := 0
	for y, l := range input {
		for x, el := range l {
			if el == "*" {
				partsToUse := []int{}
				for _, part := range parts {
					if part.isAdjacent(x, y) {
						partsToUse = append(partsToUse, part.val)
					}
				}
				if len(partsToUse) == 2 {
					res2 += partsToUse[0] * partsToUse[1]
				}
			}
		}
	}

	return fmt.Sprintf("Solution: %d", res2)
}
