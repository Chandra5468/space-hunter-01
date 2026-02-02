package main

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

	score    int
	gameOver bool

	shipImage     *ebiten.Image
	asteroidImage *ebiten.Image
}

// Update runs at 60 times/sec by ebiten
func (g *Game) Update() error {
	// CORRECT ORDER OF UPDATE
	/*
		1. early exit if game over
		2. spawn asteroids
		3. move asteroids
		4. move bullets
		5. player movement
		6. player asteroid collision
		7. bullet asteroid collision
		8. cleanup bullets
		9. cleanup asteroids
	*/

	if g.gameOver {
		if g.gameOver && ebiten.IsKeyPressed(ebiten.KeyR) {
			g.reset()
		}
		return nil
	}

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

	// asteroid movement
	for _, a := range g.asteroids {
		if playerHitsAsteroid(g.playerX, g.playerY, a) {
			g.gameOver = true
			return nil
		}
	}

	// Move spaceship
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

	// Bullet asteroid collision
	for ai := range g.asteroids {
		a := &g.asteroids[ai]

		for bi := range g.bullets {
			b := &g.bullets[bi]

			if b.y < 0 {
				continue // already dead
			}

			if bulletHitsAsteroid(*b, *a) {
				a.mass -= 5
				b.y = -1000   // mark bullet as dead
				g.score += 10 // score per hit
				break         // one bullet per asteroid per frame
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

	// clean up asteroids
	activeAsteroids := g.asteroids[:0]
	for _, a := range g.asteroids {
		if a.mass > 0 { // add this or condition after testing || a.y < 620
			activeAsteroids = append(activeAsteroids, a)
		} else {
			g.score += 50 // bonus for destroying asteroid
		}
	}
	g.asteroids = activeAsteroids

	return nil
}

// every frame
func (g *Game) Draw(screen *ebiten.Image) {

	// to see score increasing live
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("Score: %d", g.score),
	)

	if g.gameOver {
		ebitenutil.DebugPrintAt(
			screen,
			"GAME OVER\nPress R to Restart",
			280, 260,
		)
		return
	}

	// vector.FillRect(
	// 	screen,
	// 	float32(g.playerX),
	// 	float32(g.playerY),
	// 	float32(sshipWidth),
	// 	float32(sshipHeight),
	// 	color.White,
	// 	true,
	// )
	// Instead of fill Rect we will load spaceship image
	op := &ebiten.DrawImageOptions{}
	shipW, shipH := g.shipImage.Size()
	// g.shipImage.Bounds()

	op.GeoM.Translate(
		g.playerX-float64(shipW)/2,
		g.playerY-float64(shipH)/2,
	)

	screen.DrawImage(g.shipImage, op)

	// now bullets will appear every frame
	for i := range g.bullets {
		g.bullets[i].Draw(screen)
	}

	// draw asteroids
	for i := range g.asteroids {
		g.asteroids[i].Draw(screen, g.asteroidImage)
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

func (g *Game) reset() {
	g.bullets = nil
	g.asteroids = nil
	g.score = 0
	g.fireCoolDown = 0
	g.asteroidSpawnCooldown = 60
	g.playerX = 800 / 2
	g.playerY = 600 / 2
	g.gameOver = false
}
