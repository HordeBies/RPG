package ui

import (
	"github.com/BiesGo/sdlWorkSpace/rpg/game"
	"github.com/veandco/go-sdl2/sdl"
)

type entityInterface interface {
	getX() int
	getY() int
	getRect() *sdl.Rect
}

func getX(intf interface{}) int {
	switch t := intf.(type) {
	case Door:
		return t.x
	case mainCharacter:
		return t.x
	}
	panic("error")
}

type entity struct {
	x, y    int
	srcRect *sdl.Rect
}

type Door struct {
	entity
	is_open bool
	//srcRectStorage []*sdl.Rect
}

func newDoor(obj game.Entity) *Door {
	if obj.Tile == game.DoorC {
		return &Door{entity{obj.X, obj.Y, &textureIndex[obj.Tile][0]}, false} // , make([]*sdl.Rect, 2)
	} else if obj.Tile == game.DoorO {
		return &Door{entity{obj.X, obj.Y, &textureIndex[obj.Tile][0]}, true}
	}
	panic("error")
}

func (d Door) getX() int {
	return d.x
}
func (d Door) getY() int {
	return d.y
}
func (d Door) getRect() *sdl.Rect {
	return d.srcRect
}

type mainCharacter struct {
	entity
}

func createMainCharacter(obj game.Entity) mainCharacter {
	return mainCharacter{entity{obj.X, obj.Y, &textureIndex[obj.Tile][0]}}
}

func (mc mainCharacter) getX() int {
	return mc.x
}
func (mc mainCharacter) getY() int {
	return mc.y
}
func (mc mainCharacter) getRect() *sdl.Rect {
	return mc.srcRect
}
