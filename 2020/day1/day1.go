package day1

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)

func getProductOfMatchingSumOfTwoItems(intSlice []int, sum int) int {
	for i, num := range intSlice {
		for t, num2 := range intSlice {
			if i == t {
				continue
			}
			if num+num2 == sum {
				return num * num2
			}
		}
	}
	return -1
}

func getProductOfMatchingSumOfThreeItems(intSlice []int, sum int) int {
	for i, num := range intSlice {
		for t, num2 := range intSlice {
			for z, num3 := range intSlice {
				if i == t || t == z || i == z {
					continue
				}
				if num+num2+num3 == sum {
					return num * num2 * num3
				}
			}
		}
	}
	return -1
}

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2020, 1, "\n")
	return fmt.Sprintf("Solution: %d", getProductOfMatchingSumOfTwoItems(nums, 2020))
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2020, 1, "\n")
	return fmt.Sprintf("Solution: %d", getProductOfMatchingSumOfThreeItems(nums, 2020))
}
