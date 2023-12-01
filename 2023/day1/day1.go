package day1

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2023, 1, "\n")
	res := 0
	for _, l := range input {
		var first int32
		var last int32
		for _, c := range l {
			if c < 49 || c > 57 {
				continue
			}
			if first == 0 {
				first = c - 48
			}
			last = c - 48
		}
		add, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
		if err != nil {
			log.Fatal(err)
		}
		res += add
	}

	return fmt.Sprintf("Solution: %d", res)
}

func handleLineP2(line string) int {
	first := -1
	last := -1
	from := 0
	for i, c := range line {
		d, err := strconv.Atoi(string(c))
		if err == nil {
			if first == -1 {
				first = d
			}
			last = d
			from = i
			continue
		}
		for k, v := range digitMap {
			if strings.Contains(line[from:i+1], k) {
				if first == -1 {
					first = v
				}
				last = v
				from = i
				continue
			}
		}
	}
	add, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
	if err != nil {
		log.Fatal(err)
	}

	return add
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2023, 1, "\n")
	res := 0
	for _, l := range input {
		res += handleLineP2(l)
	}

	return fmt.Sprintf("Solution: %d", res)
}
