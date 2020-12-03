package day3

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strings"
)


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 3, "\n")
	for i, inp := range input {
		input[i] = strings.Repeat(inp, (len(input)/len(inp)) * 4)
	}
	mount := Mountain{}
	mount.ParseMap(input)
	mount.CurrentPosition = []int{0,0}
	mount.GoDownToBottom(3, 1)
	return fmt.Sprintf("%d", mount.TreesEncountered)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 3, "\n")
	for i, inp := range input {
		input[i] = strings.Repeat(inp, (len(input)/len(inp)) * 8)
	}
	mount := Mountain{}
	mount.ParseMap(input)
	mount.CurrentPosition = []int{0,0}

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	treeCounts := make([]int, len(slopes))

	for i, slope := range slopes {
		mount.GoDownToBottom(slope[0], slope[1])
		treeCounts[i] = mount.TreesEncountered
		mount.Reset()
	}

	res := 1
	for _, tc := range treeCounts {
		res *= tc
	}
	return fmt.Sprintf("%d", res)
}