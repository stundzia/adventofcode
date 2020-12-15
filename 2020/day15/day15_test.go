package day15

import "testing"

func TestPlayTheGame(t *testing.T) {
	tcs := []struct{
		test string
		startingNums []int
		turnToGet int
		expected int
	}{
		{
			"Starting nums: 1,3,2 turn to return: 2020",
			[]int{1,3,2},
			2020,
			1,
		},
		{
			"Starting nums: 2,1,3 turn to return: 2020",
			[]int{2,1,3},
			2020,
			10,
		},
		{
			"Starting nums: 3,1,2 turn to return: 2020",
			[]int{3,1,2},
			2020,
			1836,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			res := PlayTheGame(tc.startingNums, tc.turnToGet)
			if res != tc.expected {
				t.Errorf("expected turn %d to be %d, but got %d", tc.turnToGet, tc.expected, res)
			}
		})
	}
}

