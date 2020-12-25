package day25

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)

const SubjectDivider = 20201227
const SubjectNumber = 7


func getLoopSize(pKey int) int {
	loopSize := 0
	val := 1
	for ; val != pKey; {
		val *= SubjectNumber
		val %= SubjectDivider
		loopSize++
	}
	return loopSize
}

func transform(pKey, loopSize int) int {
	val := 1
	for i := 0; i < loopSize; i++ {
		val *= pKey
		val %= SubjectDivider
	}
	return val
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2020, 25, "\n")

	cardLoopSize := getLoopSize(input[0])
	doorLoopSize := getLoopSize(input[1])

	return fmt.Sprintf("%d (should be equal to %d)", transform(input[0], doorLoopSize), transform(input[1], cardLoopSize))
}

func DoGold() string {
	return "There is not part 2 (or part 2 is all the other days and all their parts). Merry Christmas!"
}
