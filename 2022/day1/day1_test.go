package day1

import (
	"testing"
)

func TestCountIncreases(t *testing.T) {
	testSlice := []int{1721, 979, 366, 299, 675, 1456, 2222, 45, 55}
	res := 0
	for _, v := range testSlice {
		res += v
	}
	expected := 7818
	if res != expected {
		t.Errorf("Expected %d, but got %d", expected, res)
	}
}

//
//func TestCountWindowIncreases(t *testing.T) {
//	testSlice := []int{1721, 979, 366, 299, 675, 1456, 2222, 45, 55, 1233, 14124, 1234}
//	res := countWindowIncreases(testSlice)
//	expected := 4
//	if res != expected {
//		t.Errorf("Expected %d, but got %d", expected, res)
//	}
//}
