package day6

import (
	"fmt"

	"github.com/stundzia/adventofcode/utils"
)

type fishieStates struct {
	states map[int]int

}

func newFishieStates(states []int) *fishieStates {
	fs := &fishieStates{map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}}
	for _, state := range states {
		fs.states[state]++
	}
	return fs
}

func (fs *fishieStates) getTotalCount() int {
	count := 0
	for _, v := range fs.states {
		count += v
	}
	return count
}

func (fs *fishieStates) passDay() {
	_0s, _1s, _2s, _3s, _4s, _5s, _6s, _7s, _8s := fs.states[0], fs.states[1], fs.states[2], fs.states[3], fs.states[4], fs.states[5], fs.states[6], fs.states[7], fs.states[8]

	fs.states[6] = _0s + _7s
	fs.states[8] = _0s
	fs.states[7] = _8s
	fs.states[5] = _6s
	fs.states[4] = _5s
	fs.states[3] = _4s
	fs.states[2] = _3s
	fs.states[1] = _2s
	fs.states[0] = _1s
}

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2021, 6, ",")
	fs := newFishieStates(nums)
	for i := 0; i < 80; i++ {
		fs.passDay()
	}
	return fmt.Sprintf("Solution: %d", fs.getTotalCount())
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2021, 6, ",")
	fs := newFishieStates(nums)
	for i := 0; i < 256; i++ {
		fs.passDay()
	}
	return fmt.Sprintf("Solution: %d", fs.getTotalCount())
}
