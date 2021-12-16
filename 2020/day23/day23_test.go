package day23

import (
	"github.com/stundzia/adventofcode/utils"
	"testing"
)

func TestSample(t *testing.T) {
	tcs := []struct {
		test     string
		slice    []int
		expected int
	}{
		{
			"Slice: 1,3,2 expected sum: 6",
			[]int{1, 3, 2},
			6,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			sum := utils.SumIntSlice(tc.slice)
			if sum != tc.expected {
				t.Errorf("expected %d, but got %d", tc.expected, sum)
			}
		})
	}
}
