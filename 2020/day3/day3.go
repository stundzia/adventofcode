package day3

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 3, "\n")
	mount := GetMountainFromMap(input)
	mount.GoDownToBottom(3, 1)
	return fmt.Sprintf("%d", mount.TreesEncountered)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 3, "\n")
	mount := GetMountainFromMap(input)

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
