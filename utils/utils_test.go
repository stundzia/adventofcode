package utils

import (
	"math"
	"testing"
)

func TestSumIntSlice(t *testing.T) {
	tcs := []struct {
		testInput      []int
		expectedOutput int
	}{
		{
			[]int{1, 2, 3},
			6,
		},
		{
			[]int{55, 77, 0, 22, 0, -20},
			134,
		},
		{
			[]int{-33, 33, 44, -44},
			0,
		},
		{
			[]int{0, 0, 0, 0, 0, 0, 0},
			0,
		},
	}
	for _, tc := range tcs {
		if res := SumIntSlice(tc.testInput); res != tc.expectedOutput {
			t.Errorf("incorrect int slice sum, expected %d, but got %d", tc.expectedOutput, res)
		}
	}
}

func TestSlicesIntEqual(t *testing.T) {
	a := []int{123, 44, 55, -29, 42}
	b := []int{123, 44, 55, -29, 42}
	if !SlicesIntEqual(a, b) {
		t.Errorf("expected slices equal to be true")
	}
	a = []int{123, 44, 55, -29, 42, 0}
	b = []int{123, 44, 55, -29, 42}
	if SlicesIntEqual(a, b) {
		t.Errorf("expected slices equal to be false")
	}
	a = []int{123, 44, 55, -29, 41}
	b = []int{123, 44, 55, -29, 42}
	if SlicesIntEqual(a, b) {
		t.Errorf("expected slices equal to be false")
	}
}

func TestSliceStringToInt(t *testing.T) {
	testSlice := []string{"123", "44", "55", "-29", "42"}
	expected := []int{123, 44, 55, -29, 42}
	if !SlicesIntEqual(SliceStringToInt(testSlice), expected) {
		t.Errorf("converted slice did not match expected")
	}
}

func TestRotateCoords(t *testing.T) {
	tcs := []struct {
		test    string
		X       float64
		Y       float64
		degrees float64
		newX    float64
		newY    float64
	}{
		{
			"10,1 by 90 degrees",
			10.0,
			1.0,
			90.0,
			1.0,
			-10.0,
		},
		{
			"-4,5 by 180 degrees",
			-4.0,
			5.0,
			180.0,
			4.0,
			-5.0,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			if newX, newY := RotateCoordinates(tc.X, tc.Y, tc.degrees); math.Round(newX) != math.Round(tc.newX) || math.Round(newY) != math.Round(tc.newY) {
				t.Errorf("expected new coords to be %f, %f but got %f, %f", tc.newX, tc.newY, newX, newY)
			}
		})
	}
}
