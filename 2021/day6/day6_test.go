package day6

import (
	"testing"
)

func TestFishieStates(t *testing.T) {
	testSlice := []int{3,4,3,1,2}
	fs := newFishieStates(testSlice)
	if fs.getTotalCount() != len(testSlice) {
		t.Errorf("did not init with proper fish count, expected %d, but got %d", len(testSlice), fs.getTotalCount())
	}

	for i := 0; i < 80; i++ {
		fs.passDay()
	}

	if fs.getTotalCount() != 5934 {
		t.Errorf("expected total count after 80 days to be 5934, but got %d", fs.getTotalCount())
	}

	for i := 0; i < 256 - 80; i++ {
		fs.passDay()
	}
	if fs.getTotalCount() != 26984457539 {
		t.Errorf("expected total count after 80 days to be 26984457539, but got %d", fs.getTotalCount())
	}
}
