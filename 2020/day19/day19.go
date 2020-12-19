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
	rm := NewRuleMatcher(rules)
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
	return fmt.Sprintf("%d", len(input))
}
