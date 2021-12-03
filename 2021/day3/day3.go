package day3

import (
	"fmt"
	"strconv"

	"github.com/stundzia/adventofcode/utils"
)

func s(nums []string) int {
	var gamma string
	var epsilon string
	for i := 0; i < len(nums[0]); i ++ {
		ones := 0
		zeros := 0
		for _, n := range nums {
			switch n[i] {
			case '1':
				ones++
			case '0':
				zeros++
			}
		}
		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(g * e)
}

func s2(nums []string) int {
	validO2 := nums
	validCO2 := nums

	var O2 string
	var CO2 string

	for i := 0; i < len(nums[0]); i ++ {
		ones := []string{}
		zeros := []string{}
		for _, n := range validO2 {
			switch n[i] {
			case '1':
				ones = append(ones, n)
			case '0':
				zeros = append(zeros, n)
			}
		}
		if len(zeros) > len(ones) {
			validO2 = zeros
		} else {
			validO2 = ones
		}
		if len(validO2) == 1 {
			O2 = validO2[0]
			break
		}
	}

	for i := 0; i < len(nums[0]); i ++ {
		ones := []string{}
		zeros := []string{}
		for _, n := range validCO2 {
			switch n[i] {
			case '1':
				ones = append(ones, n)
			case '0':
				zeros = append(zeros, n)
			}
		}
		if len(zeros) > len(ones) {
			validCO2 = ones
		} else {
			validCO2 = zeros
		}
		if len(validCO2) == 1 {
			CO2 = validCO2[0]
			break
		}
	}

	g, _ := strconv.ParseInt(O2, 2, 64)
	e, _ := strconv.ParseInt(CO2, 2, 64)
	return int(g * e)
}

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 3, "\n")
	return fmt.Sprintf("Solution: %d", s(nums))
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 3, "\n")
	return fmt.Sprintf("Solution: %d", s2(nums))
}
