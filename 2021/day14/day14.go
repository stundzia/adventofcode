package day14

import (
	"fmt"
	"math"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

func toPairs(el string) map[int]string {
	res := map[int]string{}
	for i := 0; i < len(el) - 1; i++ {
		res[i] = string(el[i]) + string(el[i+1])
	}
	return res
}

func letterMapInit(el string) map[string]int {
	res := map[string]int{}
	pairs := toPairs(el)
	fmt.Println(el)
	fmt.Println(pairs)
	for _, pair := range pairs {
		if _, exists := res[pair]; !exists {
			res[pair] = 0
		}
		res[pair]++
	}
	return res
}

func pairToMap(pair string, ruleMap map[string]string) map[string]int {
	res := map[string]int{}
	if val, exists := ruleMap[pair]; exists {
		res[val[:2]] = 1
		res[val[1:]] = 1
	} else {
		res[pair] = 1
	}
	return res
}

func pairMapStep(pairMap map[string]int, ruleMap map[string]string) map[string]int {
	res := map[string]int{}
	for pair, count := range pairMap {
		if count < 1 {
			continue
		}
		nm := pairToMap(pair, ruleMap)
		for pair, c := range nm {
			if _, exists := res[pair]; !exists {
				res[pair] = 0
			}
			res[pair] = res[pair] + (count * c)
		}
	}
	return res
}

func getResFromPairMap(pairMap map[string]int) int {
	letterMap := map[uint8]int{}
	for k, v := range pairMap {
		if _, exists := letterMap[k[0]]; !exists {
			letterMap[k[0]] = 0
		}
		if _, exists := letterMap[k[1]]; !exists {
			letterMap[k[1]] = 0
		}
		letterMap[k[0]] += v
		letterMap[k[1]] += v
	}
	min := math.MaxInt
	max := 0
	for _, v := range letterMap {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	// + 1 due to floor to int in case of uneven number
	return (max - min + 1) / 2
}

func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 14, "\n")
	el := lines[0]
	ruleMap := map[string]string{}
	for _, rule := range lines[2:] {
		parts := strings.Split(rule, " -> ")
		ruleMap[parts[0]] = string(parts[0][0]) + parts[1] + string(parts[0][1])
	}

	lm := letterMapInit(el)
	for i := 0; i < 10; i++ {
		lm = pairMapStep(lm, ruleMap)
	}

	return fmt.Sprintf("Solution: %d", getResFromPairMap(lm)) // +- 1
}



func DoGold() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 14, "\n")
	el := lines[0]
	ruleMap := map[string]string{}
	for _, rule := range lines[2:] {
		parts := strings.Split(rule, " -> ")
		ruleMap[parts[0]] = string(parts[0][0]) + parts[1] + string(parts[0][1])
	}

	lm := letterMapInit(el)
	for i := 0; i < 40; i++ {
		lm = pairMapStep(lm, ruleMap)
	}

	return fmt.Sprintf("Solution: %d", getResFromPairMap(lm)) // +- 1
}
