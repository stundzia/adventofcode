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

func getOperationAndParameterModes(opcode int) (op, param1, param2, param3 int) {
	op = opcode % 100
	param1 = opcode % 1000 / 100
	param2 = opcode % 10000 / 1000
	param3 = opcode % 100000 / 10000
	return
}

func (c *Computer) Run() (int, error) {
	position := 0
	main:
		for {
			if position >= len(c.Opcodes) {
				return 0, NewPositionOutOfRangeError(fmt.Sprintf("position %d is out of range for opcodes with len %d", position, len(c.Opcodes)))
			}

			opCode, modeParam1, modeParam2, _ := getOperationAndParameterModes(c.Opcodes[position])
			switch opCode {

			case 1:
				var param1 int
				if modeParam1 == 0 {
					param1 = c.Opcodes[c.Opcodes[position + 1]]
				} else {
					param1 = c.Opcodes[position + 1]
				}
				var param2 int
				if modeParam2 == 0 {
					param2 = c.Opcodes[c.Opcodes[position + 2]]
				} else {
					param2 = c.Opcodes[position + 2]
				}
				address := c.Opcodes[position + 3]
				c.Opcodes[address] = param1 + param2
				position += 4

			case 2:
				var param1 int
				if modeParam1 == 0 {
					param1 = c.Opcodes[c.Opcodes[position + 1]]
				} else {
					param1 = c.Opcodes[position + 1]
				}
				var param2 int
				if modeParam2 == 0 {
					param2 = c.Opcodes[c.Opcodes[position + 2]]
				} else {
					param2 = c.Opcodes[position + 2]
				}
				address := c.Opcodes[position + 3]
				c.Opcodes[address] = param1 * param2
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