package day6

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type race struct {
	time      int
	distance  int
	winCombos atomic.Int64
}

func (r *race) calcWinCombos() {
	for i := 0; i < r.time; i++ {
		if (r.time-i)*i > r.distance {
			r.winCombos.Add(1)
		}
	}
}

func (r *race) calcWinCombosAsync() {
	wg := &sync.WaitGroup{}
	ranges := splitIntoRangesInclusive(r.time, runtime.NumCPU())
	for _, rg := range ranges {
		wg.Add(1)
		go r.calcComboRange(wg, rg[0], rg[1])
	}
	wg.Wait()
}

func (r *race) calcComboRange(wg *sync.WaitGroup, iMin, iMax int) {
	defer wg.Done()
	for i := iMin; i <= iMax; i++ {
		if (r.time-i)*i > r.distance {
			r.winCombos.Add(1)
		}
	}
}

func splitIntoRangesInclusive(i, count int) [][2]int {
	res := [][2]int{}
	rangeSize := i / count
	for it := 0; it < i; it += rangeSize {
		end := it + rangeSize
		if end >= i {
			end = i + 1
		}
		res = append(res, [2]int{it, end - 1})
	}

	return res
}

func DoSilver() string {
	//Time:        41     77     70     96
	//Distance:   249   1362   1127   1011
	r1 := &race{
		time:      41,
		distance:  249,
		winCombos: atomic.Int64{},
	}
	r1.calcWinCombos()
	r2 := &race{
		time:      77,
		distance:  1362,
		winCombos: atomic.Int64{},
	}
	r2.calcWinCombos()
	r3 := &race{
		time:      70,
		distance:  1127,
		winCombos: atomic.Int64{},
	}
	r3.calcWinCombos()
	r4 := &race{
		time:      96,
		distance:  1011,
		winCombos: atomic.Int64{},
	}
	r4.calcWinCombos()

	res := r1.winCombos.Load() * r2.winCombos.Load() * r3.winCombos.Load() * r4.winCombos.Load()

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	r1 := &race{
		time:      41777096,
		distance:  249136211271011,
		winCombos: atomic.Int64{},
	}
	r1.calcWinCombosAsync()

	res := r1.winCombos.Load()

	return fmt.Sprintf("Solution: %d", res)
}
