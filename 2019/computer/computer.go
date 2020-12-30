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
		InputPipe:  make(chan int, 1),
		OutputPipe: make(chan int, 1),
	}
}

func (c *Computer) Run() (int, error) {
	position := 0
	main:
		for {
			if position >= len(c.Opcodes) {
				return 0, NewPositionOutOfRangeError(fmt.Sprintf("position %d is out of range for opcodes with len %d", position, len(c.Opcodes)))
			}

			opCode := c.Opcodes[position]
			switch opCode {

			case 1:
				op1 := c.Opcodes[c.Opcodes[position + 1]]
				op2 := c.Opcodes[c.Opcodes[position + 2]]
				address := c.Opcodes[position + 3]
				c.Opcodes[address] = op1 + op2
				position += 4

			case 2:
				op1 := c.Opcodes[c.Opcodes[position + 1]]
				op2 := c.Opcodes[c.Opcodes[position + 2]]
				address := c.Opcodes[position + 3]
				c.Opcodes[address] = op1 * op2
				position += 4

			case 3:
				address := c.Opcodes[position + 1]
				c.Opcodes[address] = <- c.InputPipe
				position += 2

			case 4:
				address := c.Opcodes[position + 1]
				c.OutputPipe <- c.Opcodes[address]
				position += 2

			case 99:
				break main

			default:
				return 0, NewUnknownOpcodeError(fmt.Sprintf("unknown opcode received: %d", opCode))
			}
		}
	return c.Opcodes[0], nil
}