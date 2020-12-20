package day20

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strconv"
	"strings"
)

type Image struct {
	Tiles []*Tile
	Matrix [12][12]*Tile
}

type Tile struct {
	Img *Image
	Id int
	Matrix [][]string
	Borders map[string]string
	PairingSlice []int
	PairingMap map[string]*Tile
}

func NewImageFromInput(input []string) *Image {
	img := &Image{
		Tiles: []*Tile{},
		Matrix: [12][12]*Tile{},
	}
	for _, ts := range input {
		img.NewTileFromStringSlice(strings.Split(ts, "\n"))
	}
	return img
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
		Id:      id,
		Matrix:  matrix,
	}
	img.Tiles = append(img.Tiles, tile)
	return tile
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
		return
	}
	for i, line := range t.Matrix {
		t.Matrix[i] = utils.ReverseStringSlice(line)
	}
}

func (t *Tile) RotateClockwise() {
	cols := [][]string{}
	for i := len(t.Matrix) - 1; i >= 0; i-- {
		cols = append(cols, t.Matrix[i])
	}
	newMatrx := [][]string{}
	for i := 0; i < len(t.Matrix); i++ {
		line := []string{}
		for _, col := range cols {
			line = append(line, col[i])
		}
		newMatrx = append(newMatrx, line)
	}
	t.Matrix = newMatrx
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