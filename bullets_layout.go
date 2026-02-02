package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const bulletCooldown = 10 // frames (~6 bullets/sec)

type Bullet struct {
	x float64
	y float64
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		float32(b.x),
		float32(b.y),
		4,
		4,
		color.RGBA{255, 165, 80, 255}, // yellow bullets
		true,
	)
}
