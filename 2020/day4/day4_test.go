package day4

import (
	"testing"
)

func TestFieldValid(t *testing.T) {
	if !fieldValid("byr", "2002") {
		t.Errorf("Expected byr:2002 to be valid")
	}
	if fieldValid("byr", "2003") {
		t.Errorf("Expected byr:2003 to be invalid")
	}
	if !fieldValid("hgt", "60in") {
		t.Errorf("Expected hgt:60in to be valid")
	}
	if !fieldValid("hgt", "190cm") {
		t.Errorf("Expected hgt:190cm to be valid")
	}
	if fieldValid("hgt", "190in") {
		t.Errorf("Expected hgt:190in to be invalid")
	}
	if fieldValid("hgt", "190") {
		t.Errorf("Expected hgt:190 to be invalid")
	}
	if !fieldValid("hcl", "#123abc") {
		t.Errorf("Expected hcl:#123abc to be valid")
	}
	if fieldValid("hcl", "#123abz") {
		t.Errorf("Expected hcl:#123abz to be invalid")
	}
	if fieldValid("hcl", "123abc") {
		t.Errorf("Expected hcl:123abc to be invalid")
	}
	if !fieldValid("ecl", "brn") {
		t.Errorf("Expected ecl:brn to be valid")
	}
	if fieldValid("ecl", "wat") {
		t.Errorf("Expected ecl:wat to be invalid")
	}
	if !fieldValid("pid", "000000001") {
		t.Errorf("Expected pid:000000001 to be valid")
	}
	if fieldValid("pid", "0123456789") {
		t.Errorf("Expected pid:0123456789 to be invalid")
	}
}

