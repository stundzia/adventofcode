package day8

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

var digitSignalCountMap = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

var digitToSignalProperMap = map[int][]string {
	1: {"c", "f"},
	2: {"a", "c", "d", "e", "g"},
	3: {"a", "c", "d", "f", "g"},
	4: {"b", "c", "d", "f"},
}

func parseLine(line string) ([]string, []string) {
	lineParts := strings.Split(line, " | ")
	signals := strings.Split(lineParts[0], " ")
	outputs := strings.Split(lineParts[1], " ")
	return signals, outputs
}

func parseDigit(digitMap map[int]string, digitStr string) (map[int]string, int) {
	strLen := len(digitStr)
	for k, v := range digitMap {
		if digitStr == v {
			return digitMap, k
		}
	}
	//_, have8 := digitMap[8]
	_, have7 := digitMap[7]
	_, have6 := digitMap[6]
	_, have4 := digitMap[4]
	_, have3 := digitMap[3]
	_, have1 := digitMap[1]
	switch strLen {
	case 2:
		digitMap[1] = digitStr
		return digitMap, 1
	case 3:
		digitMap[7] = digitStr
		return digitMap, 7
	case 4:
		digitMap[4] = digitStr
		return digitMap, 4
	case 5:
		if have7 {
			if allLettersInOther(digitMap[7], digitStr) {
				digitMap[3] = digitStr
				return digitMap, 3
			}
		}
		if have1 {
			if allLettersInOther(digitMap[1], digitStr) {
				digitMap[3] = digitStr
				return digitMap, 3
			}
		}

		if have4 {
			matchLetters := matchingLettersInOther(digitStr, digitMap[4])
			if matchLetters == 2 {
				digitMap[2] = digitStr
				return digitMap, 2
			}
			if matchLetters == 3 {
				digitMap[5] = digitStr
				return digitMap, 5
			}
		}
		if have6 {
			if allLettersInOther(digitStr, digitMap[6]) {
				digitMap[5] = digitStr
				return digitMap, 5
			}
		}
		//matchLetters := matchingLettersInOther(digitStr, digitMap[4])
		//if matchLetters == 2 {
		//	return 2
		//}
		//if matchLetters == 3 {
		//	return 5
		//}
	case 6:
		if have7 && !allLettersInOther(digitMap[7], digitStr) {
			digitMap[6] = digitStr
			return digitMap, 6
		}
		if have4 && allLettersInOther(digitMap[4], digitStr) {
			digitMap[9] = digitStr
			return digitMap, 9
		}
		if have4 && have7 {
			digitMap[0] = digitStr
			return digitMap, 0
		}
		if have3 && allLettersInOther(digitMap[3], digitStr) {
			digitMap[9] = digitStr
			return digitMap, 9
		}
	case 7:
		digitMap[8] = digitStr
		return digitMap, 8
	}
	return digitMap, -1
}

func allLettersInOther(main string, other string) bool {
	mainLoop:
	for _, s := range main {
		for _, o := range other {
			if s == o {
				continue mainLoop
			}
		}
		return false
	}
	return true
}

func matchingLettersInOther(main string, other string) int {
	matches := 0
mainLoop:
	for _, s := range main {
		for _, o := range other {
			if s == o {
				matches++
				continue mainLoop
			}
		}
	}
	return matches
}

func parseLine2(line string) int {
	digitToLetterMap := map[int]string{}
	lineParts := strings.Split(line, " | ")
	signals := strings.Split(lineParts[0], " ")
	outputs := strings.Split(lineParts[1], " ")
	for _, s := range signals {
		switch len(s) {
		case 2:
			digitToLetterMap[1] = s
		case 3:
			digitToLetterMap[7] = s
		case 4:
			digitToLetterMap[4] = s
		case 7:
			digitToLetterMap[8] = s
		}
	}
	output := ""
	for _, s := range signals {
		digitToLetterMap, _ = parseDigit(digitToLetterMap, s)
	}
	for _, o := range outputs {
		_, oo := parseDigit(digitToLetterMap, o)
		if oo == -1 {
			log.Fatal("FUUUUUUUCK")
		}
		output += strconv.Itoa(oo)
	}
	outputInt, _ := strconv.Atoi(output)
	return outputInt
}

func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 8, "\n")
	count := 0
	for _, l := range lines {
		_, outputs := parseLine(l)
		for _, o := range outputs {
			if _, ok := digitSignalCountMap[len(o)]; ok {
				count++
			}
		}
	}
	return fmt.Sprintf("Solution: %d", count)
}

func DoGold() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 8, "\n")
	res := 0
	for _, l := range lines {
		d := parseLine2(l)
		fmt.Println("d: ", d)
		res += d
	}
	return fmt.Sprintf("Solution: %d", res)
}
