package computer

import (
	"github.com/stundzia/adventofcode/utils"
	"testing"
)

func TestComputerBasicRun(t *testing.T) {
	opcodes := []int{
		1,9,10,3,
		2,3,11,0,
		99,
		30,40,50,
	}
	comp := NewComputer(opcodes)
	res, _ := comp.Run()
	if res != 3500 {
		t.Errorf("incorrect opcode at position 0, expected 3500, but found %d", res)
	}
}

func TestComputerWithImmediateMode(t *testing.T) {
	opcodes := []int{
		1,9,10,3,
		2,3,11,0,
		1102,23,42,0,
		99,
		30,40,50,
	}
	comp := NewComputer(opcodes)
	res, _ := comp.Run()
	if res != 23 * 42 {
		t.Errorf("incorrect opcode at position 0, expected 966, but found %d", res)
	}
}

func TestComputerWithInputOutput(t *testing.T) {
	opcodes := []int{
		1,19,20,3,
		2,3,21,0,
		3,128,
		4,128,
		3,256,
		4,256,
		4,0,
		99,
		30,40,50,
	}
	comp := NewComputer(opcodes)
	comp.InputPipe <- 25
	go comp.Run()
	output := <- comp.OutputPipe

	if output != 25 {
		t.Errorf("incorrect output, expected 25, but got %d", output)
	}

	comp.InputPipe <- 42
	output = <- comp.OutputPipe

	if output != 42 {
		t.Errorf("incorrect output, expected 42, but got %d", output)
	}

	res := <- comp.OutputPipe
	if res != 3500 {
		t.Errorf("incorrect result (address 0), expected 3500, but got %d", res)
	}
}

func TestComputerAdvanced(t *testing.T) {
	opcodes := []int{
		3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
		1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
		999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99,
	}
	comp := NewComputer(opcodes)
	go comp.Run()
	comp.InputPipe <- 9
	output := <- comp.OutputPipe

	if output != 1001 {
		t.Errorf("incorrect output, expected 1001, but got %d", output)
	}
}


func TestGetOperationAndParameterModes(t *testing.T) {
	tcs := []struct{
		test string
		opcode int
		expected []int
	}{
		{
			"opcode: 1002, should be: 2, 0, 1, 0",
			1002,
			[]int{2,0,1,0},
		},
		{
			"opcode: 11004, should be: 4, 0, 1, 1",
			11004,
			[]int{4,0,1,1},
		},
		{
			"opcode: 4, should be: 4, 0, 0, 0",
			4,
			[]int{4,0,0,0},
		},
		{
			"opcode: 10101, should be: 1, 1, 0, 1",
			10101,
			[]int{1,1,0,1},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			op, param1, param2, param3 := getOperationAndParameterModes(tc.opcode)
			if !utils.SlicesIntEqual([]int{op, param1, param2, param3}, tc.expected) {
				t.Errorf("expected %v, but got %v", tc.expected, []int{op, param1, param2, param3})
			}
		})
	}
}