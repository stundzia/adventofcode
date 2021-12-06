package day6

import (
	"fmt"
	"time"

	"github.com/stundzia/adventofcode/utils"
)

type fishies struct {
	fish map[int]int
	lastFishID int
	nextDayLastFishID int
}

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
	eights := fs.states[8]
	sevens := fs.states[7]
	sixes := fs.states[6]
	fives := fs.states[5]
	fours := fs.states[4]
	threes := fs.states[3]
	twos := fs.states[2]
	ones := fs.states[1]
	zeros := fs.states[0]

	fs.states[6] = zeros + sevens
	fs.states[8] = zeros
	fs.states[7] = eights
	fs.states[5] = sixes
	fs.states[4] = fives
	fs.states[3] = fours
	fs.states[2] = threes
	fs.states[1] = twos
	fs.states[0] = ones
}

func (f *fishies) addFish() {
	f.nextDayLastFishID++
	id := f.nextDayLastFishID
	f.fish[id] = 8
}

func (f *fishies) showFish() {
	for i := 0; i <= f.nextDayLastFishID; i++ {
		fmt.Printf("%d|", f.fish[i])
	}
	fmt.Println("")
}

func (f *fishies) passDay() {
	for i := 0; i <= f.lastFishID; i++ {
		if f.fish[i] == 0 {
			f.fish[i] = 6
			f.addFish()
		} else {
			f.fish[i]--
		}
	}
	f.lastFishID = f.nextDayLastFishID
}

func newFishies(state []int) *fishies {
	f := &fishies{
		fish:              map[int]int{},
		lastFishID:        0,
		nextDayLastFishID: 0,
	}
	for id, fish := range state {
		f.fish[id] = fish
		f.lastFishID = id
		f.nextDayLastFishID = id
	}
	return f
}


func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2021, 6, ",")
	f := newFishies(nums)
	f.showFish()
	fs := newFishieStates(nums)
	for i := 0; i < 80; i++ {
		f.passDay()
		fs.passDay()
	}
	fmt.Println("fs: ", fs.getTotalCount())
	return fmt.Sprintf("Solution: %d", f.nextDayLastFishID + 1)
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsIntSlice(2021, 6, ",")
	fs := newFishieStates(nums)
	for i := 0; i < 256; i++ {
		start := time.Now()
		fs.passDay()
		fmt.Println("days done: ", i, " in ", time.Now().Sub(start))
	}
	return fmt.Sprintf("Solution: %d", fs.getTotalCount())
}
