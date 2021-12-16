package day2

import (
	"strconv"
	"strings"
)

type submarine struct {
	Depth    int
	Position int
	Aim      int
}

func (s *submarine) handleCommand(command string) {
	parts := strings.Split(command, " ")
	cmd := parts[0]
	amountString := parts[1]
	amount, _ := strconv.Atoi(amountString)
	switch cmd {
	case "forward":
		s.Position += amount
	case "down":
		s.Depth += amount
	case "up":
		s.Depth -= amount
	}
}

func (s *submarine) handleCommandV2(command string) {
	parts := strings.Split(command, " ")
	cmd := parts[0]
	amountString := parts[1]
	amount, _ := strconv.Atoi(amountString)
	switch cmd {
	case "forward":
		s.Position += amount
		s.Depth += amount * s.Aim
	case "down":
		s.Aim += amount
	case "up":
		s.Aim -= amount
	}
}
