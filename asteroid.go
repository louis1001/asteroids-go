package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"math/rand"
	"time"
	"fmt"
)

const (
	asteroidSize = 100
	rubbleSize = 50
)

type AsteroidKind int

const (
	rubble AsteroidKind = iota
	rock = iota
	meteor = iota
)

var asteroidImages = []*ebiten.Image{}
var rubbleImages = []*ebiten.Image{}

const (
	asteroidImagePrefix = "asteroid"
	asteroidImageCount = 5

	asteroidMaxSpeed = 0.3
	asteroidMaxRotSpeed = 0.03
)

const (
	rubbleImagePrefix = "rubble"
	rubbleImageCount = 5
)

type Asteroid struct {
	image *ebiten.Image

	kind AsteroidKind
	size float64

	position Vec2
	velocity Vec2

	rotation float64
	rotationSpeed float64

	alive bool
}

func (a *Asteroid) Kill() {
	a.alive = false
}

func (a *Asteroid) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-a.size/2.0, -a.size/2.0)
	op.GeoM.Rotate(a.rotation)
	op.GeoM.Translate(a.position.x, a.position.y)

	screen.DrawImage(a.image, op)
}

func (a *Asteroid) Update() {
	a.position.Add(a.velocity)

	a.position.Wrap(
		borderBuffer,
		borderBuffer,
		screenWidth-borderBuffer,
		screenHeight-borderBuffer,
	)

	a.rotation += a.rotationSpeed
}

func (a *Asteroid) Destroy(dir Vec2) (debris []*Asteroid) {
	a.Kill()
	debris = []*Asteroid{}
	if a.kind == rubble {
		return debris
	}

	ammount := int(a.kind)+1
	for i := 0; i < ammount; i++ {
		newAsteroid := NewAsteroid(a.kind-1)

		randRotation := randFloat(-math.Pi/4, math.Pi/4)
		newDirection := dir
		newDirection.Rotate(randRotation)
		newAsteroid.velocity = MulVecScal(newDirection, a.velocity.Mag()*4)

		newPosition := a.position.Clone()
		newPosition.Add(MulVecScal(newDirection, 20))

		newAsteroid.position = newPosition

		debris = append(debris, newAsteroid)
	}

	return debris
}

func randFloat(min, max float64) float64 {
	return rand.Float64() * (max-min) + min
}

func randInt(min, max int) int {
	return (rand.Int() % (max-min)) + min
}

func pickRandomImage(slc []*ebiten.Image) *ebiten.Image {
	imageIndex := randInt(0, len(slc))
	return slc[imageIndex]
}

func GenerateAsteroidImages() {
	rand.Seed(time.Now().Unix())

	for i := 0; i < asteroidImageCount; i++ {
		fileName := fmt.Sprintf("./assets/%s%d.png", asteroidImagePrefix, i)
		loadedImage := GetImageFromFilePath(fileName)
		asteroidImages = append(asteroidImages, loadedImage)
	}

	for i := 0; i < rubbleImageCount; i++ {
		fileName := fmt.Sprintf("./assets/%s%d.png", rubbleImagePrefix, i)
		loadedImage := GetImageFromFilePath(fileName)
		rubbleImages = append(rubbleImages, loadedImage)
	}
}

func NewAsteroid(kind AsteroidKind) *Asteroid {
	randPos := Vector(
		float64(randInt(0, screenWidth)),
		float64(randInt(0, screenHeight)),
	)
	randVel := Vector(
		randFloat(-asteroidMaxSpeed, asteroidMaxSpeed),
		randFloat(-asteroidMaxSpeed, asteroidMaxSpeed),
	)

	var img *ebiten.Image
	var size float64
	if kind == rubble {
		img = pickRandomImage(rubbleImages)
		size = rubbleSize
	} else if kind == rock {
		img = pickRandomImage(asteroidImages)
		size = asteroidSize
	}

	rot := randFloat(0, math.Pi)

	return &Asteroid{
		img,
		kind,
		size,
		randPos,
		randVel,
		rot,
		randFloat(-asteroidMaxRotSpeed, asteroidMaxRotSpeed),
		true,
	}
}