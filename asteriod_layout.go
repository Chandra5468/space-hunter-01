package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Asteroid struct {
	x     float64
	y     float64
	speed float64
	mass  int
}

const asteroidSpawnRate = 90 // frames (~1.5 sec) i.e each asteroid appears for every 1.5 seconds

func (a *Asteroid) Radius() float64 {
	return math.Sqrt(float64(a.mass)) * 6
}

func (a *Asteroid) Draw(screen *ebiten.Image, img *ebiten.Image) {
	/*
		r := float32(a.Radius())
		colorVal := uint8(100 + a.mass*3)
		vector.FillCircle(
			screen,
			float32(a.x),
			float32(a.y),
			r,
			color.RGBA{colorVal, colorVal, colorVal, 255}, // darker is damaged asteroid, ligher is healthy
			true,
		)
	*/

	op := &ebiten.DrawImageOptions{}

	w, h := img.Size()
	scale := a.Radius() / float64(w/2)

	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(
		a.x-float64(w)*scale/2,
		a.y-float64(h)*scale/2,
	)

	screen.DrawImage(img, op)
}
