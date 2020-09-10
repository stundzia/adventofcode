package day1

import "testing"

func TestGetModuleFuelReq(t *testing.T) {
	tcs := []struct {
		testMass    int
		expectedFuelMass int
	}{
		{
			12,
			2,
		},
		{
			14,
			2,
		},
		{
			1969,
			654,
		},
		{
			100756,
			33583,
		},
	}
	for _, tc := range tcs {
		if res := getMassFuelReq(tc.testMass); res != tc.expectedFuelMass {
			t.Errorf("incorrect fuel mass for module mass %d; expected %d, but got %d", tc.testMass, tc.expectedFuelMass, res)
		}
	}
}
