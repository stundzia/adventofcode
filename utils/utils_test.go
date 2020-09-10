package utils

import (
	"fmt"
	"testing"
)

func TestReadInputFileContentsAsString(t *testing.T) {
	contentString, err := ReadInputFileContentsAsString(3042, 42)
	exp := `this is a test file for utils testing
this one is for testing loading content as a string
as well as a string slice
using
newlines
as a
separator`
	if err != nil {
		t.Errorf("error loading file: %s", err)
	}
	if contentString != exp {
		t.Error("content mismatch")
	}
}

func TestReadInputFileContentsAsStringSlice(t *testing.T) {
	contentStringSlice, err := ReadInputFileContentsAsStringSlice(3042, 42, "\n")
	exp := []string{
		"this is a test file for utils testing",
		"this one is for testing loading content as a string",
		"as well as a string slice",
		"using",
		"newlines",
		"as a",
		"separator",
	}
	if err != nil {
		t.Errorf("error loading file: %s", err)
	}
	for index, value := range exp {
		if contentStringSlice[index] != value {
			t.Error("content mismatch")
		}
	}
}

func TestReadInputFileContentsAsIntSlice(t *testing.T) {
	contentIntSlice, err := ReadInputFileContentsAsIntSlice(3042, 106, ",")
	exp := []int{123,55,66,7,8,9,2024}
	if err != nil {
		t.Errorf("error loading file: %s", err)
	}
	for index, value := range exp {
		if contentIntSlice[index] != value {
			t.Error(fmt.Sprintf("content mismatch, want %d, but got %d", value, contentIntSlice[index]))
		}
	}
}

func TestSumIntSlice(t *testing.T) {
	tcs := []struct {
		testInput    []int
		expectedOutput int
	}{
		{
			[]int{1,2,3},
			6,
		},
		{
			[]int{55,77,0,22,0,-20},
			134,
		},
		{
			[]int{-33,33,44,-44},
			0,
		},
		{
			[]int{0,0,0,0,0,0,0},
			0,
		},
	}
	for _, tc := range tcs {
		if res := SumIntSlice(tc.testInput); res != tc.expectedOutput {
			t.Errorf("incorrect int slice sum, expected %d, but got %d", tc.expectedOutput, res)
		}
	}
}