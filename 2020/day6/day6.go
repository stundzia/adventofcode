package day6

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strconv"
)

func countGroup(group string) int {
	letters := map[rune]interface{}{}
	for _, l := range group {
		letters[l] = struct{}{}
	}
	return len(letters)
}

func countGroupV2(group []string) int {
	sum := 0
	letters := map[rune]int{}
	for _, person := range group {
		for _, l := range person {
			letters[l]++
		}
	}
	fmt.Println(letters)
	fmt.Println(len(group))
	for _, c := range letters {
		if c == len(group)-1 {
			sum++
		}
	}
	return sum
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 6, "\n")
	groups := []string{}
	group := ""
	for _, line := range input {
		if line == "" {
			groups = append(groups, fmt.Sprintf("%s", group))
			group = ""
			continue
		}
		group = fmt.Sprintf("%s%s", group, line)
	}
	sum := 0
	for _, g := range groups {
		sum += countGroup(g)
	}
	// 6358 is wrong
	return strconv.Itoa(sum)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 6, "\n")
	groups := [][]string{}
	group := []string{}
	for _, person := range input {
		group = append(group, person)
		if person == "" {
			groups = append(groups, group)
			group = []string{}
			continue
		}
	}
	sum := 0
	for _, g := range groups {
		sum += countGroupV2(g)
	}
	// 6358 is wrong
	return strconv.Itoa(sum)
}
