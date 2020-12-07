package day5

import (
	"fmt"
	"testing"
)

func TestParseSeat(t *testing.T) {
	cases := map[string][]int{
		"BFFFBBFRRR": {70, 7, 567}, // 0b1000110
		"FFFBBBFRRR": {14, 7, 119},
		"BBFFBBFRLL": {102, 4, 820},
		"FBFBBFFRLR": {44, 5, 357},
	}
	for tc, res := range cases {
		row, column, seatID := parseSeat(tc)
		fmt.Printf("%d %d %d", row, column, seatID)
		if row != res[0] || column != res[1] {
			t.Errorf("Expected %s to be %v but got: %d %d %d", tc, res, row, column, seatID)
		}
	}
}
