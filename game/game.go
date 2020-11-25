package game

import (
	"fmt"
	"strconv"
)

type GameUI interface {
	Init()
	Draw(*Level, int)
	SelectLevel() *Level
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
	Entities   []Tile
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

type PlayerE struct {
	Entity
}

type DoorE struct {
	Entity
	is_open bool
}

type Level struct {
	GridWorld GridWorld
	Map       [][]Tile
	Player    PlayerE
	LevelName string
	MaxLayers int
}

func (gw *GridWorld) ToString() {
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
	for levelindex := 1; levelindex <= 2; levelindex++ {
		level := Level{}
		level.LevelName = "level" + strconv.Itoa(levelindex)
		level.loadLevelFromFile()
		ui.AddPreview(level)
	}
}

func Run(ui GameUI) {
	ui.Init()
	createPreviews(ui)
	level := ui.SelectLevel()
	if level == nil {
		return
	}
	level.loadLevelFromFile()
	ui.Draw(level, 2)
}
