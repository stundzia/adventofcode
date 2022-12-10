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

type ropev2 struct {
	head *ropePartV2
	tail *ropePartV2
}

type ropePartV2 struct {
	head    *ropePartV2
	tail    *ropePartV2
	X       int
	Y       int
	visited map[string]struct{}
}

type ropePart struct {
	X       int
	Y       int
	visited map[string]struct{}
}

func (rp *ropePart) checkIn() {
	rp.visited[fmt.Sprintf("%d_%d", rp.X, rp.Y)] = struct{}{}
}

func (rp *ropePartV2) checkIn() {
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

func (r *ropev2) handle(cmd string) {
	parts := strings.Split(cmd, " ")
	d := parts[0]
	amount, _ := strconv.Atoi(parts[1])
	for i := 0; i < amount; i++ {
		r.moveHead(d)
	}

}

func (rp *ropePartV2) followHead() {
	rp.checkIn()
	if utils.AbsInt(rp.X-rp.head.X) > 1 && utils.AbsInt(rp.Y-rp.head.Y) > 1 {
		rp.X += moveAmount(rp.head.X, rp.X)
		rp.Y += moveAmount(rp.head.Y, rp.Y)
		rp.checkIn()
		return
	}
	if utils.AbsInt(rp.X-rp.head.X) > 1 {
		if utils.AbsInt(rp.Y-rp.head.Y) > 0 {
			rp.Y += rp.head.Y - rp.Y
		}
		rp.X += moveAmount(rp.head.X, rp.X)
		rp.checkIn()
	}
	if utils.AbsInt(rp.Y-rp.head.Y) > 1 {
		if utils.AbsInt(rp.X-rp.head.X) > 0 {
			rp.X += rp.head.X - rp.X
		}
		rp.Y += moveAmount(rp.head.Y, rp.Y)
		rp.checkIn()
	}
}

func moveAmount(cHead, cTail int) int {
	diff := cHead - cTail
	if diff == -2 {
		return -1
	}
	if diff == 2 {
		return 1
	}
	return 0
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

func (r *ropev2) moveHead(direction string) {
	r.tail.checkIn()
	r.head.checkIn()
	switch direction {
	case "U":
		r.head.Y++
	case "D":
		r.head.Y--
	case "L":
		r.head.X--
	case "R":
		r.head.X++
	}
	for t := r.head.tail; t.tail != nil; t = t.tail {
		t.followHead()
	}
	r.tail.followHead()
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
	r := &ropev2{
		head: &ropePartV2{
			X:       0,
			Y:       0,
			visited: map[string]struct{}{},
		},
	}
	rp := r.head
	for i := 0; i < 9; i++ {
		rp.tail = &ropePartV2{
			head:    rp,
			tail:    nil,
			X:       0,
			Y:       0,
			visited: map[string]struct{}{},
		}
		rp = rp.tail
	}
	r.tail = rp
	rp.checkIn()
	for _, cmd := range cmds {
		r.handle(cmd)
	}

	return fmt.Sprintf("Solution: %d", len(r.tail.visited))
}
