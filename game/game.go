package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type GameUI interface {
	Init()
	Draw(*Level, bool) bool
	SelectLevel() (*Level, bool)
	AddPreview(Level)
}

type Tile rune

const (
	StoneWall     Tile = '#'
	DirtFloor     Tile = '.'
	DoorC         Tile = '|'
	DoorO         Tile = '/'
	MainCharacter Tile = 'P'
	ChestC        Tile = 'C'
	ChestO        Tile = 'c'
	//Monster       Tile = 'm' // will be used as specific monster like rat or snake etc
	Blank   Tile = 0
	Pending Tile = -1
)

type Grid struct {
	Layers     []Tile
	Background Tile
}

type Row struct {
	x, y  int
	Grids []Grid
}

type Pos struct {
	X, Y int
}

type GridWorld struct {
	Rows []Row
}

type Entity struct {
	Pos
	Tile Tile
}

type Character struct {
	Entity
	Hitpoints    int
	Name         string
	Strength     int
	Speed        float64
	ActionPoints float64
}

type Monster struct {
	Character
}

type Player struct {
	Character
}

func (player *Player) Move() {

}

type Level struct {
	GridWorld GridWorld
	LevelName string
	Entities  []Entity
}

type Level2 struct {
	Map      [][]Tile
	Player   *Player
	Monsters map[Pos]*Monster
	//Debug    map[Pos]bool
}

func LoadLevelFromFile2(fileName string) *Level2 {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	levelLines := make([]string, 0)
	longestRow := 0
	index := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levelLines = append(levelLines, scanner.Text())
		if longestRow < len(levelLines[index]) {
			longestRow = len(levelLines[index])
		}
		index++
	}
	defer file.Close()

	level := &Level2{}
	level.Player = &Player{}

	// TODO where should we initialize the player?
	level.Player.ActionPoints = 0
	level.Player.Strength = 5
	level.Player.Hitpoints = 20
	level.Player.Tile = 'R'
	level.Player.Speed = 1.0
	// level.Player.Name = "PurpleWIZARD"

	level.Map = make([][]Tile, len(levelLines))
	//level.Monsters = make(map[Pos]*Monster)

	for i := range level.Map {
		level.Map[i] = make([]Tile, longestRow)
	}

	for y := 0; y < len(level.Map); y++ {
		line := levelLines[y]
		for x, c := range line {

			var t Tile
			switch c {
			case ' ', '\t', '\n', '\r':
				t = Blank
			case '#':
				t = StoneWall
			case '.':
				t = DirtFloor
			case '|':
				t = DoorC
			case '-':
				t = DoorO
			case 'P':
				level.Player.X = x
				level.Player.Y = y
				t = Pending
			// case 'R':
			// 	level.Monsters[Pos{x, y}] = NewRat(Pos{x, y})
			// 	t = Pending
			// case 'S':
			// 	level.Monsters[Pos{x, y}] = NewSpider(Pos{x, y})
			// 	t = Pending
			default:
				panic("Invalid character in the map file")
			}
			level.Map[y][x] = t

		}
	}

	// If tile pending, it assigns the background of that tile. E.g. when player is encountered on the map, it puts a dirt floor under that tile
	for y, row := range level.Map {
		for x, tile := range row {
			if tile == Pending {
				level.Map[y][x] = level.bfsFloor(Pos{x, y})
			}
		}
	}

	return level
}

func (level *Level) ToString() {
	gw := level.GridWorld
	for y := range gw.Rows {
		for x := range gw.Rows[y].Grids {
			fmt.Print(gw.Rows[y].Grids[x].Background.toString())
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
	case DoorC:
		return "|"
	case DoorO:
		return "-"
	case Blank:
		return " "
	case MainCharacter:
		return "P"
	case ChestC:
		return "C"
	case ChestO:
		return "c"
	case Monster:
		return "m"
	default:
		panic("unknown toString tile")
	}
}

func createPreviews(ui GameUI) {
	for levelindex := 1; levelindex <= 4; levelindex++ {
		level := Level{}
		level.LevelName = "level" + strconv.Itoa(levelindex)
		level.loadLevelFromFile()
		ui.AddPreview(level)
	}
}

func Run(ui GameUI) {
	rand.Seed(time.Now().UnixNano())
	isReplayed := true
	for isReplayed {
		ui.Init()
		createPreviews(ui)
		var editBeforeStart bool
		level, editBeforeStart := ui.SelectLevel()
		if level == nil {
			return
		}
		level.loadLevelFromFile()
		isReplayed = ui.Draw(level, editBeforeStart)
	}
}
