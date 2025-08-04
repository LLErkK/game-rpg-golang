package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	X, Y int
	Character
}

func NewPlayer(character Character) *Player {
	currentMap := GetCurrentMap()
	return &Player{
		X:         currentMap.SpawnX,
		Y:         currentMap.SpawnY,
		Character: character,
	}
}
func (p *Player) Draw() {
	color := rl.Blue
	switch p.Type {
	case KnightType:
		color = rl.Blue
	case RogueType:
		color = rl.Green
	case BerserkerType:
		color = rl.Red
	}

	screenX := int32((p.X - int(camera.X)) * tileSize)
	screenY := int32((p.Y - int(camera.Y)) * tileSize)

	rl.DrawRectangle(screenX, screenY, tileSize, tileSize, color)
}
