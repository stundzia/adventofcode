package day15

import (
	"fmt"
	"strconv"

	"github.com/stundzia/adventofcode/utils"
)

type cave struct {
	grid map[string]*node
	maxX int
	maxY int
}

type node struct {
	cave *cave
	val int
	x int
	y int
	n *node
	s *node
	w *node
	e *node
}

func (c *cave) getNode(x, y int) *node {
	return c.grid[fmt.Sprintf("%d,%d", x, y)]
}

func (n *node) getNeighbours() {
	n.n = n.cave.getNode(n.x, n.y - 1)
	n.s = n.cave.getNode(n.x, n.y + 1)
	n.e = n.cave.getNode(n.x + 1, n.y)
	n.w = n.cave.getNode(n.x - 1, n.y)
}

func newCave(nums []string) *cave {
	c := &cave{
		grid: map[string]*node{},
	}
	for y, line := range nums {
		for x, val := range line {
			valInt, _ := strconv.Atoi(string(val))
			c.grid[fmt.Sprintf("%d,%d", x, y)] = &node{val: valInt, x: x, y: y, cave: c}
			c.maxX = x
		}
		c.maxY = y
	}
	for _, n := range c.grid {
		n.getNeighbours()
	}
	return c
}

func newCavePart2(nums []string) *cave {
	c := &cave{
		grid: map[string]*node{},
	}
	maxY := len(nums)
	maxX := len(nums[0])
	for y, line := range nums {
		for x, val := range line {
			valInt, _ := strconv.Atoi(string(val))
			c.grid[fmt.Sprintf("%d,%d", x, y)] = &node{val: valInt, x: x, y: y, cave: c}

			for xMult := 0; xMult < 5; xMult++ {
				for yMult := 0; yMult < 5; yMult++ {
					if xMult + yMult == 0 {
						continue
					}
					newX := x + xMult * maxX
					newY := y + yMult * maxY
					newVal := valInt + xMult + yMult
					if newVal > 9 {
						newVal = newVal - 9
					}
					c.grid[fmt.Sprintf("%d,%d", newX, newY)] = &node{val: newVal, x: newX, y: newY, cave: c}

					if newX > c.maxX {
						c.maxX = newX
					}
					if newY > c.maxY {
						c.maxY = newY
					}
				}
			}
		}
	}
	for _, n := range c.grid {
		n.getNeighbours()
	}
	return c
}

func (c *cave) manhattenDistance(a, b *node) int {
	return utils.AbsInt(a.x - b.x) + utils.AbsInt(a.y - b.y)
}

func (c *cave) findLeastRiskyPathVal(start, end *node) int {
	riskMap := map[*node]int{start: start.val}
	for {
		didSomething := false
		for _, n := range c.grid {
			currVal, inMap := riskMap[n]
			if val, exists := riskMap[n.s]; exists {
				if !inMap || currVal > val + n.val {
					riskMap[n] = val + n.val
					didSomething = true
				}
			}
			if val, exists := riskMap[n.n]; exists {
				if !inMap || currVal > val + n.val {
					riskMap[n] = val + n.val
					didSomething = true
				}
			}
			if val, exists := riskMap[n.e]; exists {
				if !inMap || currVal > val + n.val {
					riskMap[n] = val + n.val
					didSomething = true
				}
			}
			if val, exists := riskMap[n.w]; exists {
				if !inMap || currVal > val + n.val {
					riskMap[n] = val + n.val
					didSomething = true
				}
			}
		}
		// Mega optimization lel
		if !didSomething {
			break
		}
	}
	return riskMap[end]
}

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 15, "\n")
	cve := newCave(nums)
	start := cve.getNode(0,0)
	end := cve.getNode(cve.maxX, cve.maxY)
	r := cve.findLeastRiskyPathVal(start, end)

	return fmt.Sprintf("Solution: %d", r - start.val)
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2021, 15, "\n")

	cve := newCavePart2(nums)
	start := cve.getNode(0,0)
	end := cve.getNode(cve.maxX, cve.maxY)
	r := cve.findLeastRiskyPathVal(start, end)

	return fmt.Sprintf("Solution: %d", r - start.val)
}
