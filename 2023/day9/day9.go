package day9

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
)

type sequence struct {
	og []int
}

func parseSequence(line []int) sequence {
	s := sequence{og: line}

	return s
}

func (s sequence) diffs() sequence {
	diffs := []int{}
	for i := 0; i < len(s.og)-1; i++ {
		diffs = append(diffs, s.og[i+1]-s.og[i])
	}

	return parseSequence(diffs)
}

func (s sequence) allZero() bool {
	for _, n := range s.og {
		if n != 0 {
			return false
		}
	}
	return true
}

func (s sequence) nextVal() int {
	if s.allZero() {
		return 0
	}
	return s.og[len(s.og)-1] + s.diffs().nextVal()
}

func (s sequence) previousVal() int {
	if s.allZero() {
		return 0
	}
	return s.og[0] - s.diffs().previousVal()
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsIntSlices(2023, 9, " ")
	res := 0
	for _, l := range input {
		res += parseSequence(l).nextVal()
	}

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsIntSlices(2023, 9, " ")
	res := 0
	for _, l := range input {
		res += parseSequence(l).previousVal()
	}

	return fmt.Sprintf("Solution: %d", res)
}
