package game

import (
	"fmt"
	"strconv"
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
	DoorO         Tile = '-'
	Blank         Tile = 0
	MainCharacter Tile = 'P'
)

type Grid struct {
	Layers     []Tile
	Background Tile
}

type Row struct {
	x, y  int
	Grids []Grid
}

type GridWorld struct {
	Rows []Row
}

type Entity struct {
	X, Y int
	Tile Tile
}

type Level struct {
	GridWorld GridWorld
	LevelName string
	Entities  []Entity
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
