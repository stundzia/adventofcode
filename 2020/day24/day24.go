package day24

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 24, "\n")
	floor := NewFloor()
	for _, path := range input {
		floor.TraverseTiles(path).Flip()
	}
	return fmt.Sprintf("%d", floor.BlackTileCount)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 24, "\n")
	return fmt.Sprintf("%d", len(input))
}
