package day4

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

type card struct {
	id                 int
	winningCardNumbers []int
	cardNumbers        []int
	points             int
	matchingNums       int
}

func parseCard(n int, line string) card {
	line = strings.Replace(line, "  ", " ", -1)
	thisCard := card{
		id:                 n,
		winningCardNumbers: []int{},
		cardNumbers:        []int{},
		points:             0,
	}
	numbers := strings.Split(line, ": ")[1]
	numbersParts := strings.Split(numbers, " | ")
	winningParts := strings.Split(numbersParts[0], " ")
	cardParts := strings.Split(numbersParts[1], " ")
	for _, w := range winningParts {
		nw, err := strconv.Atoi(w)
		if err != nil {
			log.Fatal("err: ", err)
		}
		thisCard.winningCardNumbers = append(thisCard.winningCardNumbers, nw)
	}

	for _, c := range cardParts {
		nc, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal("err: ", err)
		}
		thisCard.cardNumbers = append(thisCard.cardNumbers, nc)
	}
main:
	for _, num := range thisCard.cardNumbers {
		for _, winNum := range thisCard.winningCardNumbers {
			if num == winNum {
				thisCard.matchingNums++
				if thisCard.points == 0 {
					thisCard.points = 1
				} else {
					thisCard.points *= 2
				}
				continue main
			}
		}
	}

	return thisCard
}

func (c card) cardsToCopy() []int {
	var res []int
	if c.matchingNums == 0 {
		return res
	}
	for i := 0; i < c.matchingNums; i++ {
		res = append(res, c.id+i+1)
	}

	return res
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2023, 4, "\n")
	cards := map[int]card{}

	res := 0
	for i, line := range input {
		cards[i+1] = parseCard(i+1, line)
		c := parseCard(i+1, line)
		res += c.points
	}

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2023, 4, "\n")
	cards := []card{}
	for i, line := range input {
		c := parseCard(i+1, line)
		cards = append(cards, c)
	}

	for i := 0; i < len(cards); i++ {
		if ccs := cards[i].cardsToCopy(); ccs != nil {
			for _, cc := range ccs {
				cards = append(cards, cards[cc-1])
			}
		}
	}

	return fmt.Sprintf("Solution: %d", len(cards))
}
