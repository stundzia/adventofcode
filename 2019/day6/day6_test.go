package day6

import (
	"testing"
)

func TestSystem(t *testing.T) {
	input := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	ss := NewSystemFromInput(input)
	if dDtc := ss.bodies["D"].getDistanceToCenter(); dDtc != 3 {
		t.Errorf("expected D to have distance 3 from center, got %d", dDtc)
	}
	if lDtc := ss.bodies["L"].getDistanceToCenter(); lDtc != 7 {
		t.Errorf("expected L to have distance 7 from center, got %d", lDtc)
	}
	if comDtc := ss.bodies["COM"].getDistanceToCenter(); comDtc != 0 {
		t.Errorf("expected COM to have distance 0 from center, got %d", comDtc)
	}
	if totalOrbits := ss.getTotalOrbits(); totalOrbits != 54 {
		t.Errorf("the total number of orbits should be 54, but got %d", totalOrbits)
	}
	if commonBody, steps := ss.findCommonBody("YOU", "SAN"); commonBody.name != "D" || steps != 6 {
		t.Errorf("expected common body to be D and steps to be 4, but got %s and %d respectively", commonBody.name, steps)
	}
}