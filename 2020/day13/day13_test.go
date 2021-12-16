package day13

import "testing"

func TestBusSchedule_GetFirstSequentialTimestamp(t *testing.T) {
	tcs := []struct {
		test     string
		input    string
		expected uint64
	}{
		{
			"Buses: 17,x,13,19",
			"17,x,13,19",
			3417,
		},
		{
			"Buses: 1789,37,47,1889",
			"1789,37,47,1889",
			1202161486,
		},
		{
			"Buses: 67,x,7,59,61",
			"67,x,7,59,61",
			779210,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.test, func(t *testing.T) {
			bs := NewBusSchedule(tc.input)
			if timestamp := bs.GetFirstSequentialTimestamp(); timestamp != tc.expected {
				t.Errorf("expected timestamp to be %d, but got %d", tc.expected, timestamp)
			}
		})
	}
}
