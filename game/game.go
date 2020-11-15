package game

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
	Layers []Tile
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
	level.loadLayersFromFile("level1", 2)
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
