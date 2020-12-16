package day16

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strings"
)


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 16, "\n\n")
	rules := strings.Split(input[0], "\n")
	myTicket := utils.SliceStringToInt(strings.Split(strings.Split(input[1], "\n")[1], ","))
	nearbyTicketsStrings := strings.Split(input[2], "\n")[1:]
	nearbyTickets := make([][]int, len(nearbyTicketsStrings))
	for i, nt := range nearbyTicketsStrings {
		nearbyTickets[i] = utils.SliceStringToInt(strings.Split(nt, ","))
	}
	tv := NewTicketValidator(myTicket, rules, nearbyTickets)
	return fmt.Sprintf("%d", tv.ErrorRate)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 16, "\n\n")
	rules := strings.Split(input[0], "\n")
	myTicket := utils.SliceStringToInt(strings.Split(strings.Split(input[1], "\n")[1], ","))
	nearbyTicketsStrings := strings.Split(input[2], "\n")[1:]
	nearbyTickets := make([][]int, len(nearbyTicketsStrings))
	for i, nt := range nearbyTicketsStrings {
		nearbyTickets[i] = utils.SliceStringToInt(strings.Split(nt, ","))
	}
	tv := NewTicketValidator(myTicket, rules, nearbyTickets)
	tv.discardInvalidTickets()
	fieldMap := map[int]string{}
	assigned := []string{}
	for t := 0; t < 13; t++ {
		for i, _ := range myTicket {
			candidates := tv.getFieldCandidates(i, assigned)
			if len(candidates) == 1 {
				fieldMap[i] = candidates[0]
				assigned = append(assigned, candidates[0])
			}
		}
	}
	fmt.Println(fieldMap)
	res := 1
	for index, field := range fieldMap {
		if len(field) >= 9 && field[:9] == "departure" {
			res *= tv.MyTicket[index]
		}
	}
	return fmt.Sprintf("%d", res)
}
