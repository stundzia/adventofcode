package day1

import (
	"fmt"
	"math"

	"github.com/stundzia/adventofcode/utils"
)

func swapMinIfHigher(cals [3]int, newCals int) [3]int {
	min := math.MaxInt
	minIndex := 0
	for i, c := range cals {
		if c < min {
			min = c
			minIndex = i
		}
	}
	if newCals > min {
		cals[minIndex] = newCals
	}

	return cals
}

func DoSilver() string {
	elfCals, _ := utils.ReadInputFileContentsAsIntSliceLines(2022, 1)
	maxCals := 0
	for _, cals := range elfCals {
		if utils.SumIntSlice(cals) > maxCals {
			maxCals = utils.SumIntSlice(cals)
		}
	}

	return fmt.Sprintf("Solution: %d", maxCals)
}

func DoGold() string {
	elfCals, _ := utils.ReadInputFileContentsAsIntSliceLines(2022, 1)
	topCals := [3]int{0, 0, 0}
	for _, cals := range elfCals {
		topCals = swapMinIfHigher(topCals, utils.SumIntSlice(cals))
	}

	return fmt.Sprintf("Solution: %d", topCals[0]+topCals[1]+topCals[2])
}
