package day5

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

func parseLine(l string) (x1, x2, y1, y2 int) {
	parts := strings.Split(l, " -> ")
	beginCoords := strings.Split(parts[0], ",")
	endCoords := strings.Split(parts[1], ",")
	x1, _ = strconv.Atoi(beginCoords[0])
	y1, _ = strconv.Atoi(beginCoords[1])
	x2, _ = strconv.Atoi(endCoords[0])
	y2, _ = strconv.Atoi(endCoords[1])
	return x1, x2, y1, y2
}

func part1(lines []string) int {
	coordsMap := map[string]int{}
	for _, line := range lines {
		x1, x2, y1, y2 := parseLine(line)
		if x1 == x2 {
			if y2 > y1 {
				for i := y1; i <= y2; i++ {
					key := fmt.Sprintf("%d_%d", x1, i)
					if _, ok := coordsMap[key]; !ok {
						coordsMap[key] = 0
					}
					coordsMap[key]++
				}
			} else {
				for i := y1; i >= y2; i-- {
					key := fmt.Sprintf("%d_%d", x1, i)
					if _, ok := coordsMap[key]; !ok {
						coordsMap[key] = 0
					}
					coordsMap[key]++
				}
			}
		}
		if y1 == y2 {
			if x2 > x1 {
				for i := x1; i <= x2; i++ {
					key := fmt.Sprintf("%d_%d", i, y1)
					if _, ok := coordsMap[key]; !ok {
						coordsMap[key] = 0
					}
					coordsMap[key]++
				}
			} else {
				for i := x1; i >= x2; i-- {
					key := fmt.Sprintf("%d_%d", i, y1)
					if _, ok := coordsMap[key]; !ok {
						coordsMap[key] = 0
					}
					coordsMap[key]++
				}
			}
		}
	}
	count := 0
	for _, v := range coordsMap {
		if v >= 2 {
			count++
		}
	}
	return count
}

func part2(lines []string) int {
	coordsMap := map[string]int{}
	for _, line := range lines {
		x1, x2, y1, y2 := parseLine(line)
		if x1 == x2 {
			if y2 > y1 {
				for i := y1; i <= y2; i++ {
					key := fmt.Sprintf("%d_%d", x1, i)
					if _, ok := coordsMap[key]; !ok {
						coordsMap[key] = 0
					}
					coordsMap[key]++
				}
			} else {
				for i := y1; i >= y2; i-- {
					key := fmt.Sprintf("%d_%d", x1, i)
					if _, ok := coordsMap[key]; !ok {
						coordsMap[key] = 0
					}
					coordsMap[key]++
				}
			}
		}
		if y1 == y2 {
			if x2 > x1 {
				for i := x1; i <= x2; i++ {
					key := fmt.Sprintf("%d_%d", i, y1)
					if _, ok := coordsMap[key]; !ok {
						coordsMap[key] = 0
					}
					coordsMap[key]++
				}
			} else {
				for i := x1; i >= x2; i-- {
					key := fmt.Sprintf("%d_%d", i, y1)
					if _, ok := coordsMap[key]; !ok {
						coordsMap[key] = 0
					}
					coordsMap[key]++
				}
			}
		}

		if x1 != x2 && y1 != y2 {
			xs := getRangeSlice(x1, x2)
			ys := getRangeSlice(y1, y2)
			if len(xs) != len(ys) {
				log.Fatal("SOMETHING FUCKY!!!")
			}
			for i, x := range xs {
				key := fmt.Sprintf("%d_%d", x, ys[i])
				if _, ok := coordsMap[key]; !ok {
					coordsMap[key] = 0
				}
				coordsMap[key]++
			}
		}
	}
	count := 0
	for _, v := range coordsMap {
		if v >= 2 {
			count++
		}
	}
	return count
}

func getRangeSlice(begin, end int) []int {
	res := []int{}
	if end > begin {
		for i := begin; i <= end; i++ {
			res = append(res, i)
		}
	} else {
		for i := begin; i >= end; i-- {
			res = append(res, i)
		}
	}
	return res
}

func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 5, "\n")
	return fmt.Sprintf("Solution: %d", part1(lines))
}

func DoGold() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 5, "\n")
	return fmt.Sprintf("Solution: %d", part2(lines))
}
