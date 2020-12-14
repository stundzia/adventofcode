package day14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Computer struct {
	Memory map[int]int64
	Mask1 int64
	Mask0 int64
}

var commandRegex = regexp.MustCompile("([a-z]{3,4})(\\[\\d+\\])?\\s=\\s(.*)")


func NewComputer() *Computer {
	c := &Computer{
		Memory: map[int]int64{},
		Mask1:   0,
		Mask0:   0,
	}
	return c
}

func (c *Computer) handleMaskCommand(value string) {
	mask1Str := strings.Replace(value, "X", "0", -1)
	mask1, _ := strconv.ParseInt(mask1Str, 2, 64)
	c.Mask1 = mask1
	mask0Str := strings.Replace(value, "1", "X", -1)
	mask0Str = strings.Replace(mask0Str, "0", "1", -1)
	mask0Str = strings.Replace(mask0Str, "X", "0", -1)
	mask0, _ := strconv.ParseInt(mask0Str, 2, 64)
	c.Mask1 = mask1
	c.Mask0 = mask0
}

func (c *Computer) putToMemory(address int, value int64) {
	val := value | c.Mask1
	val = val & ^c.Mask0
	c.Memory[address] = val
}

func (c *Computer) handleMemoryCommand(address, value string) {
	// Lazy, I know
	address = strings.Replace(address, "[", "", -1)
	address = strings.Replace(address, "]", "", -1)
	addressInt, _ := strconv.Atoi(address)
	valueInt, _ := strconv.Atoi(value)
	c.putToMemory(addressInt, int64(valueInt))
}

func (c *Computer) parseInputCommand(command string) {
	parts := commandRegex.FindAllStringSubmatch(command, -1)
	switch parts[0][1] {
	case "mask":
		c.handleMaskCommand(parts[0][3])
		break
	case "mem":
		c.handleMemoryCommand(parts[0][2], parts[0][3])
		break
	default:
		fmt.Println("Unknown command: ", parts[0][1])
		break
	}
}

func (c *Computer) MemorySum() int {
	var sum int64
	for _, val := range c.Memory {
		sum += val
	}
	return int(sum)
}