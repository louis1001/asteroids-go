package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)


func init() {
	GeneratePlayerImage()
	GenerateBulletImage()
	GenerateAsteroidImages()
}

const (
	screenWidth = 640
	screenHeight = 480
)

type Game struct {
	player Player
	bullets []*Bullet
	asteroids []*Asteroid

	alreadyShooting bool
}

func (g *Game) Shoot() {
	newBullet := NewBullet(
		*g.player.GetPosition(),
		*g.player.GetDirection(),
	)

	g.bullets = append(g.bullets, newBullet)
}

func (g *Game) CleanBullets() {
	newBulletList := []*Bullet{}

	for _, b := range g.bullets {
		if b.alive {
			newBulletList = append(newBulletList, b)
		}
	}

	g.bullets = newBulletList
}

func (g *Game) Update() error {
	g.player.Update()

	for _, asteroid := range g.asteroids {
		asteroid.Update()
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.Accelerate(1)
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !g.alreadyShooting {
			g.Shoot()
			g.alreadyShooting = true
		}
	} else if g.alreadyShooting {
		g.alreadyShooting = false
	}

	deadBullet := false

	for _, bullet := range g.bullets {
		bullet.Update()
		if !(deadBullet || bullet.alive) {
			deadBullet = true
		}
	}

	if deadBullet {
		g.CleanBullets()
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.Rotate(-1)
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.Rotate(1)
	} else {
		g.player.Rotate(0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, bullet := range g.bullets {
		bullet.Draw(screen)
	}

	for _, asteroid := range g.asteroids {
		asteroid.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Asteroids")

	asteroids := []*Asteroid{
		NewAsteroid(),
		NewAsteroid(),
		NewAsteroid(),
		NewAsteroid(),
		NewAsteroid(),
		NewAsteroid(),
		NewAsteroid(),
		NewAsteroid(),
	}

	if err := ebiten.RunGame(&Game{NewPlayer(), []*Bullet{}, asteroids, false}); err != nil {
		log.Fatal(err)
	}
}