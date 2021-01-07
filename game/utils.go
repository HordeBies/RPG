package game

func inRange(level *Level2, pos Pos) bool {
	return pos.X < len(level.Map[0]) && pos.Y < len(level.Map) && pos.X >= 0 && pos.Y >= 0
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
