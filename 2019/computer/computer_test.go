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