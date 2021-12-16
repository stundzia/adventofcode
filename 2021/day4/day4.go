package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stundzia/adventofcode/utils"
)

type card struct {
	rows [][]int
	cols [][]int
}

func linesToCard(lines []string) *card {
	rowsInt := [][]int{}
	for _, line := range lines {
		line = strings.Replace(line, "  ", " ", -1)
		line = strings.TrimSpace(line)
		numsStr := strings.Split(line, " ")
		numsInt := []int{}
		for _, n := range numsStr {
			nInt, _ := strconv.Atoi(n)
			numsInt = append(numsInt, nInt)
		}
		rowsInt = append(rowsInt, numsInt)
	}
	c := &card{
		rows: rowsInt,
		cols: [][]int{},
	}

	for i := 0; i < 5; i++ {
		col := []int{}
		for t := 0; t < 5; t++ {
			col = append(col, c.rows[t][i])
		}
		c.cols = append(c.cols, col)
	}

	return c
}

func (c *card) markNum(num int) {
	for t, row := range c.rows {
		for i, n := range row {
			if n == num {
				c.rows[t][i] = 0
				c.cols[i][t] = 0
			}
		}
	}
}

func (c *card) hasWon() bool {
	for _, row := range c.rows {
		if utils.SumIntSlice(row) == 0 {
			return true
		}
	}
	for _, col := range c.cols {
		if utils.SumIntSlice(col) == 0 {
			return true
		}
	}
	return false
}

func (c *card) getScore() int {
	score := 0
	for _, row := range c.rows {
		score += utils.SumIntSlice(row)
	}
	return score
}

func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 4, "\n")
	winNumsStr := lines[0]
	winNumsInt := []int{}
	for _, num := range strings.Split(winNumsStr, ",") {
		n, _ := strconv.Atoi(num)
		winNumsInt = append(winNumsInt, n)
	}
	cardLines := []string{}
	cards := []*card{}
	for i, line := range lines {
		if i == 0 {
			continue
		}
		if line == "" {
			if len(cardLines) == 5 {
				cards = append(cards, linesToCard(cardLines))
			}
			cardLines = []string{}

		} else {
			cardLines = append(cardLines, line)
		}
	}
	for _, num := range winNumsInt {
		for _, c := range cards {
			c.markNum(num)
			if c.hasWon() {
				return fmt.Sprintf("Solution: %d", c.getScore()*num)
			}
		}
	}
	return fmt.Sprintf("Failed to find solution")
}

func DoGold() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 4, "\n")
	winNumsStr := lines[0]
	winNumsInt := []int{}
	for _, num := range strings.Split(winNumsStr, ",") {
		n, _ := strconv.Atoi(num)
		winNumsInt = append(winNumsInt, n)
	}
	cardLines := []string{}
	cards := map[int]*card{}
	for i, line := range lines {
		if i == 0 {
			continue
		}
		if line == "" {
			if len(cardLines) == 5 {
				cards[i] = linesToCard(cardLines)
			}
			cardLines = []string{}

		} else {
			cardLines = append(cardLines, line)
		}
	}
	for _, num := range winNumsInt {
		for k, c := range cards {
			c.markNum(num)
			if c.hasWon() {
				if len(cards) == 1 {
					return fmt.Sprintf("Solution: %d", c.getScore()*num)
				} else {
					delete(cards, k)
				}
			}
		}
	}
	return fmt.Sprintf("Failed to find solution")
}
