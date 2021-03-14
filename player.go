package main

import(
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

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
