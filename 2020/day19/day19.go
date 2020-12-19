package day19

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strings"
)


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 19, "\n\n")
	rules := strings.Split(input[0], "\n")
	msgs := strings.Split(input[1], "\n")
	rm := NewRuleMatcher(rules, 100)
	validCount := 0
	for _, msg := range msgs {
		if rm.RuleMap[0].MessageValid(msg) {
			validCount++
		}
	}
	return fmt.Sprintf("%d", validCount)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 19, "\n\n")
	rules := strings.Split(input[0], "\n")
	msgs := strings.Split(input[1], "\n")

	maxLen := 1
	for _, msg := range msgs {
		if len(msg) > maxLen {
			maxLen = len(msg)
		}
	}
	rm := NewRuleMatcher(rules, maxLen)
	// Part 2 condition:
	rm.RuleMap[8].RuleString = "8: 42 | 42 8"
	rm.RuleMap[8].RuleString = "8: 42 | 42 42 | 42 42 42 | 42 42 42 42"
	//rm.RuleMap[8].RuleMatch = [][]int{{42}, {42, 8}}
	rm.RuleMap[8].RuleMatch = [][]int{{42}, {42, 42}, {42, 42, 42}, {42, 42, 42, 42}, {42, 42, 42, 42, 42}}
	//rm.RuleMap[11].RuleString = "11: 42 31 | 42 11 31"
	rm.RuleMap[11].RuleString = "11: 42 31 | 42 42 31 31"
	//rm.RuleMap[11].RuleMatch = [][]int{{42, 31}, {42, 11, 31}}
	rm.RuleMap[11].RuleMatch = [][]int{{42, 31}, {42, 42, 31, 31}}
	// Basically the rule is "starts with rule 42, and ends with rule 31"


	validCount := 0
	for _, msg := range msgs {
		if rm.MessageValidPart2v2(msg) {
			validCount++
		}
	}
	return fmt.Sprintf("%d", validCount)
}
