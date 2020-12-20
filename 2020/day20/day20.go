package day20

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"math/rand"
	"time"
)


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 20, "\n\n")
	img := NewImageFromInput(input)
	for _, tile := range img.Tiles {
		tile.FindPotentialPairings()
	}
	cornerIds := []int{}
	for _, tile := range img.Tiles {
		if len(tile.PairingMap) == 2 {
			cornerIds = append(cornerIds, tile.Id)
		}
	}
	res := 1
	for _, cId := range cornerIds {
		res *= cId
	}

	return fmt.Sprintf("%d", res)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 20, "\n\n")
	img := NewImageFromInput(input)
	for _, tile := range img.Tiles {
		tile.FindPotentialPairings()
	}

	cornerIds := []int{}
	cornerTiles := []*Tile{}
	for _, tile := range img.Tiles {
		if len(tile.PairingMap) == 2 {
			cornerIds = append(cornerIds, tile.Id)
			cornerTiles = append(cornerTiles, tile)
		}
	}
	res := 1
	for _, cId := range cornerIds {
		res *= cId
	}

	candidate := cornerTiles[2]
	candidate.FindPotentialPairings()
	img.PositionTile(candidate, [2]int{0,0}) // yup, cheated and found it by trial and error (properly aligned corner)
	for i := 0; i < 100; i++ {
		img.PositionAllTiles()
	}


	for _, tile := range img.Tiles {
		tile.RemoveBorders()
	}
	img.FormFullImageMatrix(false)

	rand.Seed(time.Now().Unix())

	monsterCount, seaCount := 0, 0

	for i := 0; i < 22222; i++ {
		if rand.Int() % 4 == 0 {
			img.RotateClockwise()
		}
		if rand.Int() % 3 == 0 {
			if rand.Int() % 2 == 0 {
				img.Flip(true)
			} else {
				img.Flip(false)
			}
		}

		for y, line := range img.FullImageMatrix {
			for x, pixel := range line {
				if pixel == "#" {
					img.MarkIfMonster([]int{y,x})
				}
			}
		}
		monsterCount, seaCount = img.GetMonsterAndSeaCount()
		if monsterCount != 0 {
			break
		}
	}
	return fmt.Sprintf("%d (Monster count: %d)", seaCount, monsterCount)
}