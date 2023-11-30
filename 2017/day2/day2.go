package day2

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"log"
)

func lineChecksumP1(line []int) int {
	return utils.GetMaxFromIntSlice(line) - utils.GetMinFromIntSlice(line)
}

func lineChecksumP2(line []int) int {
	for i1, a := range line {
		for i2, b := range line {
			if i1 != i2 && a%b == 0 {
				return a / b
			}
		}
	}

	log.Fatal("FUCK FUCK FUCK WHAT THE FUCK")

	return -1
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsIntGridWithSeparator(2017, 2, "\t")

	res := 0

	for _, line := range input {
		res += lineChecksumP1(line)
	}

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsIntGridWithSeparator(2017, 2, "\t")

	res := 0

	for _, line := range input {
		res += lineChecksumP2(line)
	}

	return fmt.Sprintf("Solution: %d", res)
}
