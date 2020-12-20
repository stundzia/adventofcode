package day20

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"os"
	"strconv"
	"strings"
)

type Image struct {
	Tiles []*Tile
	Matrix [12][12]*Tile
	FullImageMatrix [][]string
}

type Tile struct {
	Img *Image
	Id int
	Coords [2]int
	Matrix [][]string
	Borders map[string]string
	PairingSlice []int
	PairingMap map[string]*Tile
	Adjacent map[string]*Tile
}

var oppositeSides = map[string]string {
	"top": "bot",
	"bot": "top",
	"left": "right",
	"right": "left",
}

var monsterPositionsFromTail = [][2]int{
	{0,0},
	{1,1},
	{1,4},
	{0,5},
	{0,6},
	{1,7},
	{1,10},
	{0,11},
	{0,12},
	{1,13},
	{1,16},
	{0,17},
	{0,18},
	{-1,18},
	{0,19},
}

func NewImageFromInput(input []string) *Image {
	img := &Image{
		Tiles: []*Tile{},
		Matrix: [12][12]*Tile{},
		FullImageMatrix: [][]string{},
	}
	for _, ts := range input {
		img.NewTileFromStringSlice(strings.Split(ts, "\n"))
	}
	return img
}


func (img *Image) GetCoordsValue(coords []int) string {
	sideLen := len(img.FullImageMatrix)
	if coords[0] >= sideLen || coords[1] >= sideLen {
		return "e"
	}
	return img.FullImageMatrix[coords[0]][coords[1]]
}

func (t *Tile) RemoveBorders() {
	t.Matrix = t.Matrix[1:len(t.Matrix) - 1]
	for i, line := range t.Matrix {
		t.Matrix[i] = line[1:len(line) -1]
	}
}

func (img *Image) MarkIfMonster(tailCoords []int) {
	if img.PixelIsMonsterTail(tailCoords) {
		for _, delta := range monsterPositionsFromTail {
			coords := []int{tailCoords[0] + delta[0], tailCoords[1] + delta[1]}
			img.FullImageMatrix[coords[0]][coords[1]] = "O"
		}
	}
}

func (img *Image) PixelIsMonsterTail(tailCoords []int) bool {
	for _, delta := range monsterPositionsFromTail {
		coords := []int{tailCoords[0] + delta[0], tailCoords[1] + delta[1]}
		if img.GetCoordsValue(coords) != "#" {
			return false
		}
	}
	return true
}


func (img *Image) NewTileFromStringSlice(ss []string) *Tile {
	idStr := ss[0][5:9]
	id, _ := strconv.Atoi(idStr)
	matrix := [][]string{}
	for _, line := range ss[1:] {
		matrix = append(matrix, strings.Split(line, ""))
	}

	tile := &Tile{
		Img: img,
		Coords: [2]int{-1, -1},
		Id:      id,
		Matrix:  matrix,
		Adjacent: map[string]*Tile{},
	}
	img.Tiles = append(img.Tiles, tile)
	return tile
}

func (img *Image) PositionTile(tile *Tile, position [2]int) {
	img.Matrix[position[0]][position[1]] = tile
	tile.Coords = position
}

func (img *Image) PositionAllTiles() {
	for _, line := range img.Matrix {
		for _, tile := range line {
			if tile != nil {
				tile.PositionPairings()
			}
		}
	}
}

func (t *Tile) PositionTileAtSide(tile *Tile,side string) {
	switch side {
	case "left":
		t.Img.PositionTile(tile, [2]int{t.Coords[0], t.Coords[1] - 1})
	case "right":
		t.Img.PositionTile(tile, [2]int{t.Coords[0], t.Coords[1] + 1})
	case "top":
		t.Img.PositionTile(tile, [2]int{t.Coords[0] - 1, t.Coords[1]})
	case "bot":
		t.Img.PositionTile(tile, [2]int{t.Coords[0] + 1, t.Coords[1]})
	default:
		panic("aaaahhhh, wtf?!")
	}
	tile.FindPotentialPairings()
}


func (t *Tile) PositionPairings() {
	for side, otherTile := range t.PairingMap {
		if t.Adjacent[side] != nil {
			continue
		}
		aligned := false
		for i := 0;aligned == false;i++ {
			if t.Borders[side] == otherTile.Borders[oppositeSides[side]] {
				aligned = true
				t.PositionTileAtSide(otherTile, side)
				t.Adjacent[side] = otherTile
				otherTile.Adjacent[oppositeSides[side]] = t
				break
			} else if (side == "left" || side == "right") && t.Borders[side] == utils.ReverseString(otherTile.Borders[oppositeSides[side]]) {
				otherTile.Flip(true)
				otherTile.GetBorders(true)
			} else if (side == "top" || side == "bot") && t.Borders[side] == utils.ReverseString(otherTile.Borders[oppositeSides[side]]) {
				otherTile.Flip(false)
				otherTile.GetBorders(true)
			} else {
				otherTile.RotateClockwise()
				otherTile.GetBorders(true)
			}
			if i > 10 {
				panic("failed")
			}
		}
	}
}

func (t *Tile) PrintMatrix() {
	for _, l := range t.Matrix {
		fmt.Printf("%v\n", l)
	}
}

func (t *Tile) GetBorders(forceRecalc bool) map[string]string {
	if !forceRecalc && t.Borders != nil {
		return t.Borders
	}
	top := strings.Join(t.Matrix[0], "")
	bot := strings.Join(t.Matrix[len(t.Matrix) - 1], "")
	left := ""
	right := ""
	for _, line := range t.Matrix {
		left += line[0]
		right += line[len(line) - 1]
	}
	t.Borders = map[string]string{
		"top": top,
		"bot": bot,
		"left": left,
		"right": right,
	}
	return t.Borders
}


func (t *Tile) Flip(vertically bool) {
	if vertically {
		t.Matrix = utils.ReverseStringSliceSlice(t.Matrix)
	} else {
		for i, line := range t.Matrix {
			t.Matrix[i] = utils.ReverseStringSlice(line)
		}
	}
	t.GetBorders(true)
}

func (t *Tile) RotateClockwise() {
	cols := [][]string{}
	for i := len(t.Matrix) - 1; i >= 0; i-- {
		cols = append(cols, t.Matrix[i])
	}
	newMatrix := [][]string{}
	for i := 0; i < len(t.Matrix); i++ {
		line := []string{}
		for _, col := range cols {
			line = append(line, col[i])
		}
		newMatrix = append(newMatrix, line)
	}
	t.Matrix = newMatrix
	t.GetBorders(true)
}


func (t *Tile) FindPotentialPairings() {
	pairingMap := map[string]*Tile{}
	t.GetBorders(false)
	for _, other := range t.Img.Tiles {
		if other.Id != t.Id {
			other.GetBorders(false)
			for position, border := range t.Borders {
				for _, otherBorder := range other.Borders {
					if border == otherBorder || border == utils.ReverseString(otherBorder) {
						pairingMap[position] = other
					}
				}
			}
		}
	}
	t.PairingMap = pairingMap
}

func (img *Image) FormFullImageMatrix(printIt bool) {
	lines := []string{}
	for _, tileLine := range img.Matrix {
		for y := 0; y < len(tileLine[0].Matrix); y++ {
			l := ""
			for _, tile := range tileLine {
				l += strings.Join(tile.Matrix[y], "")
			}
			lines = append(lines, l)
		}
	}
	for _, line := range lines {
		img.FullImageMatrix = append(img.FullImageMatrix, strings.Split(line, ""))
	}
	if printIt {
		for _, l := range lines {
			fmt.Println(l)
		}
	}
}


func (img *Image) Flip(vertically bool) {
	if vertically {
		img.FullImageMatrix = utils.ReverseStringSliceSlice(img.FullImageMatrix)
		return
	}
	for i, line := range img.FullImageMatrix {
		img.FullImageMatrix[i] = utils.ReverseStringSlice(line)
	}
}


func (img *Image) RotateClockwise() {
	cols := [][]string{}
	for i := len(img.FullImageMatrix) - 1; i >= 0; i-- {
		cols = append(cols, img.FullImageMatrix[i])
	}
	newMatrix := [][]string{}
	for i := 0; i < len(img.FullImageMatrix); i++ {
		line := []string{}
		for _, col := range cols {
			line = append(line, col[i])
		}
		newMatrix = append(newMatrix, line)
	}
	img.FullImageMatrix = newMatrix
}

func (img *Image) GetMonsterAndSeaCount() (monsterCount, seaCount int) {
	pixelCount := 0
	pixelCount2 := 0
	for _, line := range img.FullImageMatrix {
		pixelCount2 += len(line)
		for _, pixel := range line {
			pixelCount++
			if pixel == "#" {
				seaCount++
			}
			if pixel == "O" {
				monsterCount++
			}
		}
	}
	return
}

func (img *Image) DumpToFile(fName string) {
	f, err := os.Create(fName)
	if err != nil {
		panic("well, fuck")
	}
	for _, line := range img.FullImageMatrix {
		_, _ = f.WriteString(fmt.Sprintf("%s\n", strings.Join(line, "")))
	}
}