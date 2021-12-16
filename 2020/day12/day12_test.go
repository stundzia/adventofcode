package day12

import "testing"

func TestRotateCoords(t *testing.T) {
	tcs := []struct {
		test     string
		coords   [2]int
		degrees  int
		expected [2]int
	}{
		{
			"10,1 by 90 degrees",
			[2]int{10, 1},
			90,
			[2]int{1, -10},
		},
		{
			"-4,5 by 180 degrees",
			[2]int{-4, 5},
			180,
			[2]int{4, -5},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			if coords := rotateCoords(tc.coords, tc.degrees); coords != tc.expected {
				t.Errorf("expected new coords to be %v, but got %v", tc.expected, coords)
			}
		})
	}
}
