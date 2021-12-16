package day1

import (
	"fmt"

	"github.com/stundzia/adventofcode/utils"
)

func countIncreases(nums []int) int {
	prev := -1
	count := 0
	for _, i := range nums {
		if i > prev {
			count++
		}
		prev = i
	}
	// Ignore first increase
	count--
	return count
}

func countWindowIncreases(nums []int) int {
	count := 0
	lastWindow := -1
	for i, _ := range nums {
		if i+3 > len(nums) {
			break
		}
		currWindow := nums[i] + nums[i+1] + nums[i+2]
		if currWindow > lastWindow {
			count++
		}
		lastWindow = currWindow
	}
	// Ignore first increase
	count--
	return count
}

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2021, 1, "\n")
	return fmt.Sprintf("Solution: %d", countIncreases(nums))
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2021, 1, "\n")
	return fmt.Sprintf("Solution: %d", countWindowIncreases(nums))
}
