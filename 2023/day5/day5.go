package day5

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/sync/semaphore"

	"github.com/stundzia/adventofcode/utils"
)

// FIXME: don't use global vars, that's why p2 took so long, idiot!
var path = []string{}
var maps = map[string]sourceDestMap{}

type sourceDestMap struct {
	name   string
	ranges []mapRange
}

type mapRange struct {
	startSource int
	startDest   int
	length      int
}

func (sdm sourceDestMap) getVal(key int) int {
	for _, rng := range sdm.ranges {
		if v, found := rng.getVal(key); found {
			return v
		}
	}
	return key
}

func (mr mapRange) getVal(key int) (int, bool) {
	if key >= mr.startSource && key <= mr.startSource+mr.length {
		return mr.startDest + (key - mr.startSource), true
	}
	return key, false
}

func parseMapLines(lines []string) {
	name := strings.Replace(lines[0], " map:", "", -1)
	path = append(path, name)
	sdm := sourceDestMap{
		name:   name,
		ranges: []mapRange{},
	}
	for i := 1; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		if len(parts) != 3 {
			log.Fatal("wtf: ", parts)
		}
		destRangeStart, _ := strconv.Atoi(parts[0])
		sourceRangeStart, _ := strconv.Atoi(parts[1])
		rangeLength, _ := strconv.Atoi(parts[2])
		sdm.ranges = append(sdm.ranges, mapRange{
			startSource: sourceRangeStart,
			startDest:   destRangeStart,
			length:      rangeLength,
		})
	}
	maps[name] = sdm

}

func getSeedsToLocationMap(seeds string) map[int]int {
	seedsMap := map[int]int{}
	parts := strings.Split(seeds, ": ")
	seedsStr := strings.Split(parts[1], " ")
	for _, seed := range seedsStr {
		seedInt, _ := strconv.Atoi(seed)
		seedsMap[seedInt] = -1
		key := seedInt
		for _, v := range path {
			key = maps[v].getVal(key)
		}
		seedsMap[seedInt] = key
	}
	return seedsMap
}

func getSeedsToLocationMapPart2(seeds string) int {
	partsf := strings.Split(seeds, ": ")
	parts := strings.Split(partsf[1], " ")
	sem := semaphore.NewWeighted(600)
	wg := sync.WaitGroup{}
	mux := sync.Mutex{}
	lowestLoc := math.MaxInt
	for i := 0; i < len(parts); i += 2 {
		start, _ := strconv.Atoi(parts[i])
		length, _ := strconv.Atoi(parts[i+1])
		for ii := 0; ii < length; ii++ {
			_ = sem.Acquire(context.Background(), 1)
			wg.Add(1)
			go func(seed int) {
				key := seed
				for _, v := range path {
					key = maps[v].getVal(key)
				}
				mux.Lock()
				if key < lowestLoc {
					lowestLoc = key
				}
				mux.Unlock()
				sem.Release(1)
				wg.Done()
			}(start + ii)
		}
	}
	wg.Wait()

	return lowestLoc
}

func DoSilver() string {
	path = []string{}
	maps = map[string]sourceDestMap{}

	input, _ := utils.ReadInputFileContentsAsStringSliceLines(2023, 5)
	for _, l := range input[1:] {
		parseMapLines(l)
	}
	seedsMap := getSeedsToLocationMap(input[0][0])

	res := math.MaxInt
	for _, v := range seedsMap {
		if v < res {
			res = v
		}
	}
	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	path = []string{}
	maps = map[string]sourceDestMap{}

	input, _ := utils.ReadInputFileContentsAsStringSliceLines(2023, 5)
	for _, l := range input[1:] {
		parseMapLines(l)
	}
	res := getSeedsToLocationMapPart2(input[0][0])

	return fmt.Sprintf("Solution: %d", res)
}
