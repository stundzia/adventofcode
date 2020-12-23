package day23

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2020, 23, "")
	cg := NewCupGame(input)
	for i := 0; i < 100; i++ {
		cg.doMove()
	}
	cg.PrintCups()
	return fmt.Sprintf("%d", len(input))
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2020, 23, "")
	cg := NewCupGame(input)
	for i := 0; i < 100; i++ {
		cg.doMove()
	}
	cg.PrintCups()
	return fmt.Sprintf("%d", len(input))
}
