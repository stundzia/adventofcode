package day2

import (
	"fmt"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

func parseLine(line string) (string, string) {
	parts := strings.Split(line, " ")
	mine := parts[1]
	other := parts[0]
	return other, mine
}

// (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of
// the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
func parseOutcome(his string, mine string) int {
	switch his {
	case "A":
		switch mine {
		case "X":
			return 1 + 3
		case "Y":
			return 2 + 6
		case "Z":
			return 3 + 0
		}
	case "B":
		switch mine {
		case "X":
			return 1 + 0
		case "Y":
			return 2 + 3
		case "Z":
			return 3 + 6
		}
	case "C":
		switch mine {
		case "X":
			return 1 + 6
		case "Y":
			return 2 + 0
		case "Z":
			return 3 + 3
		}
	}
	panic("oh no")
}

// (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of
// the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
// Anyway, the second column says how the round needs to end:
// X means you need to lose,
// Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
func parseOutcomeV2(his string, result string) int {
	switch his {
	case "A":
		switch result {
		case "X":
			return 3 + 0
		case "Y":
			return 1 + 3
		case "Z":
			return 2 + 6
		}
	case "B":
		switch result {
		case "X":
			return 1 + 0
		case "Y":
			return 2 + 3
		case "Z":
			return 3 + 6
		}
	case "C":
		switch result {
		case "X":
			return 2 + 0
		case "Y":
			return 3 + 3
		case "Z":
			return 1 + 6
		}
	}
	panic("oh no")
}

func DoSilver() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2022, 2, "\n")
	res := 0
	for _, n := range nums {
		res += parseOutcome(parseLine(n))
	}

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	nums, _ := utils.ReadInputFileContentsAsStringSlice(2022, 2, "\n")
	res := 0
	for _, n := range nums {
		res += parseOutcomeV2(parseLine(n))
	}

	return fmt.Sprintf("Solution: %d", res)
}
