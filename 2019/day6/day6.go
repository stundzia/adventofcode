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
	input, _ := utils.ReadInputFileContentsAsStringSlice(2019, 6, "\n")
	ss := NewSystemFromInput(input)
	_, s := ss.findCommonBody("YOU", "SAN")
	return fmt.Sprintf("Solution: %d", s - 2)
}
