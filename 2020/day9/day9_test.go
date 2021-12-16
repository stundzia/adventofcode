package day9

import (
	"github.com/stundzia/adventofcode/utils"
	"testing"
)

var testInput = []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

func TestIsValid(t *testing.T) {
	for n, num := range testInput {
		valid := isValid(testInput, n, 5)
		if valid && num == 127 {
			t.Errorf("num %d should be invalid", num)
		}
		if !valid && num != 127 {
			t.Errorf("num %d should be valid", num)
		}
	}
}

func TestContiguousSet(t *testing.T) {
	invalid := 127
	expected := []int{15, 25, 47, 40}
	var found bool
	var set []int
	for n, _ := range testInput {
		found, set = contiguousSet(testInput, n, invalid)
		if found && !utils.SlicesIntEqual(expected, set) {
			t.Errorf("expected found set to be equal to %v, but was %v", expected, set)
		}
		if found {
			break
		}
	}
	if !found {
		t.Errorf("unable to find contiguous set")
	}
}
