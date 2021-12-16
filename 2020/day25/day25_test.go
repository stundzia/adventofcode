package day25

import (
	"testing"
)

func TestGetLoopSize(t *testing.T) {
	tcs := []struct {
		test     string
		key      int
		expected int
	}{
		{
			"Pkey: 5764801, expected loopsize 8",
			5764801,
			8,
		},
		{
			"Pkey: 17807724, expected loopsize 11",
			17807724,
			11,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			loopSize := getLoopSize(tc.key)
			if loopSize != tc.expected {
				t.Errorf("expected %d, but got %d", tc.expected, loopSize)
			}
		})
	}
}

func TestTransform(t *testing.T) {
	tcs := []struct {
		test     string
		key      int
		loopSize int
		expected int
	}{
		{
			"Pkey: 17807724, loopsize 8, expected: 14897079",
			17807724,
			8,
			14897079,
		},
		{
			"Pkey: 5764801, loopsize 11, expected: 14897079",
			5764801,
			11,
			14897079,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			encryptionKey := transform(tc.key, tc.loopSize)
			if encryptionKey != tc.expected {
				t.Errorf("expected %d, but got %d", tc.expected, encryptionKey)
			}
		})
	}
}
