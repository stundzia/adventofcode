package day9

import (
	"github.com/stundzia/adventofcode/utils"
	"strconv"
)

func isValid(nums []int, n int, preamble int) bool {
	if n < preamble {
		return true
	}
	num := nums[n]
	components := nums[n-preamble : n]
	for i, c := range components {
		for i2, c2 := range components {
			if c+c2 == num && i != i2 {
				return true
			}
		}
	}
	return false
}

func contiguousSet(nums []int, n int, invalidNum int) (found bool, set []int) {
	for ; ; n++ {
		set = append(set, nums[n])
		sum := utils.SumIntSlice(set)
		if sum < invalidNum {
			continue
		} else {
			if sum == invalidNum && len(set) > 1 {
				return true, set
			} else {
				return false, set
			}
		}
	}
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2020, 9, "\n")
	for i, num := range input {
		if !isValid(input, i, 25) {
			return strconv.Itoa(num)
		}
	}
	return strconv.Itoa(-1)
}

func DoGold() string {
	invalidNum, _ := strconv.Atoi(DoSilver())
	input, _ := utils.ReadInputFileContentsAsIntSlice(2020, 9, "\n")
	set := []int{}
	var found bool
	for i, _ := range input {
		if found, set = contiguousSet(input, i, invalidNum); found {
			break
		}
	}
	minSet := set[0]
	maxSet := 0
	for _, num := range set {
		if num < minSet {
			minSet = num
		}
		if num > maxSet {
			maxSet = num
		}
	}

	return strconv.Itoa(minSet + maxSet)
}
