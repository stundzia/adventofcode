package day12

import (
	"fmt"
	"golang.org/x/sync/semaphore"
	"strings"
	"sync"
	"unicode"

	"github.com/stundzia/adventofcode/utils"
)

type caveSystem struct {
	pathsMux sync.Mutex
	paths    map[string]struct{}
	start    *cave
	end      *cave
	caves    map[string]*cave
}

type cave struct {
	caveSystem  *caveSystem
	name        string
	connections map[string]*cave
	big         bool
	paths       map[string]struct{}
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
	c := cs.newCave(name)
	cs.caves[name] = c
	return c
}

func (cs *caveSystem) newCave(name string) *cave {
	big := false
	if len(name) <= 2 && unicode.IsUpper(rune(name[0])) {
		big = true
	}
	c := &cave{
		caveSystem:  cs,
		name:        name,
		big:         big,
		connections: map[string]*cave{},
		paths:       map[string]struct{}{},
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

func (cs *caveSystem) validatePathNoStart(path string) bool {
	itms := strings.Split(path, "-")
	usedCount := map[string]int{}
	for i, itm := range itms {
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

func (cs *caveSystem) validatePathNoStartPart2(path string) bool {
	itms := strings.Split(path, "-")
	usedCount := map[string]int{}
	for i, itm := range itms {
		if i == len(itms)-1 && itm != "end" {
			return false
		}
		if _, exists := usedCount[itm]; !exists {
			usedCount[itm] = 0
		}
		usedCount[itm]++
	}
	doubleEntryExists := false
	for itm, count := range usedCount {
		if (itm == "start" || itm == "end") && count > 1 {
			return false
		}
		if unicode.IsLower(rune(itm[0])) && count > 1 {
			if !doubleEntryExists {
				doubleEntryExists = true
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func (cs *caveSystem) tryRandomPath(sem *semaphore.Weighted) bool {
	path := []*cave{cs.start}
	current := cs.start
	visited := map[*cave]struct{}{}
	for i := 0; ; i++ {
		if current == cs.end {
			break
		}
		last := path[len(path)-1]
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
		cs.pathsMux.Lock()
		if _, exists := cs.paths[pathStr]; exists {
			cs.pathsMux.Unlock()
			sem.Release(1)
			return false
		}
		cs.paths[pathStr] = struct{}{}
		cs.pathsMux.Unlock()
	}
	sem.Release(1)
	return false
}

func doubleEntryInPath(path string) bool {
	parts := strings.Split(path, "-")
	countMap := map[string]int{}
	for _, part := range parts {
		if len(part) < 3 && unicode.IsLower(rune(part[0])) {
			if _, exists := countMap[part]; !exists {
				countMap[part] = 1
			} else {
				return true
			}
		}
	}
	return false
}

func (c *cave) getPathsToExit() {
	if c.caveSystem.end == c {
		c.paths = map[string]struct{}{c.name: {}}
		return
	}
	paths := map[string]struct{}{}
	for _, con := range c.connections {
	pathLoop:
		for path, _ := range con.paths {
			if !c.big {
				parts := strings.Split(path, "-")
				for _, p := range parts {
					if p == c.name {
						continue pathLoop
					}
				}
			}
			if c.caveSystem.validatePathNoStart(c.name + "-" + path) {
				paths[c.name+"-"+path] = struct{}{}
			}
		}
	}
	c.paths = paths
}

func (c *cave) getPathsToExit2() {
	if c.caveSystem.end == c {
		c.paths = map[string]struct{}{c.name: {}}
		return
	}
	paths := map[string]struct{}{}
	for _, con := range c.connections {
	pathLoop:
		for path, _ := range con.paths {
			deExists := doubleEntryInPath(path)
			if !c.big {
				parts := strings.Split(path, "-")
				for _, p := range parts {
					if p == c.name && deExists {
						continue pathLoop
					}
				}
			}
			if c.caveSystem.validatePathNoStartPart2(c.name + "-" + path) {
				paths[c.name+"-"+path] = struct{}{}
			}
		}
	}
	c.paths = paths
}

func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 12, "\n")
	cs := newCaveSystem(lines)
	cs.end.getPathsToExit()
	for i := 0; i < 20; i++ {
		for _, c := range cs.caves {
			c.getPathsToExit()
		}
	}
	return fmt.Sprintf("Solution: %d", len(cs.start.paths))
}

func DoGold() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 12, "\n")
	cs := newCaveSystem(lines)
	cs.end.getPathsToExit2()
	lastCount := 0
	for i := 0; i < 22; i++ {
		for _, c := range cs.caves {
			c.getPathsToExit2()
		}
		lastCount = len(cs.start.paths)
		fmt.Println("count: ", lastCount)
	}
	return fmt.Sprintf("Solution: %d", len(cs.start.paths))
}
