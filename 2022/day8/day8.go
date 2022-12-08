package day8

import (
	"fmt"
	"sync"
	"sync/atomic"

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
	if lefVisible {
		return true
	}
	rightVisible := true
	for x := t.forest.maxX; x > t.X; x-- {
		if t.forest.trees[coordsStr(x, t.Y)].height >= t.height {
			rightVisible = false
			break
		}
	}
	if rightVisible {
		return true
	}
	topVisible := true
	for y := 0; y < t.Y; y++ {
		if t.forest.trees[coordsStr(t.X, y)].height >= t.height {
			topVisible = false
			break
		}
	}
	if topVisible {
		return true
	}
	bottomVisible := true
	for y := t.forest.maxY; y > t.Y; y-- {
		if t.forest.trees[coordsStr(t.X, y)].height >= t.height {
			bottomVisible = false
			break
		}
	}
	if bottomVisible {
		return true
	}

	return false
}

func (t *tree) scenicScore() int64 {
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

	return int64(resXL * resXR * resYU * resYD)
}

func DoSilver() string {
	grid, _ := utils.ReadInputFileContentsAsIntGrid(2022, 8)
	forest := newForest(grid)
	res := &atomic.Uint32{}
	wg := sync.WaitGroup{}
	for _, t := range forest.trees {
		wg.Add(1)
		go func(tr *tree) {
			if tr.isVisible() {
				res.Add(1)
			}
			wg.Done()
		}(t)
	}
	wg.Wait()

	return fmt.Sprintf("Solution: %d", res.Load())
}

func DoGold() string {
	grid, _ := utils.ReadInputFileContentsAsIntGrid(2022, 8)
	forest := newForest(grid)

	res := atomic.Int64{}
	wg := sync.WaitGroup{}
	for _, t := range forest.trees {
		wg.Add(1)
		go func(tr *tree) {
			scenic := tr.scenicScore()
			if scenic > res.Load() {
				res.Store(scenic)
			}
			wg.Done()
		}(t)
	}
	wg.Wait()

	return fmt.Sprintf("Solution: %d", res.Load())
}
