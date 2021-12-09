package day9

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

type lavaCaveMap struct {
	points map[string]*point
}

type point struct {
	val int
	coordsX int
	coordsY int
	n *point
	s *point
	w *point
	e *point
	lowRisk bool
	basinSize int
}

func (lcm *lavaCaveMap) addPoint(x, y, val int) {
	lcm.points[fmt.Sprintf("%d,%d", x,y)] = &point{
		val:     val,
		coordsX: x,
		coordsY: y,
		n:       nil,
		s:       nil,
		w:       nil,
		e:       nil,
	}
}

func newLavaCaveMap(lines []string) *lavaCaveMap {
	lcm := &lavaCaveMap{points: map[string]*point{}}
	for y, l := range lines {
		xsStr := strings.Split(l, "")
		for x, val := range xsStr {
			v, _ := strconv.Atoi(val)
			lcm.addPoint(x, y, v)
		}
	}
	for _, point := range lcm.points {
		nCoords := fmt.Sprintf("%d,%d", point.coordsX, point.coordsY - 1)
		lowRisk := true
		if n, ok := lcm.points[nCoords]; ok {
			point.n = n
			if point.val >= n.val {
				lowRisk = false
			}
		}
		sCoords := fmt.Sprintf("%d,%d", point.coordsX, point.coordsY + 1)
		if s, ok := lcm.points[sCoords]; ok {
			point.s = s
			if point.val >= s.val {
				lowRisk = false
			}
		}
		wCoords := fmt.Sprintf("%d,%d", point.coordsX - 1, point.coordsY)
		if w, ok := lcm.points[wCoords]; ok {
			point.w = w
			if point.val >= w.val {
				lowRisk = false
			}
		}
		eCoords := fmt.Sprintf("%d,%d", point.coordsX + 1, point.coordsY)
		if e, ok := lcm.points[eCoords]; ok {
			point.e = e
			if point.val >= e.val {
				lowRisk = false
			}
		}
		point.lowRisk = lowRisk
	}
	return lcm
}

func (p *point) getRiskLevel() int {
	if !p.lowRisk {
		return 0
	}
	return p.val + 1
}

func (p *point) getBasinSize() int {
	size := 1
	doneMap := map[*point]struct{}{}
	doneMap[p] = struct{}{}
	if p.n != nil {
		add, dMap := p.n.getFlowSides("s", doneMap)
		doneMap = dMap
		size += add
	}
	if p.s != nil {
		add, dMap := p.s.getFlowSides("n", doneMap)
		doneMap = dMap
		size += add
	}
	if p.e != nil {
		add, dMap := p.e.getFlowSides("w", doneMap)
		doneMap = dMap
		size += add
	}
	if p.w != nil {
		add, dMap := p.w.getFlowSides("e", doneMap)
		doneMap = dMap
		size += add
	}
	return size
}

func (p *point) getFlowSides(flowDirection string, doneMap map[*point]struct{}) (int, map[*point]struct{}) {
	if _, ok := doneMap[p]; ok {
		return 0, doneMap
	}
	doneMap[p] = struct{}{}
	if p.val >= 9 {
		return 0, doneMap
	}
	sum := 1
	if flowDirection != "n" && p.n != nil && p.n.val > p.val {
		if _, ok := doneMap[p.n]; !ok {
			add, newMap := p.n.getFlowSides("s", doneMap)
			doneMap = newMap
			sum += add
		}
	}
	if flowDirection != "s" && p.s != nil && p.s.val > p.val {
		if _, ok := doneMap[p.s]; !ok {
			add, newMap := p.s.getFlowSides("n", doneMap)
			doneMap = newMap
			sum += add
		}
	}
	if flowDirection != "e" && p.e != nil && p.e.val > p.val {
		if _, ok := doneMap[p.e]; !ok {
			add, newMap := p.e.getFlowSides("w", doneMap)
			doneMap = newMap
			sum += add
		}
	}
	if flowDirection != "w" && p.w != nil && p.w.val > p.val {
		if _, ok := doneMap[p.w]; !ok {
			add, newMap := p.w.getFlowSides("e", doneMap)
			doneMap = newMap
			sum += add
		}

	}
	return sum, doneMap
}

func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 9, "\n")
	lcm := newLavaCaveMap(lines)
	riskSum := 0
	for _, p := range lcm.points {
		riskSum += p.getRiskLevel()
	}
	return fmt.Sprintf("Solution: %d", riskSum)
}

func DoGold() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 9, "\n")
	lcm := newLavaCaveMap(lines)
	basins := []int{}
	for _, p := range lcm.points {
		if p.lowRisk {
			basins = append(basins, p.getBasinSize())
		}
	}
	sort.Ints(basins)
	return fmt.Sprintf("Solution: %d", basins[len(basins) - 1] * basins[len(basins) - 2] * basins[len(basins) - 3])
}
