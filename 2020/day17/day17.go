package day17

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 17, "\n")
	pd := NewPockedDimensionFromInitialStateSlice(input)
	activeCount := 0
	for i := 0; i < 6; i++ {
		pd.DoCycle()
		activeCount = pd.GetActiveCount()
	}
	return fmt.Sprintf("%d", activeCount)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 17, "\n")
	pd := NewPockedDimensionFromInitialStateSlice(input)
	activeCount := 0
	for i := 0; i < 6; i++ {
		pd.DoCycle()
		activeCount = pd.GetActiveCount()
	}
	return fmt.Sprintf("%d", activeCount)
}
