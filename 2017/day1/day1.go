package day1

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)

func doIt(input []int) int {
	sum := 0
	for i, v := range input {
		if i < len(input)-1 && v == input[i+1] {
			sum += v
			continue
		}
		if i == len(input)-1 && v == input[0] {
			sum += v
		}

	}

	return sum
}

func getCheckStepPart2(inputLen, step int) int {
	steps := inputLen / 2
	res := step + steps
	if res > inputLen-1 {
		res = res - inputLen
	}

	return res
}

func doIt2(input []int) int {
	length := len(input)

	sum := 0
	for i, v := range input {
		if v == input[getCheckStepPart2(length, i)] {
			sum += v
		}
	}

	return sum
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2017, 1, "")

	res := doIt(input)

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2017, 1, "")

	res := doIt2(input)

	return fmt.Sprintf("Solution: %d", res)
}
