package day1

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)


func DoSilver() string {
	freqCahnges, _ := utils.ReadInputFileContentsAsIntSlice(2018, 1, "\n")
	resFrequency := 0
	for _, fc := range freqCahnges {
		resFrequency += fc
	}
	return fmt.Sprintf("Solution: %d", resFrequency)
}

func DoGold() string {
	frequencies := map[int]struct{}{0: {}}
	freqCahnges, _ := utils.ReadInputFileContentsAsIntSlice(2018, 1, "\n")
	resFrequency := 0
	infiniteLoop:
		for ;; {
			for _, fc := range freqCahnges {
				resFrequency += fc
				if _, ok := frequencies[resFrequency]; !ok {
					frequencies[resFrequency] = struct{}{}
				} else {
					break infiniteLoop
				}
			}
		}
	return fmt.Sprintf("Solution: %d", resFrequency)
}
