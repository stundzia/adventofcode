package day14

import (
	"fmt"
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
	for i, c := range input {
		computer.parseInputCommand(c)
		fmt.Println(i, " Done")
	}
	return strconv.Itoa(computer.MemorySum())
}
