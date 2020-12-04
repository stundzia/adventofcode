package day3

import (
	"fmt"
	"strings"
)

type Mountain struct {
	Grid             map[string]string
	TreesEncountered int
	CurrentPosition  []int
	MaxX             int
}

func (m *Mountain) ParseMap(mapSlice []string) {
	grid := make(map[string]string)
	for y, level := range mapSlice {
		for x, step := range strings.Split(level, "") {
			grid[fmt.Sprintf("%d-%d", x, y)] = step
		}
	}
	m.Grid = grid
}

func GetMountainFromMap(mapSlice []string) *Mountain {
	// Assuming width (x) is the same on every level
	m := &Mountain{
		TreesEncountered: 0,
		CurrentPosition:  []int{0, 0},
		MaxX:             len(mapSlice[0]) - 1,
	}
	m.ParseMap(mapSlice)
	return m
}

func (m *Mountain) GoDownOneLevel(vectorX int, vectorY int) bool {
	newX := m.CurrentPosition[0] + vectorX
	newY := m.CurrentPosition[1] + vectorY
	if newX > m.MaxX {
		newX = newX - m.MaxX - 1
	}
	m.CurrentPosition[0] = newX
	m.CurrentPosition[1] = newY
	val, ok := m.Grid[fmt.Sprintf("%d-%d", newX, newY)]
	if !ok {
		fmt.Println("Reached bottom, trees encountered: ", m.TreesEncountered)
		return false
	}
	if val == "#" {
		m.TreesEncountered++
	}
	return true
}

func (m *Mountain) Reset() {
	m.TreesEncountered = 0
	m.CurrentPosition = []int{0, 0}
}

func (m *Mountain) GoDownToBottom(vectorX int, vectorY int) {
	stillOnMountain := true
	for stillOnMountain == true {
		stillOnMountain = m.GoDownOneLevel(vectorX, vectorY)
	}
}
