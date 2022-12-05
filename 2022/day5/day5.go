package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

type stack struct {
	crates []string
}

func newStack() *stack {
	return &stack{crates: []string{}}
}

func (s *stack) in(crate string) {
	s.crates = append([]string{crate}, s.crates...)
}

func (s *stack) pop() string {
	if len(s.crates) == 0 {
		return ""
	}
	crate := s.crates[0]
	s.crates = s.crates[1:]
	return crate
}

func (s *stack) pop9001(count int) []string {
	if len(s.crates) < count {
		panic("wtf, count higher than size")
	}
	crates := s.crates[:count]
	s.crates = s.crates[count:]
	return crates
}

func parseSetupLine(l string) map[int]string {
	res := map[int]string{}
	ll := len(l)
	for i := 0; i < ll; i++ {
		var crate string
		switch true {
		case (i+1)*4 < ll:
			crate = l[i*4 : (i+1)*4]
		case (i+1)*4 == ll+1:
			crate = l[i*4 : (i+1)*4-1]
		}
		if crate == "" {
			return res
		}
		crate = strings.Replace(crate, "[", "", -1)
		crate = strings.Replace(crate, "]", "", -1)
		crate = strings.Replace(crate, " ", "", -1)
		if crate != "" {
			res[i] = crate
		}
	}
	return res
}

func parseActionLine(l string) (int, int, int) {
	parts := strings.Split(l, " ")
	count, _ := strconv.Atoi(parts[1])
	fromID, _ := strconv.Atoi(parts[3])
	toID, _ := strconv.Atoi(parts[5])
	fromID--
	toID--
	return count, fromID, toID
}

func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2022, 5, "\n")
	stacks := map[int]*stack{}
	for i := 7; i >= 0; i-- {
		r := parseSetupLine(lines[i])
		for id, crate := range r {
			if _, found := stacks[id]; !found {
				stacks[id] = newStack()
			}
			stacks[id].in(crate)
		}
	}

	for _, l := range lines {
		if strings.HasPrefix(l, "move") {
			count, fromID, toID := parseActionLine(l)
			for i := 0; i < count; i++ {
				c := stacks[fromID].pop()
				stacks[toID].in(c)
			}
		}
	}
	res := ""
	for i := 0; i < len(stacks); i++ {
		res = res + stacks[i].pop()
	}
	return fmt.Sprintf("Solution: " + res)
}

func DoGold() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2022, 5, "\n")
	stacks := map[int]*stack{}
	for i := 7; i >= 0; i-- {
		r := parseSetupLine(lines[i])
		for id, crate := range r {
			if _, found := stacks[id]; !found {
				stacks[id] = newStack()
			}
			stacks[id].in(crate)
		}
	}

	for _, l := range lines {
		if strings.HasPrefix(l, "move") {
			count, fromID, toID := parseActionLine(l)
			crates := stacks[fromID].pop9001(count)
			for i := len(crates) - 1; i >= 0; i-- {
				stacks[toID].in(crates[i])
			}
		}
	}

	res := ""
	for i := 0; i < len(stacks); i++ {
		res = res + stacks[i].pop()
	}
	return fmt.Sprintf("Solution: " + res)
}
