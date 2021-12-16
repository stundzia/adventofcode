package day19

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strconv"
	"strings"
)

type RuleMatcher struct {
	RuleMap map[int]*Rule
	MaxLen  int
}

type Rule struct {
	RM              *RuleMatcher
	MatchingStrings []string
	RuleString      string
	StringMatch     string
	RuleMatch       [][]int
	MaxLen          int
	MinLen          int
}

func (r *Rule) MessageValid(msg string) bool {
	if len(r.MatchingStrings) == 0 {
		r.GetMatchingStrings()
	}
	valid := utils.StringSliceContains(r.MatchingStrings, msg)

	return valid
}

func (rm *RuleMatcher) MessageValidPart2v2(msg string) bool {
	lr := rm.RuleMap[42]
	rr := rm.RuleMap[31]
	lr.GetMinMaxLen()
	rr.GetMinMaxLen()
	// These are all the same
	lrMin := lr.MinLen
	lrMax := lr.MaxLen
	rrMin := rr.MinLen
	rrMax := rr.MaxLen

	msgLen := len(msg)
	valid := false

	if msgLen%rrMax != 0 || msgLen < rrMin {
		return false
	}

	lastMatchEnd := 0
	Matched := 0

	rightMatched := false
	leftMatched := false
	leftMatchedCount := 0
	rightMatchedCount := 0

mainLoop:
	for i := lrMin; i <= msgLen; i++ {
		check := msg[lastMatchEnd:i]
		checkLen := len(check)

		if checkLen < rrMin && checkLen < lrMin {
			continue
		}

		if !rightMatched {
			if checkLen > lrMax {
				break mainLoop
			}
			if lr.MessageValid(check) {
				leftMatched = true
				leftMatchedCount++
				lastMatchEnd = i
				Matched += checkLen
			} else if rr.MessageValid(check) {
				if !leftMatched {
					break mainLoop
				}
				rightMatched = true
				rightMatchedCount++
				lastMatchEnd = i
				Matched += checkLen
			}
		} else {
			if checkLen > rrMax {
				break mainLoop
			}
			if rr.MessageValid(check) {
				lastMatchEnd = i
				rightMatchedCount++
				Matched += checkLen
			}
		}

	}
	if Matched == msgLen && leftMatched && rightMatched && leftMatchedCount >= 2 && leftMatchedCount > rightMatchedCount {
		valid = true
	}

	return valid
}

func (r *Rule) GetMatchingStrings() []string {
	if len(r.MatchingStrings) > 0 {
		return r.MatchingStrings
	}
	if r.StringMatch != "" {
		return []string{r.StringMatch}
	}
	matchingStringsMap := map[string]struct{}{}
	matchingStrings := []string{}
	for _, rm := range r.RuleMatch {
		if len(rm) == 5 {
			matches1 := r.RM.RuleMap[rm[0]].GetMatchingStrings()
			matches2 := r.RM.RuleMap[rm[1]].GetMatchingStrings()
			matches3 := r.RM.RuleMap[rm[2]].GetMatchingStrings()
			matches4 := r.RM.RuleMap[rm[3]].GetMatchingStrings()
			matches5 := r.RM.RuleMap[rm[4]].GetMatchingStrings()
			for _, m1 := range matches1 {
				for _, m2 := range matches2 {
					for _, m3 := range matches3 {
						for _, m4 := range matches4 {
							for _, m5 := range matches5 {
								toAdd := m1 + m2 + m3 + m4 + m5
								if len(toAdd) <= r.RM.MaxLen {
									matchingStringsMap[toAdd] = struct{}{}
								}
							}
						}
					}
				}
			}
		} else if len(rm) == 4 {
			matches1 := r.RM.RuleMap[rm[0]].GetMatchingStrings()
			matches2 := r.RM.RuleMap[rm[1]].GetMatchingStrings()
			matches3 := r.RM.RuleMap[rm[2]].GetMatchingStrings()
			matches4 := r.RM.RuleMap[rm[3]].GetMatchingStrings()
			for _, m1 := range matches1 {
				for _, m2 := range matches2 {
					for _, m3 := range matches3 {
						for _, m4 := range matches4 {
							toAdd := m1 + m2 + m3 + m4
							if len(toAdd) <= r.RM.MaxLen {
								matchingStringsMap[toAdd] = struct{}{}
							}
						}
					}
				}
			}
		} else if len(rm) == 3 {
			matches1 := r.RM.RuleMap[rm[0]].GetMatchingStrings()
			matches2 := r.RM.RuleMap[rm[1]].GetMatchingStrings()
			matches3 := r.RM.RuleMap[rm[2]].GetMatchingStrings()
			for _, m1 := range matches1 {
				for _, m2 := range matches2 {
					for _, m3 := range matches3 {
						toAdd := m1 + m2 + m3
						if len(toAdd) <= r.RM.MaxLen {
							matchingStringsMap[toAdd] = struct{}{}
						}
					}
				}
			}
		} else if len(rm) == 2 {
			matches1 := r.RM.RuleMap[rm[0]].GetMatchingStrings()
			matches2 := r.RM.RuleMap[rm[1]].GetMatchingStrings()
			if len(matches1) == 65792 || len(matches2) == 65792 {
				fmt.Println("THIS")
			}
			for _, m1 := range matches1 {
				for _, m2 := range matches2 {
					toAdd := m1 + m2
					if len(toAdd) <= r.RM.MaxLen {
						matchingStringsMap[toAdd] = struct{}{}
					}
				}
			}
		} else if len(rm) == 1 {
			matchingStrings = append(matchingStrings, r.RM.RuleMap[rm[0]].GetMatchingStrings()...)
		} else {
			fmt.Println("Exception: ", r.RuleString)
			panic("Not like this...")
		}
	}
	for k, _ := range matchingStringsMap {
		matchingStrings = append(matchingStrings, k)
	}
	r.MatchingStrings = matchingStrings
	return matchingStrings
}

func (r *Rule) GetMinMaxLen() {
	if len(r.MatchingStrings) == 0 {
		r.GetMatchingStrings()
	}
	minL := 100
	maxL := 0
	for _, s := range r.MatchingStrings {
		l := len(s)
		if l < minL {
			minL = l
		}
		if l > maxL {
			maxL = l
		}
	}
	r.MaxLen = maxL
	r.MinLen = minL
}

func (rm *RuleMatcher) parseRule(rule string) {
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
}

func NewRuleMatcher(rules []string, maxLen int) *RuleMatcher {
	rm := &RuleMatcher{RuleMap: map[int]*Rule{}, MaxLen: maxLen}
	for _, rule := range rules {
		rm.parseRule(rule)
	}
	return rm
}
