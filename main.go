package main

import (
	"github.com/BiesGo/sdlWorkSpace/rpg/game"
	"github.com/BiesGo/sdlWorkSpace/rpg/ui"
)

func main() {
	ui := &ui.UI2d{}
	game.Run(ui)

}

//game currently starts in Main Menu
// Play button leads user to Selection Menu
// Start button leads user to edit Menu
// IN EDIT MODE
//tools : 1 for dirt floor, 2 stone wall, 3 door, 4 mainCharacter(uses a placeholder texture)
// ps. textures are created randomly within their scope
//left click place(if possible such as wall and floor uses same layer but door is 1 layer above)
//right click remove most upper layer
// "s" hard save

//todo list
//will implement entities above background
//code beauty(there may be unused/unnecessary parts left while changing whole structure)
//auto maze builder within scope
//left right up down buttons for travelling on level(currently left upmost part is shown automaticly)
// entity collision detection in build menu ie. doors cant overlap & there cant be more than 1 main character
//...
