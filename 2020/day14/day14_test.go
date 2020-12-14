package day14

import (
	"fmt"
	"testing"
)

func TestMemoryMaskStringToMasks(t *testing.T) {
	c := NewComputer(2)
	c.MemoryMask = "01X0X1X"
	masks := c.memoryMaskStringToMasks()
	fmt.Println(masks)
	if len(masks) != 8 {
		t.Errorf("invalid number of masks")
	}
}
