package day12

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/stundzia/adventofcode/utils"
)

type caveSystem struct {
	paths map[string]struct{}
	start *cave
	end *cave
	caves map[string]*cave
}

type cave struct {
	name string
	connections map[string]*cave
	big bool
}

func newCaveSystem(lines []string) *caveSystem {
	cs := &caveSystem{
		paths: map[string]struct{}{},
		start: nil,
		end:   nil,
		caves: map[string]*cave{},
	}
	for _, line := range lines {
		parts := strings.Split(line, "-")
		c1 := cs.getOrCreateCave(parts[0])
		c2 := cs.getOrCreateCave(parts[1])
		c1.connections[c2.name] = c2
		c2.connections[c1.name] = c1
		if c1.name == "start" {
			cs.start = c1
		}
		if c2.name == "start" {
			cs.start = c2
		}
		if c1.name == "end" {
			cs.end = c1
		}
		if c2.name == "end" {
			cs.end = c2
		}
	}
	return cs
}

func (cs *caveSystem) getOrCreateCave(name string) *cave {
	if cave, exists := cs.caves[name]; exists {
		return cave
	}
	c := newCave(name)
	cs.caves[name] = c
	return c
}

func newCave(name string) *cave {
	big := false
	if len(name) <= 2 && unicode.IsUpper(rune(name[0])) {
		big = true
	}
	c := &cave{
		name: name,
		big: big,
		connections: map[string]*cave{},
	}
	return c
}

func (cs *caveSystem) validatePath(path string) bool {
	itms := strings.Split(path, "-")
	usedCount := map[string]int{}
	for i, itm := range itms {
		if i == 0 && itm != "start" {
			return false
		}
		if i == len(itms)-1 && itm != "end" {
			return false
		}
		if _, exists := usedCount[itm]; !exists {
			usedCount[itm] = 0
		}
		usedCount[itm]++
	}
	for itm, count := range usedCount {
		if unicode.IsLower(rune(itm[0])) && count > 1 {
			return false
		}
	}
	return true
}

func (cs *caveSystem) tryRandomPath() bool {
	path := []*cave{cs.start}
	current := cs.start
	visited := map[*cave]struct{}{}
	for i := 0;;i++ {
		if current == cs.end {
			break
		}
		last := path[len(path) - 1]
		for _, v := range current.connections {
			if _, wasVisited := visited[v]; wasVisited {
				continue
			}
			if v != cs.start && (last.big || last != v) {
				current = v
				path = append(path, current)
				if !v.big {
					visited[v] = struct{}{}
				}
				break
			}
		}
		if i > 10000 {
			break
		}
	}
	pathStr := ""
	for _, itm := range path {
		pathStr = pathStr + "-" + itm.name
	}
	pathStr = pathStr[1:]
	if cs.validatePath(pathStr) {
		if _, exists := cs.paths[pathStr]; exists {
			return false
		}
		cs.paths[pathStr] = struct{}{}
	}
	return false
}

func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 12, "\n")
	cs := newCaveSystem(lines)
	failedInARowCount := 0
	for ;; {
		if cs.tryRandomPath() {
			failedInARowCount = 0
		} else {
			failedInARowCount++
		}
		if failedInARowCount == 239625000 {
			break
		}
	}
	return fmt.Sprintf("Solution: %d", len(cs.paths))
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 12, "\n")
	return fmt.Sprintf("Solution: %d", len(nums))
}
