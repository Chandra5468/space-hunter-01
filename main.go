package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	audioCtx := audio.NewContext(sampleRate)
	game := &Game{
		playerX:  800 / 2,
		playerY:  600 / 2,
		bullets:  []Bullet{}, // can omit it. Go can append on nil slice.
		audioCtx: audioCtx,
	}

	game.shipImage = loadImage("assets/playerShip3_orange.png")
	game.asteroidImage = loadImage("assets/meteorBrown_big4.png")
	game.laserSound = loadSound(audioCtx, "assets/laser9.ogg")

	ebiten.SetWindowSize(800, 600) // this sets the window size
	ebiten.SetWindowTitle("Space hunter")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
