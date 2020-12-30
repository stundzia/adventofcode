package computer

import "testing"

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