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
	input, _ := utils.ReadInputFileContentsAsIntSlice(2020, 16, "\n")
	return fmt.Sprintf("%d", len(input))
}
