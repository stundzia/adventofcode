package day1

import (
	"fmt"

	"github.com/stundzia/adventofcode/utils"
)

func s(nums []int) int {
	prev := 0
	count := 0
	for _, i := range nums {
		if i > prev {
			count++
		}
		prev = i
	}
	count--
	return count
}

func s2(nums []int) int {
	count := 0
	lastWindow := 0
	for i, _ := range nums {
		if i + 3 > len(nums) {
			break
		}
		currWindow := nums[i] + nums[i + 1] + nums[i + 2]
		if currWindow > lastWindow {
			count++
		}
		lastWindow = currWindow
	}
	count--
	return count
}

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2021, 1, "\n")
	return fmt.Sprintf("Solution: %d", s(nums))
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2021, 1, "\n")
	return fmt.Sprintf("Solution: %d", s2(nums))
}
