package day8

import (
	"testing"
)

func TestParseInstruction(t *testing.T) {
	tcs := []struct{
		Case string
		Op string
		Arg int
	}{
		{
			"acc +50",
			"acc",
			50,
		},
		{
			"nop -21",
			"nop",
			-21,
		},
		{
			"jmp +461",
			"jmp",
			461,
		},
	}
	for _, tc := range tcs {
		if op, arg := parseInstruction(tc.Case); op != tc.Op || arg != tc.Arg {
			t.Errorf("case `%s` erroneusly got op, arg: %s %d", tc.Case, op, arg)
		}
	}
}
