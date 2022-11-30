package computer

import (
	"fmt"
	"sync/atomic"
)

const opcodeAddressSpaceSize = 2048

type Computer struct {
	Opcodes     [opcodeAddressSpaceSize]int
	FirstInputs []int
	InputPipe   chan int
	OutputPipe  chan int
	Running     atomic.Bool
}

func NewComputer(opcodes []int) *Computer {
	opArray := [opcodeAddressSpaceSize]int{}
	for i, val := range opcodes {
		opArray[i] = val
	}
	return &Computer{
		Opcodes:     opArray,
		FirstInputs: []int{},
		InputPipe:   make(chan int, 2),
		OutputPipe:  make(chan int, 2),
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
	c.Running.Store(true)

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
				param1 = c.Opcodes[c.Opcodes[position+1]]
			} else {
				param1 = c.Opcodes[position+1]
			}
			var param2 int
			if modeParam2 == 0 {
				param2 = c.Opcodes[c.Opcodes[position+2]]
			} else {
				param2 = c.Opcodes[position+2]
			}
			address := c.Opcodes[position+3]
			c.Opcodes[address] = param1 + param2
			position += 4
			break

		case 2:
			var param1 int
			if modeParam1 == 0 {
				param1 = c.Opcodes[c.Opcodes[position+1]]
			} else {
				param1 = c.Opcodes[position+1]
			}
			var param2 int
			if modeParam2 == 0 {
				param2 = c.Opcodes[c.Opcodes[position+2]]
			} else {
				param2 = c.Opcodes[position+2]
			}
			address := c.Opcodes[position+3]
			c.Opcodes[address] = param1 * param2
			position += 4
			break

		case 3:
			address := c.Opcodes[position+1]
			var v int
			if len(c.FirstInputs) > 0 {
				v = c.FirstInputs[0]
				c.FirstInputs = c.FirstInputs[1:]
			} else {
				v = <-c.InputPipe
			}
			c.Opcodes[address] = v
			position += 2
			break

		case 4:
			address := c.Opcodes[position+1]
			c.OutputPipe <- c.Opcodes[address]
			position += 2
			break

		case 5:
			var param1 int
			if modeParam1 == 0 {
				param1 = c.Opcodes[c.Opcodes[position+1]]
			} else {
				param1 = c.Opcodes[position+1]
			}
			var param2 int
			if modeParam2 == 0 {
				param2 = c.Opcodes[c.Opcodes[position+2]]
			} else {
				param2 = c.Opcodes[position+2]
			}
			if param1 > 0 {
				position = param2
			} else {
				position += 3
			}
			break

		case 6:
			var param1 int
			if modeParam1 == 0 {
				param1 = c.Opcodes[c.Opcodes[position+1]]
			} else {
				param1 = c.Opcodes[position+1]
			}
			var param2 int
			if modeParam2 == 0 {
				param2 = c.Opcodes[c.Opcodes[position+2]]
			} else {
				param2 = c.Opcodes[position+2]
			}
			if param1 == 0 {
				position = param2
			} else {
				position += 3
			}
			break

		case 7:
			var param1 int
			if modeParam1 == 0 {
				param1 = c.Opcodes[c.Opcodes[position+1]]
			} else {
				param1 = c.Opcodes[position+1]
			}
			var param2 int
			if modeParam2 == 0 {
				param2 = c.Opcodes[c.Opcodes[position+2]]
			} else {
				param2 = c.Opcodes[position+2]
			}
			address := c.Opcodes[position+3]
			if param1 < param2 {
				c.Opcodes[address] = 1
			} else {
				c.Opcodes[address] = 0
			}
			position += 4
			break

		case 8:
			var param1 int
			if modeParam1 == 0 {
				param1 = c.Opcodes[c.Opcodes[position+1]]
			} else {
				param1 = c.Opcodes[position+1]
			}
			var param2 int
			if modeParam2 == 0 {
				param2 = c.Opcodes[c.Opcodes[position+2]]
			} else {
				param2 = c.Opcodes[position+2]
			}
			address := c.Opcodes[position+3]
			if param1 == param2 {
				c.Opcodes[address] = 1
			} else {
				c.Opcodes[address] = 0
			}
			position += 4
			break

		case 99:
			break main

		default:
			return 0, NewUnknownOpcodeError(fmt.Sprintf("unknown opcode received: %d", opCode))
		}
	}

	c.Running.Store(false)
	return c.Opcodes[0], nil
}
