package day3

import (
	"fmt"

	"github.com/stundzia/adventofcode/utils"
)

func getPrio(c rune) int {
	if c > 96 {
		return int(c) - 96
	}
	return int(c) - 38
}

func parseRucksack(l string) (string, string) {
	compSize := len(l) / 2
	c1 := l[:compSize]
	c2 := l[compSize:]
	return c1, c2
}

func part1(comp1, comp2 string) int {
	c1map := map[rune]struct{}{}
	for _, c := range comp1 {
		c1map[c] = struct{}{}
	}
	for _, c := range comp2 {
		if _, found := c1map[c]; found {
			fmt.Println(string(c))
			fmt.Println(getPrio(c))
			return getPrio(c)
		}
	}
	return 0
}

func rucksackToMap(l string) map[rune]struct{} {
	res := map[rune]struct{}{}
	for _, c := range l {
		res[c] = struct{}{}
	}
	return res
}

func part2(group []string) int {
	b1map := rucksackToMap(group[0])
	b2map := rucksackToMap(group[1])
	b3map := rucksackToMap(group[2])
	for c, _ := range b1map {
		if _, f := b2map[c]; f {
			if _, f := b3map[c]; f {
				return getPrio(c)
			}
		}
	}
	return 0
}

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2022, 3, "\n")
	sum := 0
	for _, n := range nums {
		sum += part1(parseRucksack(n))
	}

	return fmt.Sprintf("Solution: %d", sum)
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2022, 3, "\n")

	group := []string{}
	sum := 0
	for _, n := range nums {
		fmt.Println(len(group))

		group = append(group, n)
		if len(group) == 3 {
			sum += part2(group)
			group = []string{}
		}
	}

	return fmt.Sprintf("Solution: %d", sum)
}
