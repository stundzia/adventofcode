package day5

import (
	"github.com/stundzia/adventofcode/utils"
	"strconv"
	"strings"
)

func parseSeat(seat string) (row, column, seatID int) {
	rowBi := strings.Replace(seat[:7], "B", "1", -1)
	rowBi = strings.Replace(rowBi, "F", "0", -1)
	rowInt64, _ := strconv.ParseInt(rowBi, 2, 64)
	row = int(rowInt64)

	columnBi := strings.Replace(seat[7:], "R", "1", -1)
	columnBi = strings.Replace(columnBi, "L", "0", -1)
	columnInt64, _ := strconv.ParseInt(columnBi, 2, 64)
	column = int(columnInt64)

	seatID = row*8 + column
	return row, column, seatID
}

func DoSilver() string {
	maxSeatID := 0
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 5, "\n")
	for _, seat := range input {
		_, _, seatID := parseSeat(seat)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	return strconv.Itoa(maxSeatID)
}

func DoGold() string {
	maxSeatID, _ := strconv.Atoi(DoSilver())
	seats := map[int]interface{}{}
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 5, "\n")
	for _, seat := range input {
		_, _, seatID := parseSeat(seat)
		seats[seatID] = struct{}{}
	}
	for i := 40; i < maxSeatID; i++ {
		if _, ok := seats[i]; !ok {
			return strconv.Itoa(i)
		}
	}
	return "Did not find solution"
}
