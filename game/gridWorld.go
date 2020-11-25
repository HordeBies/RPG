package game

import (
	"bufio"
	"os"
	"unicode/utf8"
)

func getTile(c rune) (t Tile) {
	switch c {
	case ' ', '\t', '\n', '\r':
		t = Blank
	case '#':
		t = StoneWall
	case '|':
		t = DoorC
	case '-':
		t = DoorO
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

	file, err := os.OpenFile("game/maps/"+level.LevelName+".map", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	for y := range gw.Rows {
		for x := range gw.Rows[y].Grids {
			if gw.Rows[y].Grids[x].Background != Blank {
				file.WriteString(gw.Rows[y].Grids[x].Background.toString())
			} else {
				file.WriteString(" ")
			}
		}
		file.WriteString("\n")
	}
	file.Close()
}

func (level *Level) loadLevelFromFile() {
	filename := level.LevelName
	file, err := os.Open("game/maps/" + filename + ".map")
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
	level.GridWorld = GridWorld{}
	level.GridWorld.Rows = make([]Row, len(levelLines), 100)

	for i, line := range levelLines {
		level.GridWorld.Rows[i].Grids = make([]Grid, utf8.RuneCountInString(line), 100)
		for j, c := range line {
			level.GridWorld.Rows[i].Grids[j].Background = getTile(c)
		}
	}
	/*
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
		}*/
}
