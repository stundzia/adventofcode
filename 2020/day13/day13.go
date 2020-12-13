package day13

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strconv"
)


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 13, "\n")
	departureTime, _ := strconv.Atoi(input[0])
	bs := NewBusSchedule(input[1])
	for _, bus := range bs.Buses {
		bus.GenerateSchedule(departureTime, departureTime + 1000)
	}
	res := 0
	main:
	for i := departureTime; i < departureTime + 1000; i++ {
		for _, b := range bs.Buses {
			if b.LeavesAtTimestamp(i) {
				res = (i - departureTime) * b.ID
				break main
			}
		}
	}
	return strconv.Itoa(res)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 13, "\n")
	bs := NewBusSchedule(input[1])

	return fmt.Sprintf("%d", bs.GetFirstSequentialTimestamp())
}
