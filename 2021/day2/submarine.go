package day2

import (
	"strconv"
	"strings"
)

type Submarine struct {
	Depth int
	Position int
	Aim int
}

func (s *Submarine) handleCommand(command string) {
	cmd := strings.Split(command, " ")[0]
	numS := strings.Split(command, " ")[1]
	num, _ := strconv.Atoi(numS)
	switch cmd {
	case "forward":
		s.Position += num
	case "down":
		s.Depth += num
	case "up":
		s.Depth -= num
	}
}

func (s *Submarine) handleCommandV2(command string) {
	cmd := strings.Split(command, " ")[0]
	numS := strings.Split(command, " ")[1]
	num, _ := strconv.Atoi(numS)
	switch cmd {
	case "forward":
		s.Position += num
		s.Depth += num * s.Aim
	case "down":
		s.Aim += num
	case "up":
		s.Aim -= num
	}
}
