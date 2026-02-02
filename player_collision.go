package main

import "math"

func playerHitsAsteroid(px, py float64, a Asteroid) bool {
	playerRadius := float64(sshipWidth) / 2
	dx := (px + playerRadius) - a.x
	dy := (py + playerRadius) - a.y

	distance := math.Sqrt(dx*dx + dy*dy)

	return distance < (playerRadius + a.Radius())
}
