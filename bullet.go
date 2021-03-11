package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

const (
	bulletSpeed = 14
	bulletMaxLife = screenWidth

	bw = 2
	bh = 8
)

var bulletImage = ebiten.NewImage(bw, bh)

func GenerateBulletImage() {
	bulletImage.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
}

type Bullet struct {
	position, direction Vec2

	travelled float64
	alive bool
}

func NewBullet(pos, dir Vec2) *Bullet {
	return &Bullet {pos, dir, 0, true}
}

func (b *Bullet) Update() {
	delta := MulVecScal(b.direction, bulletSpeed)
	b.position.Add(delta)
	b.position.Wrap(borderBuffer, borderBuffer, screenWidth-borderBuffer, screenHeight-borderBuffer)

	b.travelled += bulletSpeed
	if b.travelled >= bulletMaxLife {
		b.alive = false
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-bw/2.0, -bh/2.0)
	op.GeoM.Rotate(b.direction.Heading() - math.Pi/2)
	op.GeoM.Translate(b.position.x, b.position.y)

	screen.DrawImage(bulletImage, op)
}