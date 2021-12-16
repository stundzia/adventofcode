package day4

import (
	"fmt"
	"strconv"
)

func isValid(pass int) bool {
	passStr := strconv.Itoa(pass)
	adjecentFound := false
	var last int32
	for _, char := range passStr {
		if char < last {
			return false
		}
		if char == last {
			adjecentFound = true
		}
		last = char
	}
	return adjecentFound
}

func isValidV2(pass int) bool {
	passStr := strconv.Itoa(pass)
	adjecentCount := 0
	adjecentFound := false
	var last int32
	for _, char := range passStr {
		if char < last {
			return false
		}
		if char == last {
			adjecentCount++
		} else {
			if adjecentCount == 1 {
				adjecentFound = true
			}
			adjecentCount = 0
		}
		last = char
	}
	return adjecentFound || adjecentCount == 1
}

func getPasswordCountInRange(min, max, version int) int {
	count := 0
	for i := min; i <= max; i++ {
		if version == 1 && isValid(i) {
			count++
		} else if version == 2 && isValidV2(i) {
			count++
		}
	}
	return count
}

func DoSilver() string {
	return fmt.Sprintf("Solution: %d", getPasswordCountInRange(197487, 673251, 1))
}

func DoGold() string {
	return fmt.Sprintf("Solution: %d", getPasswordCountInRange(197487, 673251, 2))
}
