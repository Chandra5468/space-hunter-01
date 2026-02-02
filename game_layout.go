package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const playerSpeed = 4

type spaceshipDimensions float64

const (
	sshipWidth  spaceshipDimensions = 40.0
	sshipHeight spaceshipDimensions = 40.0
)

type Game struct { // storing state.
	playerX float64 // player x position : horizontal position
	playerY float64 // player y position : vertical position

	bullets   []Bullet // slice becoz : multiple bullets, dynamic add/remove
	asteroids []Asteroid

	fireCoolDown          int // if user presses on space, a lot of bullets are launched. Cool it down
	asteroidSpawnCooldown int
}

// Update runs at 60 times/sec by ebiten
func (g *Game) Update() error {

	if g.asteroidSpawnCooldown > 0 {
		g.asteroidSpawnCooldown--
	}

	if g.asteroidSpawnCooldown == 0 {
		g.asteroids = append(g.asteroids, Asteroid{
			x:     float64(rand.Intn(760) + 20),
			y:     -20,
			speed: float64(rand.Intn(3)+1) + 1,
			mass:  rand.Intn(20) + 20,
		})

		g.asteroidSpawnCooldown = asteroidSpawnRate
	}

	// Move asteroids
	for i := range g.asteroids {
		g.asteroids[i].y += g.asteroids[i].speed
	}

	// Remove off screen asteroids
	// activeAsteroids := g.asteroids[:0]

	// for _, a := range g.asteroids {
	// 	if a.y < 620 {
	// 		activeAsteroids = append(activeAsteroids, a)
	// 	}
	// }

	// g.asteroids = activeAsteroids

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

	// decrease cooldown every frame
	if g.fireCoolDown > 0 {
		g.fireCoolDown--
	}

	// fire an bullet on space (update)
	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.fireCoolDown == 0 {
		g.bullets = append(g.bullets, Bullet{
			x: g.playerX + float64(sshipWidth)/2 - 2, // width of aeroplane is 40 assuming
			y: g.playerY,
		})
		g.fireCoolDown = bulletCooldown // shooting feels controlled and professional
	}

	// move bullets logic (always)
	for i := range g.bullets {
		g.bullets[i].y -= 8
	}

	for ai := range g.asteroids {
		a := &g.asteroids[ai]

		for bi := range g.bullets {
			b := &g.bullets[bi]

			if b.y < 0 {
				continue // already dead
			}

			if bulletHitsAsteroid(*b, *a) {
				a.mass -= 5
				b.y = -1000 // mark bullet as dead
				break       // one bullet per asteroid per frame
			}
		}
	}

	// remove/clean bullets off screen
	activeBullets := g.bullets[:0]
	for _, b := range g.bullets {
		if b.y > 0 {
			activeBullets = append(activeBullets, b)
		}
	}

	g.bullets = activeBullets

	activeAsteroids := g.asteroids[:0]
	for _, a := range g.asteroids {
		if a.mass > 0 || a.y < 620 {
			activeAsteroids = append(activeAsteroids, a)
		}
	}
	g.asteroids = activeAsteroids

	// handle collision of bullets and asteroids
	return nil
}

// every frame
func (g *Game) Draw(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		float32(g.playerX),
		float32(g.playerY),
		float32(sshipWidth),
		float32(sshipHeight),
		color.White,
		true,
	)

	// now bullets will appear every frame
	for i := range g.bullets {
		g.bullets[i].Draw(screen)
	}

	// draw asteroids
	for i := range g.asteroids {
		g.asteroids[i].Draw(screen)
	}
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
