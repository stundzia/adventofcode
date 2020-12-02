package day2

import (
	"github.com/stundzia/adventofcode/utils"
	"strconv"
	"strings"
)

func passValid(min int, max int, letter string, pass string) bool {
	count := strings.Count(pass, letter)
	if count >= min && count <= max {
		return true
	}
	return false
}

func passValidV2(position1 int, position2 int, letter string, pass string) bool {
	matchesFirst := string(pass[position1 - 1]) == letter
	matchesSecond := string(pass[position2 - 1]) == letter
	if matchesFirst != matchesSecond {
		return true
	}
	return false
}


func parsePasswordAndPolicy(pp string) (num1, num2 int, letter, pass string) {
	passPolicy := strings.Split(pp, ": ")
	policy := passPolicy[0]
	pass = passPolicy[1]
	policyMinMaxLetter := strings.Split(policy, " ")
	letter = policyMinMaxLetter[1]
	minMax := strings.Split(policyMinMaxLetter[0], "-")
	num1, _ = strconv.Atoi(minMax[0])
	num2, _ = strconv.Atoi(minMax[1])
	return num1, num2, letter, pass
}


func handlePasswordAndPolicy(pp string) bool {
	minPos, maxPos, letter, pass := parsePasswordAndPolicy(pp)
	return passValid(minPos, maxPos, letter, pass)
}

func handlePasswordAndPolicy2(pp string) bool {
	passPolicy := strings.Split(pp, ": ")
	policy := passPolicy[0]
	pass := passPolicy[1]
	policyMinMaxLetter := strings.Split(policy, " ")
	letter := policyMinMaxLetter[1]
	minMax := strings.Split(policyMinMaxLetter[0], "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	return passValidV2(min, max, letter, pass)
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 2, "\n")
	validCount := 0
	for _, i := range input {
		if handlePasswordAndPolicy(i) {
			validCount++
		}
	}
	return strconv.Itoa(validCount)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 2, "\n")
	validCount := 0
	for _, i := range input {
		if handlePasswordAndPolicy2(i) {
			validCount++
		}
	}
	return strconv.Itoa(validCount)
}