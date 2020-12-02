package day2

import "testing"

func TestGetProductOfMatchingSumOfTwoItems(t *testing.T) {
	testSlice := []int{1721, 979, 366, 299, 675, 1456}
	res := getProductOfMatchingSumOfTwoItems(testSlice, 2020)
	expected := 514579
	if res != expected {
		t.Errorf("Expected %d, but got %d", expected, res)
	}
}

func TestGetProductOfMatchingSumOfThreeItems(t *testing.T) {
	testSlice := []int{1721, 979, 366, 299, 675, 1456}
	res := getProductOfMatchingSumOfThreeItems(testSlice, 2020)
	expected := 241861950
	if res != expected {
		t.Errorf("Expected %d, but got %d", expected, res)
	}
}
