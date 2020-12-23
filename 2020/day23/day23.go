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
	n := utils.GetMaxFromIntSlice(input)
	for ;len(input) < 1000000; n++ {
		input = append(input, n)
	}
	cg := NewCupGame(input)
	for i := 0; i < 10000000; i++ {
		cg.doMove()
		if i % 10000 == 0 {
			fmt.Println("moves done: ", i)
		}
	}
	cg.PrintPart2Res()
	return fmt.Sprintf("%d", len(input))
}
