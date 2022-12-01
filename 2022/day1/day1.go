package day1

import (
	"fmt"
	"math"

	"github.com/stundzia/adventofcode/utils"
)

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2022, 1, "\n")
	fmt.Println(nums)
	cals := 0
	maxCals := 0
	for _, num := range nums {
		fmt.Println(num)
		if num == 0 {
			if cals > maxCals {
				maxCals = cals
			}
			cals = 0
			continue
		}
		cals += num
	}
	//inputStr, _ := utils.ReadInputFileContentsAsString(2022, 1)
	//inputStrSlice, _ := utils.ReadInputFileContentsAsStringSlice(2022, 1, ",")
	return fmt.Sprintf("Solution: %d", maxCals)
}

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

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2022, 1, "\n")
	cals := 0
	topCals := [3]int{0, 0, 0}
	for _, num := range nums {
		if num == 0 {
			topCals = swapMinIfHigher(topCals, cals)
			cals = 0
			continue
		}
		cals += num
	}
	return fmt.Sprintf("Solution: %d", topCals[0]+topCals[1]+topCals[2])
}
