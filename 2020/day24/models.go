package day24

import "fmt"

type Floor struct {
	RefTile *Tile
	TilesMap map[[2]int]*Tile
	BlackTileCount int
}

type Tile struct {
	Coords [2]int
	OnFloor *Floor
	Black bool
}

func NewFloor() *Floor {
	floor := &Floor{
		TilesMap: map[[2]int]*Tile{},
	}
	floor.RefTile = floor.NewTile([2]int{0,0})
	return floor
}

var CoordsDiffMap = map[string][2]int {
	"e": {2,0},
	"w": {-2,0},
	"ne": {1,1},
	"nw": {-1,1},
	"se": {1,-1},
	"sw": {-1,-1},
}

func (t *Tile) Flip() {
	t.Black = !t.Black
	if t.Black {
		t.OnFloor.BlackTileCount++
	} else {
		t.OnFloor.BlackTileCount--
	}
	fmt.Printf("Flipped: %v\n", t.Coords)
}


func (f *Floor) TraverseTiles(path string) *Tile {
	direction := ""
	currentTile := f.RefTile
	PathLoop:
	for _, d := range path {
		dStr := string(d)
		if direction == "" {
			if dStr == "e" || dStr == "w" {
				diff := CoordsDiffMap[dStr]
				currCoords := currentTile.Coords
				currentTile = f.GetOrCreateTile([2]int{currCoords[0] + diff[0], currCoords[1] + diff[1]})
			} else {
				direction += dStr
				continue PathLoop
			}
		} else {
			dStr = direction + dStr
			diff := CoordsDiffMap[dStr]
			currCoords := currentTile.Coords
			currentTile = f.GetOrCreateTile([2]int{currCoords[0] + diff[0], currCoords[1] + diff[1]})
			direction = ""
			continue PathLoop
		}
	}
	return currentTile
}

func (f *Floor) NewTile(coords [2]int) *Tile {
	tile := &Tile{
		OnFloor: f,
		Coords: coords,
		Black:      false,
	}
	f.TilesMap[coords] = tile
	return tile
}

func (f *Floor) GetOrCreateTile(coords [2]int) *Tile {
	tile, ok := f.TilesMap[coords]
	if !ok {
		tile = f.NewTile(coords)
	}
	return tile
}