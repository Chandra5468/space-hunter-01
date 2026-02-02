package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const playerSpeed = 4

type Game struct { // storing state.
	playerX float64 // player x position : horizontal position
	playerY float64 // player y position : vertical position
}

// Update runs at 60 times/sec by ebiten
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.playerX -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.playerX += playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.playerY -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.playerY += playerSpeed
	}

	g.clampPlayer()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) { // every frame
	ebitenutil.DrawRect(
		screen,
		g.playerX,
		g.playerY,
		40,
		40,
		color.White,
	)
	// vector.FillRect(
	// 	screen,
	// 	float32(g.playerX),
	// 	float32(g.playerY),
	// 	40,
	// 	40,
	// 	color.White,
	// )
}

// This is Game world size
// This tells Ebiten my logical game canvas is 800 * 600
// Everything you draw uses this coordinate system. (0,0) top-left (800,600) bottom-right
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func (g *Game) clampPlayer() {
	if g.playerX < 0 {
		g.playerX = 0
	}
	if g.playerX > 800-40 {
		g.playerX = 800 - 40
	}

	if g.playerY < 0 {
		g.playerY = 0
	}
	if g.playerY > 600-40 {
		g.playerY = 600 - 40
	}
}
