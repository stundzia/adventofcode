package day7

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type game struct {
	hands []hand
}

type hand struct {
	str          string
	cards        []rune
	bid          int
	cardCountMap map[rune]int
	handType     int
}

var mapStupid = map[string]string{
	"A": "Z",
	"K": "Y",
	"Q": "X",
	"J": "W",
	"T": "V",
}

var mapStupidPart2 = map[string]string{
	"A": "Z",
	"K": "Y",
	"Q": "X",
	"J": "1",
	"T": "V",
}

func parseHand(line string) hand {
	parts := strings.Split(line, " ")
	cards := parts[0]
	for k, v := range mapStupid {
		cards = strings.Replace(cards, k, v, -1)
	}
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("wtf")
	}
	h := hand{
		str:          line,
		cards:        []rune(cards),
		bid:          bid,
		cardCountMap: map[rune]int{},
	}
	h.handType = calcHandType(cards)
	return h
}

func parseHandV2(line string) hand {
	parts := strings.Split(line, " ")
	cards := parts[0]
	for k, v := range mapStupidPart2 {
		cards = strings.Replace(cards, k, v, -1)
	}
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("wtf")
	}
	h := hand{
		str:          line,
		cards:        []rune(cards),
		bid:          bid,
		cardCountMap: map[rune]int{},
	}

	h.handType = calcHandType(cards)
	cardCountMap := map[rune]int{}
	cardsSlice := []rune{}
	for _, c := range cards {
		if _, f := cardCountMap[c]; !f {
			cardsSlice = append(cardsSlice, c)
			cardCountMap[c] = 0
		}
		cardCountMap[c]++
	}
	if _, f := cardCountMap['1']; !f {
		return h
	}

	maxCard := '_'
	maxCount := 0
	for c, v := range cardCountMap {
		if v > maxCount && c != '1' {
			maxCard = c
			maxCount = v
		}
	}

	if maxCount != 0 && maxCard != '_' {
		newCards := strings.Replace(cards, string('1'), string(maxCard), cardCountMap['1'])
		newHandType := calcHandType(newCards)
		if newHandType > h.handType {
			h.handType = newHandType
		}
	}

	return h
}

func calcHandType(cards string) int {
	cardCountMap := map[rune]int{}
	for _, c := range cards {
		if _, f := cardCountMap[c]; !f {
			cardCountMap[c] = 0
		}
		cardCountMap[c]++
	}
	switch len(cardCountMap) {
	case 1:
		return FiveOfAKind
	case 2:
		counts := []int{}
		for _, v := range cardCountMap {
			counts = append(counts, v)
			if v == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	case 3:
		for _, v := range cardCountMap {
			if v == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	case 4:
		return OnePair
	default:
		return HighCard
	}
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2023, 7, "\n")

	g := game{hands: []hand{}}
	for _, l := range input {
		g.hands = append(g.hands, parseHand(l))
	}

	sort.Slice(g.hands, func(i, j int) bool {
		if g.hands[i].handType == g.hands[j].handType {
			for it := 0; it < 5; it++ {
				if g.hands[i].cards[it] == g.hands[j].cards[it] {
					continue
				}
				return g.hands[i].cards[it] < g.hands[j].cards[it]
			}
		}
		return g.hands[i].handType < g.hands[j].handType
	})

	res := 0

	for i, h := range g.hands {
		res += (i + 1) * h.bid
	}

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2023, 7, "\n")

	g := game{hands: []hand{}}
	for _, l := range input {
		g.hands = append(g.hands, parseHandV2(l))
	}

	sort.Slice(g.hands, func(i, j int) bool {
		if g.hands[i].handType == g.hands[j].handType {
			for it := 0; it < 5; it++ {
				if g.hands[i].cards[it] == g.hands[j].cards[it] {
					continue
				}
				return g.hands[i].cards[it] < g.hands[j].cards[it]
			}
		}
		return g.hands[i].handType < g.hands[j].handType
	})

	res := 0

	for i, h := range g.hands {
		res += (i + 1) * h.bid
	}

	return fmt.Sprintf("Solution: %d", res)
}
