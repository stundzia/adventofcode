package day15

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)


func DoSilver() string {
	numberTurnMap := map[int][2]int{}
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2020, 15, ",")
	turn := 0
	lastNum := 0
	var res int
	for _, num := range nums {
		turn++
		numberTurnMap[num] = [2]int{turn, -1}
		lastNum = num
	}
	for ;; {
		turn++
		num := 0
		if turns, ok := numberTurnMap[lastNum]; ok {
			if turns[1] != -1 {
				num = turns[1] - turns[0]
			}
		}
		if turns, ok := numberTurnMap[num]; ok {
			if turns[1] == -1 {
				numberTurnMap[num] = [2]int{turns[0], turn}
			} else {
				numberTurnMap[num] = [2]int{turns[1], turn}
			}
		} else {
			numberTurnMap[num] = [2]int{turn, -1}
		}
		lastNum = num
		if turn == 2020 {
			res = num
			break
		}
	}
	return fmt.Sprintf("%d", res)
}

func DoGold() string {
	numberTurnMap := map[int][2]int{}
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2020, 15, ",")
	turn := 0
	lastNum := 0
	var res int
	for _, num := range nums {
		turn++
		numberTurnMap[num] = [2]int{turn, -1}
		lastNum = num
	}
	for ;; {
		turn++
		num := 0
		if turns, ok := numberTurnMap[lastNum]; ok {
			if turns[1] != -1 {
				num = turns[1] - turns[0]
			}
		}
		if turns, ok := numberTurnMap[num]; ok {
			if turns[1] == -1 {
				numberTurnMap[num] = [2]int{turns[0], turn}
			} else {
				numberTurnMap[num] = [2]int{turns[1], turn}
			}
		} else {
			numberTurnMap[num] = [2]int{turn, -1}
		}
		lastNum = num
		if turn == 30000000 {
			res = num
			break
		}
	}
	return fmt.Sprintf("%d", res)
}
