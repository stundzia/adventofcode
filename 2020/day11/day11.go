package day11

import (
	"github.com/stundzia/adventofcode/utils"
	"strconv"
)


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 11, "\n")
	ss := NewSeatingSpace(input)
	seatsTaken := ss.CountTaken()
	for i := 0; i < 250; i++ {
		ss.DoSeatingRound()
		newCount := ss.CountTaken()
		if newCount == seatsTaken && i > 5 {
			break
		} else {
			seatsTaken = newCount
		}
	}
	return strconv.Itoa(ss.CountTaken())
}


func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 11, "\n")
	ss := NewSeatingSpace(input)
	seatsTaken := ss.CountTaken()
	for i := 0; i < 250; i++ {
		ss.DoSeatingRoundByVisibility()
		newCount := ss.CountTaken()
		if newCount == seatsTaken && i > 5 {
			break
		} else {
			seatsTaken = newCount
		}
	}

	return strconv.Itoa(ss.CountTaken())
}
