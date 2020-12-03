package day3

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strings"
)

func asdf(min int, max int, letter string, pass string) bool {
	count := strings.Count(pass, letter)
	if count >= min && count <= max {
		return true
	}
	return false
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 3, "\n")
	for i, inp := range input {
		input[i] = strings.Repeat(inp, (len(input)/len(inp)) * 4)
	}
	mount := Mountain{}
	mount.ParseMap(input)
	mount.CurrentPosition = []int{0,0}
	inMap := true
	for ;inMap == true; {
		inMap = mount.GoDownOneLevel(3, 1)
	}
	return fmt.Sprintf("%d", mount.TreesEncountered)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 3, "\n")
	for i, inp := range input {
		input[i] = strings.Repeat(inp, (len(input)/len(inp)) * 8)
	}
	mount := Mountain{}
	mount.ParseMap(input)
	mount.CurrentPosition = []int{0,0}
	inMap := true
	for ;inMap == true; {
		inMap = mount.GoDownOneLevel(1, 1)
	}
	a := mount.TreesEncountered
	mount.TreesEncountered = 0

	mount.CurrentPosition = []int{0,0}
	inMap = true
	for ;inMap == true; {
		inMap = mount.GoDownOneLevel(3, 1)
	}
	b := mount.TreesEncountered
	mount.TreesEncountered = 0

	mount.CurrentPosition = []int{0,0}
	inMap = true
	for ;inMap == true; {
		inMap = mount.GoDownOneLevel(5, 1)
	}
	c := mount.TreesEncountered
	mount.TreesEncountered = 0

	mount.CurrentPosition = []int{0,0}
	inMap = true
	for ;inMap == true; {
		inMap = mount.GoDownOneLevel(7, 1)
	}
	d := mount.TreesEncountered
	mount.TreesEncountered = 0

	mount.CurrentPosition = []int{0,0}
	inMap = true
	for ;inMap == true; {
		inMap = mount.GoDownOneLevel(1, 2)
	}
	e := mount.TreesEncountered
	mount.TreesEncountered = 0
	return fmt.Sprintf("%d", a*b*c*d*e)
}