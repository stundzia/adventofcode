package day2

import (
	"testing"
)

func TestSubmarineHandleCommand(t *testing.T) {
	testCommands := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	sub := submarine{}
	for _, cmd := range testCommands {
		sub.handleCommand(cmd)
	}
	expectedDepth := 10
	expectedPosition := 15
	if sub.Depth != expectedDepth {
		t.Errorf("Expected depth to be %d, but got %d", expectedDepth, sub.Depth)
	}
	if sub.Position != expectedPosition {
		t.Errorf("Expected position to be %d, but got %d", expectedPosition, sub.Position)
	}
}

func TestSubmarineHandleCommandV2(t *testing.T) {
	testCommands := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	sub := submarine{}
	for _, cmd := range testCommands {
		sub.handleCommandV2(cmd)
	}
	expectedDepth := 60
	expectedPosition := 15
	if sub.Depth != expectedDepth {
		t.Errorf("Expected depth to be %d, but got %d", expectedDepth, sub.Depth)
	}
	if sub.Position != expectedPosition {
		t.Errorf("Expected position to be %d, but got %d", expectedPosition, sub.Position)
	}
}
