package day17

import (
	"testing"
)

func TestPockedDimension(t *testing.T) {
	initial := []string{
		".#.",
		"..#",
		"###",
	}
	tcs := []struct {
		test       string
		input      []string
		cycles     int
		expected   int
		dimensions int
	}{
		{
			"Slice: 3 cycles, 3 dimensions",
			initial,
			3,
			38,
			3,
		},
		{
			"Slice: 3 cycles, 3 dimensions",
			initial,
			3,
			320,
			4,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			pd := NewPockedDimensionFromInitialStateSlice(tc.input, tc.dimensions)
			for i := 0; i < tc.cycles; i++ {
				if tc.dimensions == 4 {
					pd.Do4DCycle()
				} else {
					pd.DoCycle()
				}
			}
			var res int
			if tc.dimensions == 4 {
				res = pd.Get4DActiveCount()
			} else {
				res = pd.GetActiveCount()
			}
			if res != tc.expected {
				t.Errorf("expected active count to be %d, but got %d", tc.expected, res)
			}
		})
	}
}
