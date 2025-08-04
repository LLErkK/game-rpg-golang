package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	tileSize     = 32
	mapWidth     = 10
	mapHeight    = 8
	screenWidth  = tileSize*mapWidth + 300
	screenHeight = tileSize * mapHeight
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "RPG Sederhana - Terstruktur")
	defer rl.CloseWindow()

	InitMaps()
	
	game := NewGame()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}
}
