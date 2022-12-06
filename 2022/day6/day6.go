package day6

import (
	"fmt"

	"github.com/stundzia/adventofcode/utils"
)

type buffer struct {
	size   int
	buffer []rune
}

func (b *buffer) add(r rune) {
	if len(b.buffer) < b.size {
		b.buffer = append(b.buffer, r)
		return
	}
	b.buffer = append(b.buffer[1:], r)
}

func (b *buffer) allUnique() (bool, int) {
	for i, r := range b.buffer {
		for t, r2 := range b.buffer {
			if i != t && r == r2 {
				return false, 0
			}
		}
	}
	return true, len(b.buffer)
}

func DoSilver() string {
	line, _ := utils.ReadInputFileContentsAsString(2022, 6)
	var res int
	b := &buffer{
		size:   4,
		buffer: []rune{},
	}
	for i, c := range line {
		b.add(c)
		unique, count := b.allUnique()
		if unique && count == 4 {
			res = i + 1
			break
		}
	}

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	line, _ := utils.ReadInputFileContentsAsString(2022, 6)
	var res int
	b := &buffer{
		size:   14,
		buffer: []rune{},
	}
	for i, c := range line {
		b.add(c)
		unique, count := b.allUnique()
		if unique && count == 14 {
			res = i + 1
			break
		}
	}

	return fmt.Sprintf("Solution: %d", res)
}
