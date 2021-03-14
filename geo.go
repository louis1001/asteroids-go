package main

import (
	"math"
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

func (v *Vec2) Mag() float64 {
	return math.Hypot(v.x, v.y)
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

func (v1 *Vec2) Clone() Vec2 {
	return Vec2{v1.x, v1.y}
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
