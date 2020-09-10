package computer

import (
	"fmt"
)

const opcodeAddressSpaceSize = 2048

type Computer struct {
	Opcodes [opcodeAddressSpaceSize]int
	InputPipe chan int
	OutputPipe chan int
}

func NewComputer(opcodes []int) *Computer {
	opArray := [opcodeAddressSpaceSize]int{}
	for i, val := range opcodes {
		opArray[i] = val
	}
	return &Computer{
		Opcodes:    opArray,
		InputPipe:  make(chan int),
		OutputPipe: make(chan int),
	}
}

func (c *Computer) Run() (int, error) {
	position := 0
	main:
		for {
			if position >= len(c.Opcodes) {
				return 0, NewPositionOutOfRangeError(fmt.Sprintf("position %d is out of range for opcodes with len %d", position, len(c.Opcodes)))
			}
			op := c.Opcodes[position]
			switch op {
			case 1:
				op1 := c.Opcodes[c.Opcodes[position + 1]]
				op2 := c.Opcodes[c.Opcodes[position + 2]]
				c.Opcodes[c.Opcodes[position + 3]] = op1 + op2
				position += 4

			case 2:
				op1 := c.Opcodes[c.Opcodes[position + 1]]
				op2 := c.Opcodes[c.Opcodes[position + 2]]
				c.Opcodes[c.Opcodes[position + 3]] = op1 * op2
				position += 4
			case 99:
				break main
			default:
				return 0, NewUnknownOpcodeError(fmt.Sprintf("unknown opcode received: %d", op))
			}
		}
	return c.Opcodes[0], nil
}