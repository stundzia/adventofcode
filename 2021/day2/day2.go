package day2

import (
	"fmt"

	"github.com/stundzia/adventofcode/utils"
)

func part1(commands []string) int {
	sub := &submarine{}
	for _, cmd := range commands {
		sub.handleCommand(cmd)
	}
	return sub.Position * sub.Depth
}

func part2(commands []string) int {
	sub := &submarine{}
	for _, cmd := range commands {
		sub.handleCommandV2(cmd)
	}
	return sub.Position * sub.Depth
}

func DoSilver() string {
	commands, _ := utils.ReadInputFileContentsAsStringSlice(2021, 2, "\n")
	return fmt.Sprintf("Solution: %d", part1(commands))
}

func DoGold() string {
	commands, _ := utils.ReadInputFileContentsAsStringSlice(2021, 2, "\n")
	return fmt.Sprintf("Solution: %d", part2(commands))
}
