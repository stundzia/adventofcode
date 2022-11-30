package day8

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"math"
)

type Image struct {
	raw    []int
	layers []*Layer
	xSize  int
	ySize  int
}

type Layer struct {
	pixels [][]int
}

func NewImage(input []int, xSize, ySize int) *Image {
	img := &Image{
		raw:    input,
		xSize:  xSize,
		ySize:  ySize,
		layers: []*Layer{},
	}
	for i := 0; i+(xSize*ySize) <= len(input); i += (xSize * ySize) {
		img.layers = append(img.layers, NewLayer(input[i:i+(xSize*ySize)], xSize, ySize))
	}
	return img
}

func NewLayer(pixels []int, xSize, ySize int) *Layer {
	l := &Layer{pixels: [][]int{}}
	for y := 0; y < ySize; y++ {
		l.pixels = append(l.pixels, pixels[xSize*y:xSize*(y+1)])
	}
	return l
}

func (img *Image) PrintSelf() {
	for i, l := range img.layers {
		fmt.Println("Layer ", i+1)
		for _, r := range l.pixels {
			fmt.Println(r)
		}
	}
}

func (img *Image) PrintVisible() {
	res := make([][]string, img.ySize)
	for i, _ := range res {
		res[i] = make([]string, img.xSize)
	}
	for y, _ := range res {
		for x, _ := range res[y] {
			res[y][x] = ""
		}
	}
	for _, l := range img.layers {
		for y, r := range l.pixels {
			for x, p := range r {
				v := res[y][x]
				if v != "" || p == 2 {
					continue
				}
				v = " "
				if p == 1 {
					v = "X"
				}
				res[y][x] = v
			}
		}
	}
	for _, r := range res {
		fmt.Println(r)
	}
}

func (l *Layer) DigitCount(digit int) int {
	count := 0
	for _, r := range l.pixels {
		for _, n := range r {
			if n == digit {
				count++
			}
		}
	}
	return count
}

func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2019, 8, "")
	img := NewImage(input, 25, 6)
	minZeroCount := math.MaxInt
	lowestLayer := img.layers[0]
	for _, l := range img.layers {
		zeroCount := l.DigitCount(0)
		if zeroCount < minZeroCount {
			minZeroCount = zeroCount
			lowestLayer = l
		}
	}
	return fmt.Sprintf("%d", lowestLayer.DigitCount(1)*lowestLayer.DigitCount(2))
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsIntSlice(2019, 8, "")
	img := NewImage(input, 25, 6)
	img.PrintVisible()

	return fmt.Sprintf("Up there, bud ^^")
}
