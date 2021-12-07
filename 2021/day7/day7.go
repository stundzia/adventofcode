package day7

import (
	"fmt"

	"github.com/stundzia/adventofcode/utils"
)

func diffSum(nums []int, position int) int {
	res := 0
	for _, num := range nums {
		res += utils.AbsInt(num - position)
	}
	return res
}

func diffSum2(nums []int, position int) int {
	res := 0
	for _, num := range nums {
		res += moveCost(utils.AbsInt(num - position))
	}
	return res
}

func moveCost(moveSize int) int {
	cost := 0
	for i := moveSize; i > 0; i-- {
		cost += i
	}
	return cost
}


func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2021, 7, ",")
	lowest := 999999999
	for i := 0; i < 2000; i++ {
		diff := diffSum(nums, i)
		if diff < lowest {
			lowest = diff
		}
	}
	return fmt.Sprintf("Solution: %d", lowest)
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2021, 7, ",")
	lowest := 999999999
	for i := 0; i < 2000; i++ {
		diff := diffSum2(nums, i)
		if diff < lowest {
			lowest = diff
		}
	}
	return fmt.Sprintf("Solution: %d", lowest)
}
