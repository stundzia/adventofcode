package day2

import (
	"fmt"
	"github.com/stundzia/adventofcode/2019/computer"
	"github.com/stundzia/adventofcode/utils"
	"strconv"
)

func DoSilver() string {
	opcodes, _ := utils.ReadInputFileContentsAsIntSlice(2019, 2, ",")
	opcodes[1] = 12
	opcodes[2] = 2
	c := computer.NewComputer(opcodes)
	res, _ := c.Run()
	return strconv.Itoa(res)
}

func DoGold() string {
	neededOutput := 19690720
	opcodes, _ := utils.ReadInputFileContentsAsIntSlice(2019, 2, ",")
	for noun := 0; noun < 1000; noun++ {
		for verb := 0; verb < 1000; verb++ {
			c := computer.NewComputer(opcodes)
			c.Opcodes[1] = noun
			c.Opcodes[2] = verb
			res, err := c.Run()
			if res == neededOutput && err == nil {
				return fmt.Sprintf("%d%d", noun, verb)
			}
		}
	}
	return "fuckknows"
}