package day14

import (
	"github.com/stundzia/adventofcode/utils"
	"strconv"
)

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 14, "\n")
	computer := NewComputer(1)
	for _, c := range input {
		computer.parseInputCommand(c)
	}
	return strconv.Itoa(computer.MemorySum())
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 14, "\n")
	computer := NewComputer(2)
	for _, c := range input {
		computer.parseInputCommand(c)
	}
	return strconv.Itoa(computer.MemorySum())
}
