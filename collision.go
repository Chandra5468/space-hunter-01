package main

import "math"

// computes distance between bullet and asteroid center
// if distance < asteroid radius -> hit
func bulletHitsAsteroid(b Bullet, a Asteroid) bool {
	dx := b.x - a.x
	dy := b.y - a.y

	distance := math.Sqrt(dx*dx + dy*dy)
	return distance < a.Radius()
}
