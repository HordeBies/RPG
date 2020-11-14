package game

import (
	"bufio"
	"fmt"
	"os"
)

type GameUI interface {
	Draw(*Level)
}

type Tile rune

const (
	StoneWall Tile = '#'
	DirtFloor Tile = '.'
	Door      Tile = '|'
	Blank     Tile = ' '
)

type Level struct {
	Map [][]Tile
}

func loadLevelFromFile(filename string) *Level {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	levelLines := make([]string, 0)
	longestRow := 0
	for scanner.Scan() {
		levelLines = append(levelLines, scanner.Text())
		if len(scanner.Text()) > longestRow {
			longestRow = len(scanner.Text())
		}
	}

	level := &Level{}
	level.Map = make([][]Tile, len(levelLines))
	for i := range level.Map {
		level.Map[i] = make([]Tile, longestRow)
	}

	for y := 0; y < len(level.Map); y++ {
		line := levelLines[y]
		for x, c := range line {
			switch c {
			case ' ', '\t', '\n', '\r':
				level.Map[y][x] = Blank
			case '#':
				level.Map[y][x] = StoneWall
			case '|':
				level.Map[y][x] = Door
			case '.':
				level.Map[y][x] = DirtFloor
			default:
				panic("unknown Mapping")

			}

		}
	}

	for y := 0; y < len(level.Map); y++ {
		for x := 0; x < longestRow; x++ {
			fmt.Print(level.Map[y][x], " ")
		}
		fmt.Println("")
	}

	return nil
}

func Run(ui GameUI) {
	level := loadLevelFromFile("game/maps/level1.map")
	ui.Draw(level)
}
