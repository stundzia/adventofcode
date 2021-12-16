package day4

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tcs := []struct {
		test     string
		pass     int
		expected bool
	}{
		{
			"pass: 111111, valid: yes",
			111111,
			true,
		},
		{
			"pass: 223450, valid: no",
			223450,
			false,
		},
		{
			"pass: 123789, valid: no",
			123789,
			false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			valid := isValid(tc.pass)
			if valid != tc.expected {
				t.Errorf("expected %v, but got %v", tc.expected, valid)
			}
		})
	}
}

func TestIsValidV2(t *testing.T) {
	tcs := []struct {
		test     string
		pass     int
		expected bool
	}{
		{
			"pass: 112233, valid: yes",
			112233,
			true,
		},
		{
			"pass: 123444, valid: no",
			123444,
			false,
		},
		{
			"pass: 111122, valid: yes",
			111122,
			true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			valid := isValidV2(tc.pass)
			if valid != tc.expected {
				t.Errorf("expected %v, but got %v", tc.expected, valid)
			}
		})
	}
}
