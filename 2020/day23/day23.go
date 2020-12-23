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
	n := utils.GetMaxFromIntSlice(input) + 1
	for ;len(input) < 1000000; n++ {
		input = append(input, n)
	}
	cg := NewCupGame(input)
	for i := 0; i < 10000000; i++ {
		cg.doMove()
		if i % 1000000 == 0 {
			fmt.Println("moves done: ", i)
		}
	}
	// 921194976570 is too high
	return fmt.Sprintf("%d", cg.GetPart2Res())
}
