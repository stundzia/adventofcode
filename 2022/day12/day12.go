package day12

import (
	"fmt"
	"math"
	"sync/atomic"

	"github.com/stundzia/adventofcode/utils"
)

type elevationMap struct {
	nodes   map[string]*node
	visited map[*node]struct{}
	current *node
	target  *node
	start   *node
	steps   int
}

type node struct {
	em           *elevationMap
	coords       string
	height       int
	paths        map[*node]struct{}
	pathsTo      map[*node]struct{}
	pathToTarget map[*node]int
}

func newElevationMap(input []string) *elevationMap {
	em := &elevationMap{nodes: map[string]*node{}, visited: map[*node]struct{}{}}
	for y, row := range input {
		for x, p := range row {
			coords := fmt.Sprintf("%d,%d", x, y)
			if p == 83 {
				n := &node{
					em:           em,
					coords:       coords,
					height:       0,
					paths:        map[*node]struct{}{},
					pathsTo:      map[*node]struct{}{},
					pathToTarget: map[*node]int{},
				}
				// S
				em.current = n
				em.start = n
				em.nodes[coords] = n
				continue
			}
			if p == 69 {
				// E
				n := &node{
					em:           em,
					coords:       coords,
					height:       122 - 97,
					paths:        map[*node]struct{}{},
					pathsTo:      map[*node]struct{}{},
					pathToTarget: map[*node]int{},
				}
				n.pathToTarget[n] = 0
				em.target = n
				em.nodes[coords] = n
				continue
			}
			n := &node{
				em:           em,
				coords:       coords,
				height:       int(p - 97),
				paths:        map[*node]struct{}{},
				pathsTo:      map[*node]struct{}{},
				pathToTarget: map[*node]int{},
			}
			em.nodes[coords] = n
		}
	}

	for _, n := range em.nodes {
		n.mapPossiblePaths()
	}
	return em
}

func (em *elevationMap) currentElevation() int {
	return em.current.height
}

func (em *elevationMap) mapNodes() int {
	for _, node := range em.nodes {
		node.mapPossiblePaths()
	}
	return em.current.height
}

func (n *node) mapPossiblePaths() {
	x, y := utils.CoordsStrToInts(n.coords)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if dx != 0 && dy != 0 {
				continue
			}
			other, found := n.em.nodes[utils.CoordsIntsToStr(x+dx, y+dy)]
			if !found {
				continue
			}
			if other.height-1 <= n.height {
				n.paths[other] = struct{}{}
			}
			if n.height-1 <= other.height {
				n.pathsTo[other] = struct{}{}
			}
		}
	}
}

//func (em *elevationMap) validateAsPath(coords string) bool {
//	if v, f := em.nodes[coords]; f {
//		if v-1 <= em.currentElevation() {
//			if _, f := em.visited[coords]; !f {
//				return true
//			}
//		}
//	}
//
//	return false
//}

//func (em *elevationMap) goToLocation(coords string) {
//	em.steps++
//	em.current = coords
//	em.visited[coords] = struct{}{}
//}

func getMaxVal(m map[*node]int) int {
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func (n *node) mapPathToTarget() bool {
	updated := false
	if n.em.target == n {
		return updated
	}
	l := math.MaxInt
	for no, _ := range n.paths {
		if len(no.pathToTarget) > 0 && len(no.pathToTarget) < l {
			if _, f := no.pathToTarget[n]; f {
				continue
			}
			if len(n.pathToTarget) > 0 && len(no.pathToTarget) > len(n.pathToTarget)+1 {
				continue
			}
			n.pathToTarget = map[*node]int{}
			l = len(no.pathToTarget)
			max := 0
			for k, v := range no.pathToTarget {
				n.pathToTarget[k] = v
				if v > max {
					max = v
				}
			}
			n.pathToTarget[n] = max + 1
			updated = true
		}
	}
	return updated
}

//func (em *elevationMap) adjacentCoords() {
//	res := map[string]int{}
//	x, y := utils.CoordsStrToInts(em.current)
//	res[utils.CoordsIntsToStr(x+1,y)] = em.
//}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2022, 12, "\n")

	leastSteps := &atomic.Int64{}
	leastSteps.Store(math.MaxInt)

	elMap := newElevationMap(input)

	for i := 0; i < 10288; i++ {
		updated := 0
		for _, n := range elMap.nodes {
			if n.mapPathToTarget() {
				updated++
			}
		}
		fmt.Println("updated: ", updated)
		if updated == 0 {
			break
		}
	}

	return fmt.Sprintf("Solution: %d", len(elMap.start.pathToTarget)-1)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2022, 12, "\n")

	leastSteps := &atomic.Int64{}
	leastSteps.Store(math.MaxInt)

	elMap := newElevationMap(input)

	for i := 0; i < 10288; i++ {
		updated := 0
		for _, n := range elMap.nodes {
			if n.mapPathToTarget() {
				updated++
			}
		}
		fmt.Println("updated: ", updated)
		if updated == 0 {
			break
		}
	}

	min := math.MaxInt
	for _, n := range elMap.nodes {
		if n.height == 0 && len(n.pathToTarget) > 1 && len(n.pathToTarget) < min {
			min = len(n.pathToTarget)
		}
	}

	return fmt.Sprintf("Solution: %d", min-1)
}
