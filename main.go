package main

import (
	"github.com/BiesGo/sdlWorkSpace/rpg/game"
	"github.com/BiesGo/sdlWorkSpace/rpg/ui"
)

func main() {
	ui := &ui.UI2d{}
	game.Run(ui)

}

//game currently starts in Build Menu
//tools : 1 for dirt floor, 2 stone wall, 3 door, 4 mainCharacter(uses a placeholder texture)
// ps. textures are created randomly within their scope
//left click place(if possible such as wall and floor uses same layer but door is 1 layer above)
//right click remove most upper layer
// "s" save
