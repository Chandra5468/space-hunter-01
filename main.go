package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &Game{
		playerX: 800 / 2,
		playerY: 600 / 2,
		bullets: []Bullet{}, // can omit it. Go can append on nil slice.
	}

	ebiten.SetWindowSize(800, 600) // this sets the window size
	ebiten.SetWindowTitle("Space hunter")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
