package day7

import (
	"github.com/stundzia/adventofcode/utils"
	"strconv"
	"strings"
)

func parseRule(rule string) (bag string, ruleMap map[string]int) {
	ruleSlice := strings.Split(rule, " contain ")
	topBag := strings.Replace(ruleSlice[0], " bags", "", -1)
	topBag = strings.Replace(topBag, " ", "_", -1)
	bagsString := strings.Replace(ruleSlice[1], " bags", "", -1)
	bagsString = strings.Replace(bagsString, " bag", "", -1)
	bagsString = strings.Replace(bagsString, ".", "", -1)
	bags := strings.Split(bagsString, ", ")
	ruleMap = map[string]int{}
	for _, b := range bags {
		count, _ := strconv.Atoi(string(b[0]))
		bagName := strings.Replace(b[2:], " ", "_", -1)
		ruleMap[bagName] = count
	}
	return topBag, ruleMap
}

func canContain(ruleMap map[string]map[string]int, bag string, shouldContain string) bool {
	allowed := map[string]interface{}{}
	for k, _ := range ruleMap[bag] {
		allowed[k] = struct {}{}
	}
	// TODO: don't brute force it
	for i := 0; i < 20; i++ {
		for k, _ := range allowed {
			for al, _ := range ruleMap[k] {
				allowed[al] = struct {}{}
			}
		}
	}
	_, ok := allowed[shouldContain]
	return ok
}

func mustContainCount(ruleMap map[string]map[string]int, bag string) int {
	rules := ruleMap[bag]
	totalCount := 0
	for bag, count := range rules {
		totalCount += count
		if _, ok := ruleMap[bag];ok {
			totalCount += count * mustContainCount(ruleMap, bag)
		}
	}
	return totalCount
}

func DoSilver() string {
	bagRules := map[string]map[string]int{}
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 7, "\n")
	for _, rule := range input {
		bag, ruleMap := parseRule(rule)
		bagRules[bag] = ruleMap
	}
	allowed := 0
	for bag, _ := range bagRules {
		if canContain(bagRules, bag, "shiny_gold") {
			allowed++
		}
	}
	return strconv.Itoa(allowed)
}

func DoGold() string {
	bagRules := map[string]map[string]int{}
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 7, "\n")
	for _, rule := range input {
		bag, ruleMap := parseRule(rule)
		bagRules[bag] = ruleMap
	}
	return strconv.Itoa(mustContainCount(bagRules, "shiny_gold"))
}
