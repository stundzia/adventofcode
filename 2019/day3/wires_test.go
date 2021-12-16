package day3

import (
	"testing"
)

func TestParseCommand(t *testing.T) {
	tcs := []struct {
		test             string
		expectedDistance int
		expectedDY       int
		expectedDX       int
	}{
		{
			"U234",
			234,
			1,
			0,
		},
		{
			"D23",
			23,
			-1,
			0,
		},
		{
			"R13",
			13,
			0,
			1,
		},
		{
			"L42",
			42,
			0,
			-1,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			if dX, dY, distance := parseCommand(tc.test); distance != tc.expectedDistance || dX != tc.expectedDX || dY != tc.expectedDY {
				t.Errorf("expected delta X/delta Y/distance to be equal to %d/%d/%d, but got %d/%d/%d", tc.expectedDX, tc.expectedDY, tc.expectedDistance, dX, dY, distance)
			}
		})
	}
}

func TestMarkPosition(t *testing.T) {
	tcs := []struct {
		test         string
		signature    uint8
		position     uint8
		expectedMark uint8
	}{
		{
			"signature 1, position 2",
			1,
			2,
			3,
		},
		{
			"signature 2, position 0",
			2,
			0,
			2,
		},
		{
			"signature 1, position 1",
			1,
			1,
			1,
		},
		{
			"signature 2, position 2",
			2,
			2,
			2,
		},
		{
			"signature 2, position 1",
			2,
			1,
			3,
		},
		{
			"signature 2, position 3",
			2,
			3,
			3,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			w := &Wire{
				signature: tc.signature,
				grid:      nil,
				x:         0,
				y:         0,
			}
			w.markPosition(tc.position)
			if mark := w.markPosition(tc.position); mark != tc.expectedMark {
				t.Errorf("expected mark to be %d, but got %d", tc.expectedMark, mark)
			}
		})
	}
}

func TestGetDistanceToCenter(t *testing.T) {
	grid := NewGrid()
	dist := grid.getDistanceToCenter(527, -323)
	if dist != 850 {
		t.Errorf("expected distance to be 850, got %d", dist)
	}
}
