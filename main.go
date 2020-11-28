package main

import (
	"fmt"

	"github.com/BiesGo/sdlWorkSpace/rpg/game"
	"github.com/BiesGo/sdlWorkSpace/rpg/ui"
)

func main() {
	fmt.Println("game currently starts in Main Menu \nPlay button leads user to Selection Menu\nStart button leads user to edit Menu\n\nIN EDIT MODE\ntools : 1 for dirt floor, 2 stone wall, 3 door, 4 mainCharacter(uses a placeholder texture)\nps. textures are created randomly within their scope\nleft click place(if possible such as wall and floor uses same layer but door is 1 layer above)\nright click remove most upper layer\n's' hard save\n'Backspace' reload latest save from file\nEnd of IN EDIT MODE")
	ui := &ui.UI2d{}
	game.Run(ui)

}

//game currently starts in Main Menu
// Start button leads user to Selection Menu
// in selection menu
//		start button leads user to end menu
//		edit button leads user to edit menu

// IN EDIT MODE
//for tool menu hold "tab" or "lshift"
//click on tool menu choices or use numbers 1,2,3...
//start button in tool menu leads user to end menu
// ps. textures are created randomly within their scope
//left click place(if possible such as wall and floor uses same layer but door is 1 layer above)
//right click remove most upper layer
// "s" hard save
// "Backspace" reload latest save from file
// End of IN EDIT MODE

//todo list
//code beauty(there may be unused/unnecessary parts left while changing whole structure)
//auto maze builder within scope
//left right up down buttons for travelling on level(currently left upmost part is shown automaticly)
// entity collision detection in build menu ie. doors cant overlap & there cant be more than 1 main character
// ~H~
