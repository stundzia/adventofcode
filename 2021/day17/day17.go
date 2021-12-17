package day17

import (
	"fmt"

	"github.com/stundzia/adventofcode/utils"
)

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 17, "\n")
	return fmt.Sprintf("Solution: %d", len(nums))
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 17, "\n")
	return fmt.Sprintf("Solution: %d", len(nums))
}
