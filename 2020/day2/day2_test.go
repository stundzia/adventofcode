package day2

import "testing"

func TestPassValid(t *testing.T) {
	if !passValid(1, 4, "c", "acce") {
		t.Errorf("Expected 'acce' to be valid")
	}
	if !passValid(1, 4, "d", "adcaae") {
		t.Errorf("Expected 'adcaae' to be valid")
	}
	if !passValid(1, 4, "d", "adcadaedd") {
		t.Errorf("Expected 'adcadaedd' to be valid")
	}
	if passValid(1, 4, "d", "dadcadaedd") {
		t.Errorf("Expected 'dadcadaedd' to be invalid")
	}
	if passValid(5, 7, "d", "lalddalalpawofdd") {
		t.Errorf("Expected 'lalddalalpawofdd' to be invalid")
	}
}

func TestPassValidV2(t *testing.T) {
	if passValidV2(1, 4, "c", "acce") {
		t.Errorf("Expected 'acce' to be invalid")
	}
	if passValidV2(1, 4, "a", "adcaae") {
		t.Errorf("Expected 'adcaae' to be invalid")
	}
	if !passValidV2(1, 5, "d", "adcadaedd") {
		t.Errorf("Expected 'adcadaedd' to be valid")
	}
	if !passValidV2(1, 4, "d", "dadcadaedd") {
		t.Errorf("Expected 'dadcadaedd' to be valid")
	}
	if !passValidV2(5, 7, "d", "lalddalalpawofdd") {
		t.Errorf("Expected 'lalddalalpawofdd' to be valid")
	}
}
