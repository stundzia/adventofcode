package day13

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strconv"
	"strings"
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

	start, diff := bs.GetDiff(0,5, 1)
	for i := 5; i < len(strings.Split(input[1], ",")); i+=2 {
		start, diff = bs.GetDiff(start, i, diff)
	}
	return fmt.Sprintf("%d", start)
}
