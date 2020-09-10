package computer

import "log"

type Computer struct {
	Opcodes []int
	InputPipe chan int
	OutputPipe chan int
}

func NewComputer(opcodes []int) *Computer {
	return &Computer{
		Opcodes:    opcodes,
		InputPipe:  make(chan int),
		OutputPipe: make(chan int),
	}
}

func (c *Computer) Run() int {
	position := 0
	main:
		for {
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
				log.Fatalf("unknown opcode received: %d", op)
			}
		}
	return c.Opcodes[0]
}