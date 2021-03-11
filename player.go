package main

import(
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"log"
)

type Vec2 struct {
	x, y float64
}

func VZero() Vec2 {
	return Vec2{0, 0}
}

func Vector(x, y float64) Vec2 {
	return Vec2{x, y}
}

func (v *Vec2) Rotate(theta float64) {
	v1 := *v

	v.x = math.Cos(theta)*v1.x - math.Sin(theta)*v1.y
	v.y = math.Sin(theta)*v1.x + math.Cos(theta)*v1.y
}

func (v *Vec2) Heading() float64 {
	a := math.Atan(v.y/v.x)
	if v.x < 0 {
		a += math.Pi
	}
	return a
}

func (v1 *Vec2) Add(v2 Vec2) *Vec2 {
	v1.x += v2.x
	v1.y += v2.y

	return v1
}

func MulVecScal(v Vec2, scal float64) Vec2 {
	v.x *= scal
	v.y *= scal
	return v
}

func AddVecVec(v1, v2 *Vec2) Vec2 {
	return Vec2{v1.x+v2.x, v1.y+v2.y}
}

func SubVecVec(v1 *Vec2, v2 *Vec2) Vec2 {
	return Vec2{v1.x-v2.x, v1.y-v2.y}
}

func (v *Vec2)Wrap(minX, minY, maxX, maxY float64) {
	if v.x > maxX { v.x = minX }
	if v.x < minX { v.x = maxX }

	if v.y > maxY { v.y = minY }
	if v.y < minY { v.y = maxY }
}

const pw = 50
const ph = 50

var player_image = ebiten.NewImage(pw, ph)

func GeneratePlayerImage() {
	stored_image := GetImageFromFilePath("./assets/ship.png")


	if stored_image == nil {
		// Something
		log.Fatal("Failed to load the image!")
	}

	player_image.DrawImage(stored_image, nil)
}

const borderBuffer = -float64(pw)*0.3

const rotationSpeed = 0.09
const drag = 0.2
const maxSpeed = 8

type Player struct {
	position Vec2
	direction Vec2

	rotationDirection float64

	velocity float64
}

func (p *Player) GetPosition() *Vec2 {
	return &p.position
}

func (p *Player) GetDirection() *Vec2 {
	return &p.direction
}

func NewPlayer() Player {
	return Player {Vec2{screenWidth/2, screenHeight/2}, Vec2{0, -1}, 0, 0}
}

// Add drift
// Inertia and acceleration instead of just move
// in the direction I point to

func (p *Player)Update() {
	if p.velocity > maxSpeed{
		p.velocity = maxSpeed
	} else if p.velocity > 0 {
		p.velocity -= drag
	} else {
		p.velocity = 0
	}

	if p.rotationDirection > 0 {
		p.direction.Rotate(rotationSpeed)
	} else if p.rotationDirection < 0{
		p.direction.Rotate(-rotationSpeed)
	}

	delta := MulVecScal(p.direction, p.velocity)
	p.position.Add(delta)

	p.position.Wrap(borderBuffer, borderBuffer, screenWidth-borderBuffer, screenHeight-borderBuffer)
}

func (p *Player) Rotate(dir float64) {
	p.rotationDirection = dir
}

func (p *Player) Accelerate(ammount float64) {
	p.velocity += ammount
}


func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-pw/2.0, -ph/2.0)
	op.GeoM.Rotate(p.direction.Heading())
	op.GeoM.Translate(p.position.x, p.position.y)

	screen.DrawImage(player_image, op)
}
