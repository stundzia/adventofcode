package day16

import (
	"github.com/stundzia/adventofcode/utils"
	"strconv"
	"strings"
)

type TicketValidator struct {
	MyTicket []int
	Rules map[string][2][2]int
	NearbyTickets [][]int
	CompletelyInvalidCount int
	ErrorRate int
}

func parseRangeString(rs string) [2]int {
	rangeParts := strings.Split(rs, "-")
	low, _ := strconv.Atoi(rangeParts[0])
	high, _ := strconv.Atoi(rangeParts[1])
	return [2]int{low,high}
}


func parseRuleLine(rl string) (rule string, ranges [2][2]int) {
	rlParts := strings.Split(rl, ": ")
	rule = strings.Replace(rlParts[0], " ", "_", -1)
	rangeStrs := strings.Split(rlParts[1], " or ")
	ranges[0] = parseRangeString(rangeStrs[0])
	ranges[1] = parseRangeString(rangeStrs[1])
	return rule, ranges
}

func NewTicketValidator(myTicket []int, rules []string, nearbyTickets [][]int) *TicketValidator {
	tv := &TicketValidator{
		MyTicket:      myTicket,
		Rules:         map[string][2][2]int{},
		NearbyTickets: nearbyTickets,
		CompletelyInvalidCount: 0,
		ErrorRate: 0,
	}
	for _, rule := range rules {
		key, value := parseRuleLine(rule)
		tv.Rules[key] = value
	}
	tv.getErrorMetrics()
	return tv
}

func (tv *TicketValidator) fieldValid(field string, value int) bool {
	ruleRanges := tv.Rules[field]
	for _, rl := range ruleRanges {
		if value >= rl[0] && value <= rl[1] {
			return true
		}
	}
	return false
}

func (tv *TicketValidator) valueValidForAnyRule(value int) (bool, int) {
	for field, _ := range tv.Rules {
		if tv.fieldValid(field, value) {
			return true, 0
		}
	}
	return false, value
}

func (tv *TicketValidator) hasCompletelyInvalidFields(ticket []int) bool {
	valueLoop:
	for _, val := range ticket {
		for field, _ := range tv.Rules {
			if tv.fieldValid(field, val) {
				continue valueLoop
			}
		}
		return true
	}
	return false
}

func (tv *TicketValidator) getErrorMetrics() {
	for _, ticket := range tv.NearbyTickets {
		if tv.hasCompletelyInvalidFields(ticket) {
			tv.CompletelyInvalidCount++
		}
		for _, value := range ticket {
			_, errRate := tv.valueValidForAnyRule(value)
			tv.ErrorRate += errRate
		}
	}
}


func (tv *TicketValidator) discardInvalidTicket() bool {
	for i, ticket := range tv.NearbyTickets {
		if tv.hasCompletelyInvalidFields(ticket) {
			tv.NearbyTickets = utils.RemoveFrom2DIntSlice(tv.NearbyTickets, i)
			return true
		}
	}
	return false
}

func (tv *TicketValidator) discardInvalidTickets() {
	discarded := true
	for ;discarded == true; {
		discarded = tv.discardInvalidTicket()
	}
}

func (tv *TicketValidator) getFieldCandidates(fieldIndex int, ignore []string) (candidates []string) {
	validityCountMap := map[string]int{}
	for rule, _ := range tv.Rules {
		validityCountMap[rule] = 0
	}
	for _, ticket := range tv.NearbyTickets {
		for	rule, _ := range tv.Rules {
			if tv.fieldValid(rule, ticket[fieldIndex]) {
				validityCountMap[rule]++
			}
		}
	}
	for field, val := range validityCountMap {
		if val == len(tv.NearbyTickets) && !utils.StringSliceContains(ignore, field) {
			candidates = append(candidates, field)
		}
	}
	return candidates
}