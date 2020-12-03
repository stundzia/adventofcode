package day3

import (
	"fmt"
	"strings"
)

type Mountain struct {
	Grid map[string]string
	TreesEncountered int
	CurrentPosition []int
}


func (m *Mountain) ParseMap(mapSlice []string) {
	grid := make(map[string]string)
	for y, level := range mapSlice {
		for x, step := range strings.Split(level, "") {
			grid[fmt.Sprintf("%d-%d", x,y)] = step
		}
	}
	m.Grid = grid
}


func (m *Mountain) GoDownOneLevel(vectorX int, vectorY int) bool {
	newX := m.CurrentPosition[0] + vectorX
	newY := m.CurrentPosition[1] + vectorY
	m.CurrentPosition[0] = newX
	m.CurrentPosition[1] = newY
	val, ok := m.Grid[fmt.Sprintf("%d-%d", newX, newY)]
	if !ok {
		fmt.Println("Trees encountered: ", m.TreesEncountered)
		return false
	}
	if val == "#" {
		m.TreesEncountered++
	}
	return true
}