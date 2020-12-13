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
	fmt.Println("len: ", len(strings.Split(input[1], ",")))
	for sID, b := range bs.Buses {
		fmt.Println(sID, " : ", b.ID)
	}

	var res uint64
	var t uint64
	start, diff := bs.GetDiff(379523,75, 438857)
	//start, diff = bs.GetDiff(start, 4, diff)
	fmt.Println(":ttt: ", start, diff)
	for t = start; t < 613416286793297; t+=diff {
		if bs.IsSequentialDepartureBig(t) {
			fmt.Println(t)
			fmt.Println("Diff: ", t - res)
			res = t
			break
		}
		if t % 100000000 == 0 {
			fmt.Println("Done: ", t)
			fmt.Println("Assumed %: ", float64(t)/float64(613416286793297) * 100, "%")
		}
	}
	//bs.FindFirstSequentialDeparture()
	return fmt.Sprintf("%d", res)
}
