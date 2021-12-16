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
	return cg.GetPart1Res()
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2020, 23, "")
	n := utils.GetMaxFromIntSlice(input) + 1
	for ; len(input) < 1000000; n++ {
		input = append(input, n)
	}
	cg := NewCupGame(input)
	for i := 0; i < 10000000; i++ {
		cg.doMove()
	}
	return fmt.Sprintf("%d", cg.GetPart2Res())
}
