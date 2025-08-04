package main

import rl "github.com/gen2brain/raylib-go/raylib"

type TileType int

const (
	Wall TileType = iota
	Floor
	Door
	Stairs
	Water
	Grass
	Portal
)

type GameMap struct {
	Name   string
	Width  int
	Height int
	Data   [][]TileType
	SpawnX int
	SpawnY int
}

type Camera struct {
	X, Y float32
}

var currentMapIndex = 0
var gameMaps []GameMap
var camera Camera

func InitMaps() {
	gameMaps = []GameMap{
		{
			Name:   "Starting Village",
			Width:  20,
			Height: 15,
			SpawnX: 2,
			SpawnY: 2,
			Data: [][]TileType{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
				{0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0},
				{0, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0},
				{0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
				{0, 1, 1, 0, 0, 0, 1, 1, 1, 3, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0},
				{0, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0},
				{0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			Name:   "Dark Forest",
			Width:  25,
			Height: 20,
			SpawnX: 1,
			SpawnY: 18,
			Data: [][]TileType{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 5, 5, 5, 0, 5, 5, 5, 5, 5, 0, 0, 5, 5, 5, 5, 0, 5, 5, 5, 5, 5, 0, 5, 0},
				{0, 5, 1, 5, 0, 5, 1, 1, 1, 5, 0, 0, 5, 1, 1, 5, 0, 5, 1, 1, 1, 5, 0, 5, 0},
				{0, 5, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 5, 0},
				{0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0},
				{0, 5, 1, 5, 0, 5, 1, 5, 0, 5, 1, 1, 5, 1, 5, 0, 5, 1, 1, 5, 0, 5, 1, 5, 0},
				{0, 5, 1, 5, 0, 5, 1, 5, 0, 5, 1, 1, 5, 1, 5, 0, 5, 1, 1, 5, 0, 5, 1, 5, 0},
				{0, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 0},
				{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0},
				{0, 5, 5, 1, 5, 0, 5, 1, 5, 0, 1, 1, 5, 1, 5, 0, 5, 1, 1, 5, 0, 1, 5, 5, 0},
				{0, 5, 5, 1, 5, 0, 5, 1, 5, 0, 1, 1, 5, 1, 5, 0, 5, 1, 1, 5, 0, 1, 5, 5, 0},
				{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0},
				{0, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 0},
				{0, 5, 1, 5, 0, 5, 1, 5, 0, 5, 1, 1, 5, 1, 5, 0, 5, 1, 1, 5, 0, 5, 1, 5, 0},
				{0, 5, 1, 5, 0, 5, 1, 5, 0, 5, 1, 1, 5, 1, 5, 0, 5, 1, 1, 5, 0, 5, 1, 5, 0},
				{0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0},
				{0, 5, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 5, 0},
				{0, 5, 1, 5, 0, 5, 1, 1, 1, 5, 0, 0, 5, 1, 1, 5, 0, 5, 1, 1, 1, 5, 0, 5, 0},
				{0, 5, 5, 5, 0, 5, 5, 5, 5, 5, 0, 0, 5, 5, 5, 5, 0, 5, 5, 5, 5, 5, 0, 5, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			Name:   "Crystal Cave",
			Width:  30,
			Height: 25,
			SpawnX: 15,
			SpawnY: 23,
			Data: [][]TileType{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 4, 4, 1, 1, 1, 1, 1, 4, 4, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 1, 1, 0, 0, 4, 4, 4, 4, 1, 1, 1, 4, 4, 4, 4, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 0, 0, 4, 4, 4, 4, 4, 4, 1, 4, 4, 4, 4, 4, 4, 0, 0, 1, 1, 1, 0, 0, 0, 0},
				{0, 0, 1, 1, 1, 0, 0, 4, 4, 4, 0, 0, 4, 4, 1, 4, 4, 0, 0, 4, 4, 4, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 1, 1, 0, 0, 4, 4, 4, 0, 0, 0, 0, 4, 1, 4, 0, 0, 0, 0, 4, 4, 4, 0, 0, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 0, 4, 4, 4, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 4, 4, 4, 0, 0, 1, 1, 0, 0},
				{0, 1, 1, 0, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 4, 4, 4, 0, 1, 1, 0, 0},
				{0, 1, 0, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 4, 4, 4, 0, 1, 0, 0},
				{0, 1, 1, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 4, 4, 1, 1, 0, 0},
				{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
				{0, 1, 1, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 4, 4, 1, 1, 0, 0},
				{0, 1, 0, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 4, 4, 4, 0, 1, 0, 0},
				{0, 1, 1, 0, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 4, 4, 4, 0, 1, 1, 0, 0},
				{0, 1, 1, 0, 0, 4, 4, 4, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 4, 4, 4, 0, 0, 1, 1, 0, 0},
				{0, 0, 1, 1, 0, 0, 4, 4, 4, 0, 0, 0, 0, 4, 1, 4, 0, 0, 0, 0, 4, 4, 4, 0, 0, 1, 1, 0, 0, 0},
				{0, 0, 1, 1, 1, 0, 0, 4, 4, 4, 0, 0, 4, 4, 1, 4, 4, 0, 0, 4, 4, 4, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 0, 0, 4, 4, 4, 4, 4, 4, 1, 4, 4, 4, 4, 4, 4, 0, 0, 1, 1, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 1, 1, 0, 0, 4, 4, 4, 4, 1, 1, 1, 4, 4, 4, 4, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 4, 4, 1, 1, 1, 1, 1, 4, 4, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
}

func GetCurrentMap() *GameMap {
	if currentMapIndex >= 0 && currentMapIndex < len(gameMaps) {
		return &gameMaps[currentMapIndex]
	}
	return &gameMaps[0]
}

func IsWalkable(x, y int) bool {
	currentMap := GetCurrentMap()

	if x < 0 || y < 0 || x >= currentMap.Width || y >= currentMap.Height {
		return false
	}

	tile := currentMap.Data[y][x]
	return tile == Floor || tile == Grass
}

func UpdateCamera(playerX, playerY int) {
	currentMap := GetCurrentMap()
	viewWidth := mapWidth
	viewHeight := mapHeight

	// Center camera on player
	camera.X = float32(playerX) - float32(viewWidth)/2
	camera.Y = float32(playerY) - float32(viewHeight)/2

	// Clamp camera to map bounds
	if camera.X < 0 {
		camera.X = 0
	}
	if camera.Y < 0 {
		camera.Y = 0
	}
	if camera.X > float32(currentMap.Width-viewWidth) {
		camera.X = float32(currentMap.Width - viewWidth)
	}
	if camera.Y > float32(currentMap.Height-viewHeight) {
		camera.Y = float32(currentMap.Height - viewHeight)
	}
}

func DrawMap() {
	currentMap := GetCurrentMap()

	// Calculate visible area
	startX := int(camera.X)
	startY := int(camera.Y)
	endX := startX + mapWidth
	endY := startY + mapHeight

	// Clamp to map bounds
	if endX > currentMap.Width {
		endX = currentMap.Width
	}
	if endY > currentMap.Height {
		endY = currentMap.Height
	}

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			if x >= 0 && y >= 0 && x < currentMap.Width && y < currentMap.Height {
				tile := currentMap.Data[y][x]
				color := getTileColor(tile)

				screenX := (x - startX) * tileSize
				screenY := (y - startY) * tileSize

				rl.DrawRectangle(int32(screenX), int32(screenY), tileSize, tileSize, color)
			}
		}
	}
}

func getTileColor(tile TileType) rl.Color {
	switch tile {
	case Wall:
		return rl.DarkGray
	case Floor:
		return rl.LightGray
	case Door:
		return rl.Brown
	case Stairs:
		return rl.Yellow
	case Water:
		return rl.Blue
	case Grass:
		return rl.Green
	case Portal:
		return rl.Red
	default:
		return rl.Black
	}
}

func ChangeMap(mapIndex int) {
	if mapIndex >= 0 && mapIndex < len(gameMaps) {
		currentMapIndex = mapIndex
	}
}

func GetMapCount() int {
	return len(gameMaps)
}
