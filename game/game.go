package game

import "fmt"

type GameUI interface {
	Draw(*Level, int)
}

type Tile rune

const (
	StoneWall     Tile = '#'
	DirtFloor     Tile = '.'
	Door          Tile = '|'
	Blank         Tile = 0
	MainCharacter Tile = 'P'
)

type Grid struct {
	Layers     []Tile
	Background Tile
	Entity     []Tile
}

type Row struct {
	x, y  int
	Grids []Grid
}

type GridWorld struct {
	Rows []Row
}

type Entity struct {
	x, y int
	tile Tile
}

type Player struct {
	Entity
}

type Level struct {
	GridWorld GridWorld
	Map       [][]Tile
	Player    Player
	LevelName string
	MaxLayers int
}

func (gw *GridWorld) ToString() {
	for y := range gw.Rows {
		for x := range gw.Rows[y].Grids {
			if len(gw.Rows[y].Grids[x].Layers) > 1 {
				fmt.Print("(")
				for l := range gw.Rows[y].Grids[x].Layers {
					fmt.Print(gw.Rows[y].Grids[x].Layers[l])
				}
				fmt.Print(")")
			} else if len(gw.Rows[y].Grids[x].Layers) == 1 {
				fmt.Print(gw.Rows[y].Grids[x].Layers[0])
			}
		}
		fmt.Println("")
	}
}

func (tile Tile) toString() string {
	switch tile {
	case StoneWall:
		return "#"
	case DirtFloor:
		return "."
	case Door:
		return "|"
	case Blank:
		return " "
	case MainCharacter:
		return "P"
	default:
		panic("unknown toString tile")
	}
}

func Run(ui GameUI) {
	level := &Level{}
	level.MaxLayers = 2
	level.LevelName = "level1"
	level.loadLayersFromFile()
	ui.Draw(level, 2)
	// for _, row := range level.gridWorld.rows {
	// 	for _, grid := range row.grids {
	// 		if len(grid.layers) > 1 {
	// 			fmt.Print("(")
	// 			for _, layer := range grid.layers {
	// 				fmt.Print(layer.toString())
	// 			}
	// 			fmt.Print(")")
	// 		} else {
	// 			fmt.Print(grid.layers[0].toString())
	// 		}
	// 	}
	// 	fmt.Println("")
	// }
}
