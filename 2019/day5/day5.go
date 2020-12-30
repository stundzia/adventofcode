package day5

import (
	"fmt"
	"github.com/stundzia/adventofcode/2019/computer"
	"github.com/stundzia/adventofcode/utils"
)


func DoSilver() string {
	opcodes, _ := utils.ReadInputFileContentsAsIntSlice(2019, 5, ",")
	comp := computer.NewComputer(opcodes)
	go comp.Run()
	comp.InputPipe <- 1
	var lastOutput int
	for ;; {
		lastOutput = <- comp.OutputPipe
		if !comp.Running {
			break
		}
	}
	return fmt.Sprintf("%d", lastOutput)
}

func DoGold() string {
	opcodes, _ := utils.ReadInputFileContentsAsIntSlice(2019, 5, ",")
	comp := computer.NewComputer(opcodes)
	go comp.Run()
	comp.InputPipe <- 5
	var lastOutput int
	for ;; {
		lastOutput = <- comp.OutputPipe
		if !comp.Running {
			break
		}
	}
	return fmt.Sprintf("%d", lastOutput)
}
