package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func inRange(level *Level2, pos Pos) bool {
	return pos.X < len(level.Map[0]) && pos.Y < len(level.Map) && pos.X >= 0 && pos.Y >= 0
}

func checkDoor(level *Level2, pos Pos) {
	t := level.Map[pos.Y][pos.X]
	fmt.Println(t)

	if t == DoorC {
		level.Map[pos.Y][pos.X] = DoorO
	}
}

func HandleInput(currKeyState, prevKeyState []uint8, level *Level2) {
	p := level.Player

	if currKeyState[sdl.SCANCODE_UP] != 0 && prevKeyState[sdl.SCANCODE_UP] == 0 {
		newPos := Pos{p.X, p.Y - 1}
		if canWalk(level, newPos) {
			level.Player.Pos = newPos
		} else {
			checkDoor(level, newPos)
		}
	}

	if currKeyState[sdl.SCANCODE_DOWN] != 0 && prevKeyState[sdl.SCANCODE_DOWN] == 0 {
		newPos := Pos{p.X, p.Y + 1}
		if canWalk(level, newPos) {
			level.Player.Pos = newPos
		} else {
			checkDoor(level, newPos)
		}
	}

	if currKeyState[sdl.SCANCODE_LEFT] != 0 && prevKeyState[sdl.SCANCODE_LEFT] == 0 {
		newPos := Pos{p.X - 1, p.Y}
		if canWalk(level, newPos) {
			level.Player.Pos = newPos
		} else {
			checkDoor(level, newPos)
		}
	}

	if currKeyState[sdl.SCANCODE_RIGHT] != 0 && prevKeyState[sdl.SCANCODE_RIGHT] == 0 {
		newPos := Pos{p.X + 1, p.Y}
		if canWalk(level, newPos) {
			level.Player.Pos = newPos
		} else {
			checkDoor(level, newPos)
		}
	}

	// switch ui.input. {
	// case Up:

	// case Down:
	// 	newPos := Pos{p.X, p.Y + 1}
	// 	if canWalk(level, newPos) {
	// 		level.Player.Move(newPos, level)
	// 	} else {
	// 		checkDoor(level, newPos)
	// 	}
	// case Left:
	// 	newPos := Pos{p.X - 1, p.Y}
	// 	if canWalk(level, newPos) {
	// 		level.Player.Move(newPos, level)
	// 	} else {
	// 		checkDoor(level, newPos)
	// 	}
	// case Right:
	// 	newPos := Pos{p.X + 1, p.Y}
	// 	if canWalk(level, Pos{p.X + 1, p.Y}) {
	// 		level.Player.Move(newPos, level)
	// 	} else {
	// 		checkDoor(level, newPos)
	// 	}
	// case Search:
	// 	level.astar(level.Player.Pos, Pos{3, 2})
	// case CloseWindow:
	// 	close(input.LevelChannel)
	// 	chanIndex := 0
	// 	for i, c := range game.LevelChans {
	// 		if c == input.LevelChannel {
	// 			chanIndex = i
	// 			break
	// 		}
	// 	}
	// 	game.LevelChans = append(game.LevelChans[:chanIndex], game.LevelChans[chanIndex+1:]...)
	// }
}

func canWalk(level *Level2, pos Pos) bool {

	if inRange(level, pos) {
		t := level.Map[pos.Y][pos.X]
		switch t {
		case StoneWall, DoorC, Blank:
			return false
		default:
			return true
		}
	}
	return false
}

func getNeighbours(level *Level2, pos Pos) []Pos {
	neighbours := make([]Pos, 0, 4)
	left := Pos{pos.X - 1, pos.Y}
	right := Pos{pos.X + 1, pos.Y}
	up := Pos{pos.X, pos.Y - 1}
	down := Pos{pos.X, pos.Y + 1}

	if canWalk(level, right) {
		neighbours = append(neighbours, right)
	}

	if canWalk(level, left) {
		neighbours = append(neighbours, left)
	}

	if canWalk(level, up) {
		neighbours = append(neighbours, up)
	}

	if canWalk(level, down) {
		neighbours = append(neighbours, down)
	}

	return neighbours
}

func (level *Level2) bfsFloor(start Pos) Tile {
	frontier := make([]Pos, 0, 8)
	frontier = append(frontier, start)
	visited := make(map[Pos]bool)
	visited[start] = true
	//level.Debug = visited

	for len(frontier) > 0 {
		current := frontier[0]

		currentTile := level.Map[current.Y][current.X]
		switch currentTile {
		case DirtFloor:
			return DirtFloor
		default:

		}

		frontier = frontier[1:] // pops the first appended element
		for _, next := range getNeighbours(level, current) {
			if !visited[next] {
				frontier = append(frontier, next)
				visited[next] = true
			}
		}

	}
	return DirtFloor
}
