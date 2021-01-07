package ui

import (
	"math/rand"

	"github.com/BiesGo/sdlWorkSpace/rpg/game"
	"github.com/veandco/go-sdl2/sdl"
)

func playMenu(ui *UI2d) stateFunc {

	newLevel := GlobalLevel2

	// offsetX := int32((winWidth / 2) - centerX*32)
	// offsetY := int32((winHeight / 2) - centerY*32)

	r := rand.New(rand.NewSource(1))

	renderer.Clear()
	for y, row := range newLevel.Map {
		for x, tile := range row {

			if tile != game.Blank {
				srcRects := textureIndex[tile]
				srcRect := srcRects[r.Intn(len(srcRects))] // get a random tile from a specific group of rects,
				//this makes difference if the variaton count of the current tile is greater than 1
				dstRect := sdl.Rect{int32(x * 32), int32(y * 32), 32, 32}
				// for seeing how breadth-first search works, can be said that this is going to be for debugging purposes
				renderer.Copy(textureAtlas, &srcRect, &dstRect)
			}

		}
	}

	game.HandleInput(ui.input.currKeyState, ui.input.prevKeyState, newLevel)

	playerSrcRect := textureIndex['P'][0]
	renderer.Copy(textureAtlas, &playerSrcRect, &sdl.Rect{int32(newLevel.Player.X) * 32, int32(newLevel.Player.Y) * 32, 32, 32})

	return determineToken
}
