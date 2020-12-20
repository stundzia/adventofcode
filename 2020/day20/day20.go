package day20

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 20, "\n\n")
	img := NewImageFromInput(input)
	for _, tile := range img.Tiles {
		tile.FindPotentialPairings()
	}
	cornerIds := []int{}
	for _, tile := range img.Tiles {
		if len(tile.PairingMap) == 2 {
			cornerIds = append(cornerIds, tile.Id)
		}
	}
	res := 1
	for _, cId := range cornerIds {
		res *= cId
	}

	return fmt.Sprintf("%d", res)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 20, "\n\n")
	return fmt.Sprintf("%d", len(input))
}
