package day6

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2019, 6, "\n")
	ss := NewSystemFromInput(input)
	return fmt.Sprintf("%d", ss.getTotalOrbits())
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2019, 6, "\n")
	return fmt.Sprintf("Solution: %d", len(input))
}
