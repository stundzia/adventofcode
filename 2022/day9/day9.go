package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

type rope struct {
	tail *ropePart
	head *ropePart
}

type ropePart struct {
	X       int
	Y       int
	visited map[string]struct{}
}

func (rp *ropePart) checkIn() {
	rp.visited[fmt.Sprintf("%d_%d", rp.X, rp.Y)] = struct{}{}
}

func (r *rope) handle(cmd string) {
	parts := strings.Split(cmd, " ")
	d := parts[0]
	amount, _ := strconv.Atoi(parts[1])
	for i := 0; i < amount; i++ {
		r.moveHead(d)
	}

}

func (r *rope) moveHead(direction string) {
	r.tail.checkIn()
	r.head.checkIn()
	switch direction {
	case "U":
		r.head.Y++
		if r.head.Y-r.tail.Y > 1 {
			r.tail.X += r.xDiff()
			r.tail.Y++
		}
	case "D":
		r.head.Y--
		if r.tail.Y-r.head.Y > 1 {
			r.tail.X += r.xDiff()
			r.tail.Y--
		}
	case "L":
		r.head.X--
		if r.tail.X-r.head.X > 1 {
			r.tail.Y += r.yDiff()
			r.tail.X--
		}
	case "R":
		r.head.X++
		if r.head.X-r.tail.X > 1 {
			r.tail.Y += r.yDiff()
			r.tail.X++
		}
	}
	r.head.checkIn()
	r.tail.checkIn()

}

func (r *rope) xDiff() int {
	res := r.head.X - r.tail.X
	return res
}

func (r *rope) yDiff() int {
	res := r.head.Y - r.tail.Y
	return res
}

func (r *rope) tailKeepUp() {
	if r.tail.Y == r.head.Y {
		if r.xDiff() > 1 {

		}
	}
}

func DoSilver() string {
	cmds, _ := utils.ReadInputFileContentsAsStringSlice(2022, 9, "\n")
	r := &rope{
		tail: &ropePart{
			X:       0,
			Y:       0,
			visited: map[string]struct{}{},
		},
		head: &ropePart{
			X:       0,
			Y:       0,
			visited: map[string]struct{}{},
		},
	}
	for _, cmd := range cmds {
		r.handle(cmd)
	}

	return fmt.Sprintf("Solution: %d", len(r.tail.visited))
}

func DoGold() string {
	cmds, _ := utils.ReadInputFileContentsAsStringSlice(2022, 9, "\n")
	r := &rope{
		tail: &ropePart{
			X:       0,
			Y:       0,
			visited: map[string]struct{}{},
		},
		head: &ropePart{
			X:       0,
			Y:       0,
			visited: map[string]struct{}{},
		},
	}
	for _, cmd := range cmds {
		r.handle(cmd)
	}

	return fmt.Sprintf("Solution: %d", len(r.tail.visited))
}
