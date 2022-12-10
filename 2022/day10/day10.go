package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

type cpu struct {
	register int
	cycle    int
	res1     int
	res2     map[int]string
}

func (c *cpu) printres2() {
	res := ""
	for i := 0; ; i++ {
		if _, found := c.res2[i]; !found {
			break
		}
		if i > 0 && i%40 == 0 {
			res += "\n"
		}
		res += c.res2[i]
	}

	fmt.Println(res)
}

func (c *cpu) cycleup() {
	position := c.cycle % 40
	if position >= c.register-1 && position <= c.register+1 {
		c.res2[c.cycle] = "#"
	} else {
		c.res2[c.cycle] = "."
	}
	c.cycle++
	if c.cycle == 20 || (c.cycle > 20 && (c.cycle-20)%40 == 0) {
		c.res1 += c.cycle * c.register
	}
}

func (c *cpu) execute(cmd string) {
	if cmd == "noop" {
		c.cycleup()
		return
	}
	cmdParts := strings.Split(cmd, " ")
	v, _ := strconv.Atoi(cmdParts[1])
	c.cycleup()
	c.cycleup()
	c.register += v
	return
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2022, 10, "\n")
	cpu := &cpu{
		register: 1,
		cycle:    0,
		res2:     map[int]string{},
	}
	for _, cmd := range input {
		cpu.execute(cmd)
	}

	return fmt.Sprintf("Solution: %d", cpu.res1)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2022, 10, "\n")
	cpu := &cpu{
		register: 1,
		cycle:    0,
		res2:     map[int]string{},
	}
	for _, cmd := range input {
		cpu.execute(cmd)
	}
	cpu.printres2()
	return fmt.Sprintf("Solution: ^^")
}
