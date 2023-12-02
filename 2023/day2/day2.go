package day2

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"log"
	"strconv"
	"strings"
)

type game struct {
	id       int
	pulls    []pull
	maxRed   int
	maxGreen int
	maxBlue  int
}

type pull struct {
	red   int
	green int
	blue  int
}

func (g game) power() int {
	return g.maxBlue * g.maxRed * g.maxGreen
}

func parseGame(line string) game {
	g := game{
		id:    0,
		pulls: []pull{},
	}
	parts1 := strings.Split(line, ": ")
	parts2 := strings.Split(parts1[0], " ")
	g.id, _ = strconv.Atoi(parts2[1])
	pullsParts := strings.Split(parts1[1], "; ")
	for _, p := range pullsParts {
		pl := pull{
			red:   0,
			green: 0,
			blue:  0,
		}
		countsStrs := strings.Split(p, ", ")
		for _, c := range countsStrs {
			kv := strings.Split(c, " ")
			count, err := strconv.Atoi(kv[0])
			color := kv[1]
			if err != nil {
				log.Fatal("err: ", err, "(", c, ")")
			}
			switch color {
			case "red":
				pl.red = count
				if count > g.maxRed {
					g.maxRed = count
				}
			case "green":
				pl.green = count
				if count > g.maxGreen {
					g.maxGreen = count
				}
			case "blue":
				pl.blue = count
				if count > g.maxBlue {
					g.maxBlue = count
				}
			}
		}
		g.pulls = append(g.pulls, pl)
	}

	return g
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2023, 2, "\n")
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	games := map[int]game{}
	for _, l := range input {
		g := parseGame(l)
		games[g.id] = g
	}

	res := 0

	for id, g := range games {
		if g.maxRed > maxRed || g.maxGreen > maxGreen || g.maxBlue > maxBlue {
			continue
		}

		res += id
	}

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2023, 2, "\n")
	games := map[int]game{}
	res := 0
	for _, l := range input {
		g := parseGame(l)
		games[g.id] = g
		res += g.power()
	}

	return fmt.Sprintf("Solution: %d", res)
}
