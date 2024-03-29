package day19

import (
	"fmt"

	"github.com/stundzia/adventofcode/utils"
)

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 18, "\n")
	return fmt.Sprintf("Solution: %d", len(nums))
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 18, "\n")
	return fmt.Sprintf("Solution: %d", len(nums))
}
