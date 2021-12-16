package day18

import (
	"testing"
)

func TestEvaluateExpresionWithParenthesizedAddition(t *testing.T) {
	tcs := []struct {
		test     string
		input    string
		expected int
	}{
		{
			"(2 * 3) + 5 + (4 * 5)  Ans: 31",
			"(2 * 3) + 5 + (4 * 5)",
			31,
		},
		{
			"1 + (2 * 3) + (4 * (5 + 6))  Ans: 51",
			"1 + (2 * 3) + (4 * (5 + 6))",
			51,
		},
		{
			"2 * 3 + (4 * 5)  Ans: 46",
			"2 * 3 + (4 * 5)",
			46,
		},
		{
			"5 + (8 * 3 + 9 + 3 * 4 * 3)  Ans: 1445",
			"5 + (8 * 3 + 9 + 3 * 4 * 3)",
			1445,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))  Ans: 669060",
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			669060,
		},
		{
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2  Ans: 23340",
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			23340,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			exp := parenthesizeAddition(tc.input)
			ans := evaluateExpression(exp)
			if ans != tc.expected {
				t.Errorf("expected %d, but got %d", tc.expected, ans)
			}
		})
	}
}
