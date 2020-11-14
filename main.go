package main

import (
	"github.com/BiesGo/sdlWorkSpace/rpg/game"
	"github.com/BiesGo/sdlWorkSpace/rpg/ui"
)

func main() {
	ui := &ui.UI2d{}
	game.Run(ui)

}
