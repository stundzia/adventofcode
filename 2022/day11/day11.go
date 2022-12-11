package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maja42/goval"

	"github.com/stundzia/adventofcode/utils"
)

type monkeyGroup struct {
	monkeys map[int]*monkey
	allDivs int
}

type monkey struct {
	id              int
	group           *monkeyGroup
	itemsInspected  int
	items           []int
	operation       func(int) int
	test            func(int) bool
	testDiv         int
	testTrueTarget  int
	testFalseTarget int
}

func (mg *monkeyGroup) parseMonkey(lines []string) {
	m := &monkey{items: []int{}, group: mg}
	nameParts := strings.Split(lines[0], " ")
	id, _ := strconv.Atoi(strings.Split(nameParts[1], ":")[0])
	m.id = id
	startingItemsStr := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
	for _, itm := range startingItemsStr {
		itmInt, _ := strconv.Atoi(itm)
		m.items = append(m.items, itmInt)
	}
	m.operation = parseOperation(strings.Split(lines[2], ": ")[1])
	m.test = parseTest(strings.Split(lines[3], ": ")[1])

	testParts := strings.Split(lines[3], " ")
	div, _ := strconv.Atoi(testParts[len(testParts)-1])
	m.testDiv = div
	mg.allDivs *= div

	l4parts := strings.Split(lines[4], " ")
	l5parts := strings.Split(lines[5], " ")
	trueTarget, _ := strconv.Atoi(l4parts[len(l4parts)-1])
	falseTarget, _ := strconv.Atoi(l5parts[len(l5parts)-1])
	m.testTrueTarget = trueTarget
	m.testFalseTarget = falseTarget

	mg.monkeys[id] = m
}

func (mg *monkeyGroup) doRound() {
	for i := 0; i < len(mg.monkeys); i++ {
		mg.monkeys[i].doRound()
	}
}

func (mg *monkeyGroup) doRoundPart2() {
	for i := 0; i < len(mg.monkeys); i++ {
		mg.monkeys[i].doRoundPart2()
	}
}

func (m *monkey) doRound() {
	for len(m.items) > 0 {
		m.handleItem()
	}
}

func (m *monkey) doRoundPart2() {
	for len(m.items) > 0 {
		m.handleItemPart2()
	}
}

func (m *monkey) handleItem() {
	if len(m.items) == 0 {
		return
	}
	m.itemsInspected++
	item := m.items[0]
	m.items = m.items[1:]
	item = m.operation(item)
	item = item / 3
	if m.test(item) {
		m.group.monkeys[m.testTrueTarget].items = append(m.group.monkeys[m.testTrueTarget].items, item)
		return
	}
	m.group.monkeys[m.testFalseTarget].items = append(m.group.monkeys[m.testFalseTarget].items, item)
}

func (m *monkey) handleItemPart2() {
	if len(m.items) == 0 {
		return
	}
	m.itemsInspected++
	item := m.items[0]
	m.items = m.items[1:]
	item = m.operation(item)
	item = item % m.group.allDivs
	if m.test(item) {
		m.group.monkeys[m.testTrueTarget].items = append(m.group.monkeys[m.testTrueTarget].items, item)
		return
	}
	m.group.monkeys[m.testFalseTarget].items = append(m.group.monkeys[m.testFalseTarget].items, item)
}

func parseOperation(operation string) func(int) int {
	op := strings.Split(operation, " = ")[1]
	f := func(n int) int {
		eval := goval.NewEvaluator()
		res, err := eval.Evaluate(op, map[string]interface{}{"old": n}, nil)
		if err != nil {
			panic(err)
		}
		return res.(int)
	}
	return f
}

func parseTest(test string) func(int) bool {
	parts := strings.Split(test, " by ")
	v, _ := strconv.Atoi(parts[1])
	f := func(n int) bool {
		if n%v == 0 {
			return true
		}
		return false
	}
	return f
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2022, 11, "\n")
	monkeyLines := []string{}
	mg := monkeyGroup{map[int]*monkey{}, 1}
	for _, l := range input {
		if l == "" {
			mg.parseMonkey(monkeyLines)
			monkeyLines = []string{}
			continue
		}
		monkeyLines = append(monkeyLines, l)
	}
	mg.parseMonkey(monkeyLines)

	for i := 0; i < 20; i++ {
		mg.doRound()
	}

	top := 0
	second := 0
	for _, m := range mg.monkeys {
		if m.itemsInspected >= top {
			top = m.itemsInspected
		}
	}
	for _, m := range mg.monkeys {
		if m.itemsInspected < top && m.itemsInspected >= second {
			second = m.itemsInspected
		}
	}

	return fmt.Sprintf("Solution: %d", top*second)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2022, 11, "\n")
	monkeyLines := []string{}
	mg := monkeyGroup{map[int]*monkey{}, 1}
	for _, l := range input {
		if l == "" {
			mg.parseMonkey(monkeyLines)
			monkeyLines = []string{}
			continue
		}
		monkeyLines = append(monkeyLines, l)
	}
	mg.parseMonkey(monkeyLines)
	for i := 0; i < 10000; i++ {
		mg.doRoundPart2()
	}

	top := 0
	second := 0
	for _, m := range mg.monkeys {
		if m.itemsInspected >= top {
			top = m.itemsInspected
		}
	}
	for _, m := range mg.monkeys {
		if m.itemsInspected < top && m.itemsInspected >= second {
			second = m.itemsInspected
		}
	}

	return fmt.Sprintf("Solution: %d", top*second)
}
