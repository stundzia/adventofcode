package day8

import (
	"fmt"

	"github.com/stundzia/adventofcode/utils"
)

type forest struct {
	trees map[string]*tree
	maxX  int
	maxY  int
}

type tree struct {
	forest *forest
	height int
	X      int
	Y      int
}

func newForest(grid [][]int) *forest {
	f := &forest{trees: map[string]*tree{}, maxY: len(grid) - 1, maxX: len(grid[0]) - 1}
	for y, row := range grid {
		for x, height := range row {
			f.trees[fmt.Sprintf("%d_%d", x, y)] = &tree{
				forest: f,
				height: height,
				X:      x,
				Y:      y,
			}
		}
	}
	return f
}

func coordsStr(x, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}

func (t *tree) isVisible() bool {
	lefVisible := true
	for x := 0; x < t.X; x++ {
		if t.forest.trees[coordsStr(x, t.Y)].height >= t.height {
			lefVisible = false
			break
		}
	}
	rightVisible := true
	for x := t.forest.maxX; x > t.X; x-- {
		if t.forest.trees[coordsStr(x, t.Y)].height >= t.height {
			rightVisible = false
			break
		}
	}
	topVisible := true
	for y := 0; y < t.Y; y++ {
		if t.forest.trees[coordsStr(t.X, y)].height >= t.height {
			topVisible = false
			break
		}
	}
	bottomVisible := true
	for y := t.forest.maxY; y > t.Y; y-- {
		if t.forest.trees[coordsStr(t.X, y)].height >= t.height {
			bottomVisible = false
			break
		}
	}

	return rightVisible || lefVisible || bottomVisible || topVisible
}

func (t *tree) scenicScore() int {
	resXR := 0
	for x := t.X + 1; x <= t.forest.maxX; x++ {
		resXR++
		if t.forest.trees[coordsStr(x, t.Y)].height >= t.height {
			break
		}
	}

	resXL := 0
	for x := t.X - 1; x >= 0; x-- {
		resXL++
		if t.forest.trees[coordsStr(x, t.Y)].height >= t.height {
			break
		}
	}

	resYD := 0
	for y := t.Y + 1; y <= t.forest.maxY; y++ {
		resYD++
		if t.forest.trees[coordsStr(t.X, y)].height >= t.height {
			break
		}
	}

	resYU := 0
	for y := t.Y - 1; y >= 0; y-- {
		resYU++
		if t.forest.trees[coordsStr(t.X, y)].height >= t.height {
			break
		}
	}

	return resXL * resXR * resYU * resYD
}

func DoSilver() string {
	grid, _ := utils.ReadInputFileContentsAsIntGrid(2022, 8)
	forest := newForest(grid)
	res := 0
	for _, t := range forest.trees {
		if t.isVisible() {
			res++
		}
	}

	return fmt.Sprintf("Solution: %d", res)
}

func DoGold() string {
	grid, _ := utils.ReadInputFileContentsAsIntGrid(2022, 8)
	forest := newForest(grid)

	res := 0
	for _, t := range forest.trees {
		scenic := t.scenicScore()
		if scenic > res {
			res = scenic
		}
	}

	return fmt.Sprintf("Solution: %d", res)
}
