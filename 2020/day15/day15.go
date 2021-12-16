package day15

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)

func PlayTheGame(startingNums []int, turnToGet int) int {
	numberTurnMap := map[int][2]int{}
	turn := 0
	lastNum := 0
	for _, num := range startingNums {
		turn++
		numberTurnMap[num] = [2]int{turn, -1}
		lastNum = num
	}
	for {
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
		if turn == turnToGet {
			return num
		}
	}
}

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2020, 15, ",")
	res := PlayTheGame(nums, 2020)
	return fmt.Sprintf("%d", res)
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2020, 15, ",")
	res := PlayTheGame(nums, 30000000)
	return fmt.Sprintf("%d", res)
}
