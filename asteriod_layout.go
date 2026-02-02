package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

func (a *Asteroid) Draw(screen *ebiten.Image) {
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
}
