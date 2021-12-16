package day10

import (
	"fmt"
	"sort"

	"github.com/stundzia/adventofcode/utils"
)

var chunkOpenerCloserMap = map[rune]rune{
	'[': ']',
	'(': ')',
	'{': '}',
	'<': '>',
}

var chunkCloserOpenerMap = map[rune]rune{
	']': '[',
	')': '(',
	'}': '{',
	'>': '<',
}

var scoreMapPart1 = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var scoreMapPart2 = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

type openerStack struct {
	openers []rune
	count   int
}

func (s *openerStack) Push(c rune) {
	s.openers = append(s.openers[:s.count], c)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *openerStack) Pop() rune {
	if s.count == 0 {
		return 0
	}
	s.count--
	return s.openers[s.count]
}

func newOpenerStack(size int) *openerStack {
	return &openerStack{
		openers: make([]rune, size),
	}
}

func parseLine(line string) int {
	openerStk := newOpenerStack(100)
	for _, c := range line {
		if _, ok := chunkOpenerCloserMap[c]; ok {
			openerStk.Push(c)
		} else {
			opener := chunkCloserOpenerMap[c]
			toMatch := openerStk.Pop()
			if toMatch == 0 {
				continue
			}
			if toMatch != opener {
				return scoreMapPart1[c]
			}
		}
	}
	return 0
}

func parseLine2(line string) int {
	openerStk := newOpenerStack(100)
	for _, c := range line {
		if _, ok := chunkOpenerCloserMap[c]; ok {
			openerStk.Push(c)
		} else {
			opener := chunkCloserOpenerMap[c]
			toMatch := openerStk.Pop()
			if toMatch == 0 {
				continue
			}
			if toMatch != opener {
				return 0
			}
		}
	}
	closingStr := ""
	for openerStk.count > 0 {
		closingStr = closingStr + string(chunkOpenerCloserMap[openerStk.Pop()])
	}
	score := 0
	for _, c := range closingStr {
		score = score * 5
		score = score + scoreMapPart2[c]
	}
	return score
}

func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 10, "\n")
	score := 0
	for _, l := range lines {
		score += parseLine(l)
	}
	return fmt.Sprintf("Solution: %d", score)
}

func DoGold() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 10, "\n")
	scores := []int{}
	for _, l := range lines {
		score := parseLine2(l)
		if score > 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return fmt.Sprintf("Solution: %d", scores[len(scores)/2])
}
