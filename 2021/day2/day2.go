package day2

import (
	"fmt"

	"github.com/stundzia/adventofcode/utils"
)

func s(ss []string) int {
	sub := &Submarine{
		Depth:    0,
		Position: 0,
	}
	for _, s := range ss {
		sub.handleCommand(s)
	}
	return sub.Position * sub.Depth
}

func s2(ss []string) int {
	sub := &Submarine{
		Depth:    0,
		Position: 0,
		Aim: 0,
	}
	for _, s := range ss {
		sub.handleCommandV2(s)
	}
	return sub.Position * sub.Depth
}

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 2, "\n")
	return fmt.Sprintf("Solution: %d", s(nums))
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 2, "\n")
	return fmt.Sprintf("Solution: %d", s2(nums))
}
