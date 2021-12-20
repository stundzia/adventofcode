package day20

import (
	"fmt"
	"math"
	"strconv"

	"github.com/stundzia/adventofcode/utils"
)

type imgHandler struct {
	algo string
	img map[string]rune
}

var defaultPixelInt = 0

func newImgHandler(algo string, lines []string) *imgHandler {
	ih := &imgHandler{
		algo: algo,
		img: map[string]rune{},
	}
	gg := 0
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				gg++
			}
			ih.img[fmt.Sprintf("%d,%d", x, y)] = char
		}
	}
	return ih
}

func (ih *imgHandler) getPixel(x, y int) rune {
	if pixel, exists := ih.img[fmt.Sprintf("%d,%d", x, y)]; exists {
		return pixel
	}
	return '.'
}

func (ih *imgHandler) getPixelAsInt(x, y int) int {
	if pixel, exists := ih.img[fmt.Sprintf("%d,%d", x, y)]; exists {
		if pixel == '#' {
			return 1
		} else {
			return 0
		}
	}
	return defaultPixelInt
}

func (ih *imgHandler) pixelAlgoIndex(x, y int) int {
	binaryString := ""
	for yL := y - 1; yL <= y + 1; yL++ {
		for xL := x - 1; xL <= x + 1; xL++ {
			binaryString = binaryString + strconv.Itoa(ih.getPixelAsInt(xL, yL))
		}
	}
	biInInt, _ := strconv.ParseInt(binaryString, 2, 64)
	return int(biInInt)
}

func (ih *imgHandler) getEnhancedPixelValue(x, y int) rune {
	algoIndex := ih.pixelAlgoIndex(x, y)
	return rune(ih.algo[algoIndex])
}

func (ih *imgHandler) enhanceImg(sizeUp int) {
	res := map[string]rune{}

	minX, maxX, minY, maxY := ih.getCoordsRanges()
	for x := minX - sizeUp; x < maxX + sizeUp; x++ {
		for y := minY - sizeUp; y < maxY + sizeUp; y++ {
			res[fmt.Sprintf("%d,%d", x,y)] = ih.getEnhancedPixelValue(x, y)
		}
	}
	ih.img = res
}

func (ih *imgHandler) getCoordsRanges() (minX, maxX, minY, maxY int) {
	minX = math.MaxInt
	minY = math.MaxInt
	for coords, _ := range ih.img {
		x, y := utils.CoordsStrToInts(coords)
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	return
}


func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 20, "\n")
	algoLine := lines[0]
	imgLines := lines[2:]
	ih := newImgHandler(algoLine, imgLines)
	ih.enhanceImg(3)
	defaultPixelInt = 1
	ih.enhanceImg(3)
	minX, maxX, minY, maxY := ih.getCoordsRanges()
	lightCount := 0
	for y := minY + 1; y < maxY - 1; y++ {
		for x := minX + 1; x < maxX - 1; x++ {
			if ih.getPixel(x, y) == '#' {
				lightCount++
			}
		}
	}
	return fmt.Sprintf("Solution: %d", lightCount)
}

func DoGold() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2021, 20, "\n")
	algoLine := lines[0]
	imgLines := lines[2:]
	ih := newImgHandler(algoLine, imgLines)
	lightCount := 0
	for i := 0; i < 50; i++ {
		defaultPixelInt = i % 2
		ih.enhanceImg(3)
	}
	minX, maxX, minY, maxY := ih.getCoordsRanges()
	for y := minY + 1; y < maxY - 1; y++ {
		for x := minX + 1; x < maxX - 1; x++ {
			if ih.getPixel(x, y) == '#' {
				lightCount++
			}
		}
	}
	return fmt.Sprintf("Solution: %d", lightCount)
}
