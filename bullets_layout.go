package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

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
		10,
		color.RGBA{255, 255, 0, 255}, // yellow bullets
		true,
	)
}
