package day19

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strconv"
	"strings"
)

type RuleMatcher struct {
	RuleMap map[int]*Rule
}


type Rule struct {
	RM *RuleMatcher
	MatchingStrings []string
	RuleString string
	StringMatch string
	RuleMatch [][]int
}


func (r *Rule) MessageValid(msg string) bool {
	if len(r.MatchingStrings) == 0 {
		r.GetMatchingStrings()
	}
	return utils.StringSliceContains(r.MatchingStrings, msg)
}


func (r *Rule) GetMatchingStrings() []string {
	if len(r.MatchingStrings) > 0 {
		return r.MatchingStrings
	}
	if r.StringMatch != "" {
		return []string{r.StringMatch}
	}
	matchingStrings := []string{}
	for _, rm := range r.RuleMatch {
		if len(rm) == 3 {
			matches1 := r.RM.RuleMap[rm[0]].GetMatchingStrings()
			matches2 := r.RM.RuleMap[rm[1]].GetMatchingStrings()
			matches3 := r.RM.RuleMap[rm[2]].GetMatchingStrings()
			for _, m1 := range matches1 {
				for _, m2 := range matches2 {
					for _, m3 := range matches3 {
						matchingStrings = append(matchingStrings, m1 + m2 + m3)
					}
				}
			}
		} else if len(rm) == 2 {
			matches1 := r.RM.RuleMap[rm[0]].GetMatchingStrings()
			matches2 := r.RM.RuleMap[rm[1]].GetMatchingStrings()
			for _, m1 := range matches1 {
				for _, m2 := range matches2 {
					matchingStrings = append(matchingStrings, m1 + m2)
				}
			}
		} else if len(rm) == 1 {
			matchingStrings = append(matchingStrings, r.RM.RuleMap[rm[0]].GetMatchingStrings()...)
		} else {
			fmt.Println("Exception: ", r.RuleString)
			panic("Not like this...")
		}
	}
	r.MatchingStrings = matchingStrings
	return matchingStrings
}

func (rm *RuleMatcher) parseRule(rule string) {
	//80: 131 97
	//3: 123 107 | 97 66
	//101: 1 123 | 95 97
	//10: 97 138 | 123 107
	//83: 123 27 | 97 12
	//54: 97 126 | 123 107
	//61: 97 50 | 123 114
	//123: "a"
	parts := strings.Split(rule, ": ")
	ruleId, _ := strconv.Atoi(parts[0])
	ruleStr := parts[1]
	strMatch := ""
	ruleMatch := [][]int{}
	if len(ruleStr) == 3 && string(ruleStr[0]) == "\"" && string(ruleStr[2]) == "\"" {
		strMatch = string(ruleStr[1])
	} else {
		rmParts := strings.Split(ruleStr, " | ")
		for _, part := range rmParts {
			pp := strings.Split(part, " ")
			ruleMatch = append(ruleMatch, utils.SliceStringToInt(pp))
		}
	}
	ruleRef := &Rule{
		RM:              rm,
		MatchingStrings: nil,
		RuleString:      rule,
		StringMatch:     strMatch,
		RuleMatch:       ruleMatch,
	}
	rm.RuleMap[ruleId] = ruleRef
	fmt.Println(rm.RuleMap[ruleId].RuleString)
	fmt.Println(rm.RuleMap[ruleId].StringMatch)
	fmt.Println(rm.RuleMap[ruleId].RuleMatch)
}

func NewRuleMatcher(rules []string) *RuleMatcher {
	rm := &RuleMatcher{RuleMap: map[int]*Rule{}}
	for _, rule := range rules {
		rm.parseRule(rule)
	}
	return rm
}