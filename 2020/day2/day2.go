package day2

import (
	"fmt"
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

func passValidV2(should int, shouldnt int, letter string, pass string) bool {
	if (string(pass[should - 1]) == letter && string(pass[shouldnt - 1]) != letter) || (string(pass[should - 1]) != letter && string(pass[shouldnt - 1]) == letter) {
		return true
	}
	return false
}


func handlePasswordAndPolicy(pp string) bool {
	passPolicy := strings.Split(pp, ": ")
	policy := passPolicy[0]
	pass := passPolicy[1]
	policyMinMaxLetter := strings.Split(policy, " ")
	letter := policyMinMaxLetter[1]
	minMax := strings.Split(policyMinMaxLetter[0], "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	return passValid(min, max, letter, pass)
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
	sum := 0
	for _, i := range input {
		if handlePasswordAndPolicy(i) {
			sum++
		}
	}
	return fmt.Sprintf(strconv.Itoa(sum))
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 2, "\n")
	sum := 0
	for _, i := range input {
		if handlePasswordAndPolicy2(i) {
			sum++
		}
	}
	return fmt.Sprintf(strconv.Itoa(sum))
}