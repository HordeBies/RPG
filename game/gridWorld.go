package game

import (
	"bufio"
	"os"
	"strconv"
	"unicode/utf8"
)

func getTile(c rune) (t Tile) {
	switch c {
	case ' ', '\t', '\n', '\r':
		t = Blank
	case '#':
		t = StoneWall
	case '|':
		t = Door
	case '.':
		t = DirtFloor
	case 'P':
		t = MainCharacter
	default:
		panic("unknown gridworld mapping")
	}
	return t
}

func Save(level *Level) {
	gw := level.GridWorld

	for currLayer := 0; currLayer < level.MaxLayers; currLayer++ {
		file, err := os.OpenFile("game/maps/"+level.LevelName+"Layer"+strconv.Itoa(currLayer+1)+".map", os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		for y := range gw.Rows {
			for x := range gw.Rows[y].Grids {
				if gw.Rows[y].Grids[x].Layers != nil && len(gw.Rows[y].Grids[x].Layers) > currLayer {
					file.WriteString(gw.Rows[y].Grids[x].Layers[currLayer].toString())
				} else {
					file.WriteString(" ")
				}
			}
			file.WriteString("\n")
		}
		file.Close()
	}
}

func (level *Level) loadLayersFromFile() {
	filename := level.LevelName
	layerCount := level.MaxLayers
	file, err := os.Open("game/maps/" + filename + "Layer1.map")
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
	level.GridWorld.Rows = make([]Row, len(levelLines), 100)

	for i, line := range levelLines {
		level.GridWorld.Rows[i].Grids = make([]Grid, utf8.RuneCountInString(line), 100)
		for j, c := range line {
			level.GridWorld.Rows[i].Grids[j].Layers = append(level.GridWorld.Rows[i].Grids[j].Layers, getTile(c))
		}
	}

	for i := 2; i <= layerCount; i++ {
		layerfile, err := os.Open("game/maps/" + filename + "Layer" + strconv.Itoa(i) + ".map")
		if err != nil {
			panic(err)
		}
		layerscanner := bufio.NewScanner(layerfile)
		layerLines := make([]string, 0)
		for layerscanner.Scan() {
			layerLines = append(layerLines, layerscanner.Text())
		}
		for i, line := range layerLines {
			for j, c := range line {
				currTile := getTile(c)
				if currTile != Blank {
					level.GridWorld.Rows[i].Grids[j].Layers = append(level.GridWorld.Rows[i].Grids[j].Layers, currTile)
				}
			}
		}
	}
}
