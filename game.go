package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState int

const (
	ChoosingCharacter GameState = iota
	Playing
	GameOver
)

type Game struct {
	State           GameState
	Player          *Player
	Characters      []Character
	Selected        int
	LastMove        time.Time
	StepDelay       time.Duration
	ShowStats       bool
	LastStatsToggle time.Time
}

func NewGame() *Game {
	return &Game{
		State:      ChoosingCharacter,
		Characters: GetStarterCharacters(),
		Selected:   0,
		StepDelay:  350 * time.Millisecond,
		ShowStats:  false,
	}
}

func (g *Game) Update() {
	switch g.State {
	case ChoosingCharacter:
		g.updateCharacterSelection()
	case Playing:
		g.updatePlaying()
	case GameOver:
		g.updateGameOver()
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	switch g.State {
	case ChoosingCharacter:
		g.drawCharacterSelection()
	case Playing:
		g.drawGame()
	case GameOver:
		g.drawGameOver()
	}

	rl.EndDrawing()
}

func (g *Game) updateCharacterSelection() {
	if rl.IsKeyPressed(rl.KeyRight) {
		g.Selected = (g.Selected + 1) % len(g.Characters)
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		g.Selected = (g.Selected - 1 + len(g.Characters)) % len(g.Characters)
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		chosen := g.Characters[g.Selected]
		g.Player = NewPlayer(chosen)
		g.State = Playing
	}
}
func (g *Game) updatePlaying() {
	g.updateMovement()

	// Toggle stats display
	now := time.Now()
	if rl.IsKeyPressed(rl.KeyTab) && now.Sub(g.LastStatsToggle) > 200*time.Millisecond {
		g.ShowStats = !g.ShowStats
		g.LastStatsToggle = now
	}

	// Test level up (untuk testing)
	if rl.IsKeyPressed(rl.KeySpace) {
		g.Player.Character.GainExp(50)
	}
	if rl.IsKeyPressed(rl.KeyM) {
		nextMap := (currentMapIndex + 1) % GetMapCount()
		ChangeMap(nextMap)
		newMap := GetCurrentMap()
		g.Player.X = newMap.SpawnX
		g.Player.Y = newMap.SpawnY
	}
	if rl.IsKeyPressed(rl.KeyN) {
		nextMap := (currentMapIndex - 1 + GetMapCount()) % GetMapCount()
		ChangeMap(nextMap)
		newMap := GetCurrentMap()
		g.Player.X = newMap.SpawnX
		g.Player.Y = newMap.SpawnY
	}
}

// Hanya untuk test
func (g *Game) updateGameOver() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		*g = *NewGame() // Reset game
	}
}

func (g *Game) drawCharacterSelection() {
	rl.DrawText("Pilih Karakter (← → Enter)", 100, 30, 20, rl.Black)
	rl.DrawText("Setiap karakter memiliki scaling yang berbeda!", 100, 60, 16, rl.DarkGray)

	for i, c := range g.Characters {
		color := rl.Gray
		if i == g.Selected {
			color = rl.DarkBlue
		}

		x := int32(100 + i*200)
		y := int32(100)

		rl.DrawText(c.Name, x, y, 20, color)
		rl.DrawText("HP: "+itoa(c.HP), x, y+30, 16, color)
		rl.DrawText("ATK: "+itoa(c.ATK), x, y+50, 16, color)
		rl.DrawText("DEF: "+itoa(c.DEF), x, y+70, 16, color)
		rl.DrawText("SPD: "+itoa(c.SPD), x, y+90, 16, color)
		rl.DrawText("CRIT: "+ftoa(c.CRIT), x, y+110, 16, color)

		// Show scaling info
		scaling := characterScalings[c.Type]
		rl.DrawText("--- Per Level ---", x, y+140, 12, rl.DarkGray)
		rl.DrawText("HP: +"+itoa(scaling.HPGrowth), x, y+155, 12, rl.DarkGray)
		rl.DrawText("ATK: +"+itoa(scaling.ATKGrowth), x, y+170, 12, rl.DarkGray)
		rl.DrawText("DEF: +"+itoa(scaling.DEFGrowth), x, y+185, 12, rl.DarkGray)
		rl.DrawText("SPD: +"+itoa(scaling.SPDGrowth), x, y+200, 12, rl.DarkGray)
	}
}
func (g *Game) drawGame() {

	UpdateCamera(g.Player.X, g.Player.Y)

	DrawMap()
	g.Player.Draw()

	// Draw UI
	g.drawUI()

	if g.ShowStats {
		g.drawDetailedStats()
	}
}

func (g *Game) drawUI() {
	// Basic player info
	uiX := int32(mapWidth*tileSize + 10)
	currentMap := GetCurrentMap()
	rl.DrawText("Map: "+currentMap.Name, uiX, 10, 14, rl.DarkBlue)
	rl.DrawText("Player: "+g.Player.Name, uiX, 30, 16, rl.Black)
	rl.DrawText("Level: "+itoa(g.Player.Level), uiX, 50, 16, rl.Black)
	rl.DrawText("HP: "+itoa(g.Player.HP)+"/"+itoa(g.Player.MaxHP), uiX, 70, 16, rl.Black)
	rl.DrawText("EXP: "+itoa(g.Player.EXP)+"/"+itoa(g.Player.EXPToNextLevel), uiX, 90, 16, rl.Black)
	rl.DrawText("Pos: ("+itoa(g.Player.X)+","+itoa(g.Player.Y)+")", uiX, 110, 12, rl.DarkGray)

	// Controls
	rl.DrawText("TAB: Toggle Stats", uiX, 140, 12, rl.DarkGray)
	rl.DrawText("SPACE: Gain EXP (test)", uiX, 155, 12, rl.DarkGray)
	rl.DrawText("M: Next Map", uiX, 170, 12, rl.DarkGray)
	rl.DrawText("N: Prev Map", uiX, 185, 12, rl.DarkGray)
}

func (g *Game) drawDetailedStats() {
	uiX := int32(mapWidth*tileSize + 10)
	startY := int32(160)

	rl.DrawText("--- Detailed Stats ---", uiX, startY, 14, rl.DarkBlue)
	rl.DrawText("ATK: "+itoa(g.Player.ATK), uiX, startY+20, 12, rl.Black)
	rl.DrawText("DEF: "+itoa(g.Player.DEF), uiX, startY+35, 12, rl.Black)
	rl.DrawText("SPD: "+itoa(g.Player.SPD), uiX, startY+50, 12, rl.Black)
	rl.DrawText("CRIT: "+ftoa(g.Player.CRIT), uiX, startY+65, 12, rl.Black)

	// Show next level growth
	scaling := characterScalings[g.Player.Type]
	rl.DrawText("--- Next Level ---", uiX, startY+90, 12, rl.DarkGreen)
	rl.DrawText("HP: +"+itoa(scaling.HPGrowth), uiX, startY+105, 11, rl.DarkGreen)
	rl.DrawText("ATK: +"+itoa(scaling.ATKGrowth), uiX, startY+120, 11, rl.DarkGreen)
	rl.DrawText("DEF: +"+itoa(scaling.DEFGrowth), uiX, startY+135, 11, rl.DarkGreen)
	rl.DrawText("SPD: +"+itoa(scaling.SPDGrowth), uiX, startY+150, 11, rl.DarkGreen)
}

func (g *Game) drawGameOver() {
	rl.DrawText("Game Over!", 200, 100, 32, rl.Red)
	rl.DrawText("Press Enter to restart", 200, 150, 20, rl.Black)
}

func (g *Game) updateMovement() {
	now := time.Now()
	if now.Sub(g.LastMove) >= g.StepDelay {
		moved := false
		if rl.IsKeyDown(rl.KeyRight) && IsWalkable(g.Player.X+1, g.Player.Y) {
			g.Player.X++
			moved = true
		}
		if rl.IsKeyDown(rl.KeyLeft) && IsWalkable(g.Player.X-1, g.Player.Y) {
			g.Player.X--
			moved = true
		}
		if rl.IsKeyDown(rl.KeyUp) && IsWalkable(g.Player.X, g.Player.Y-1) {
			g.Player.Y--
			moved = true
		}
		if rl.IsKeyDown(rl.KeyDown) && IsWalkable(g.Player.X, g.Player.Y+1) {
			g.Player.Y++
			moved = true
		}
		if moved {
			g.LastMove = now
		}
	}
}
