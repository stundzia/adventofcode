package day1

import (
	"testing"

	"github.com/stundzia/adventofcode/utils"
)

func TestSomething(t *testing.T) {
	testSlice := []int{1721, 979, 366, 299, 675, 1456}
	res := utils.SumIntSlice(testSlice)
	expected := 5496
	if res != expected {
		t.Errorf("Expected %d, but got %d", expected, res)
	}
}
