package main

type Circle struct {
	Center *Vec2
	Radius float64
	Velocity *Vec2
}

type Rect struct {
	Corner *Vec2
	Size *Vec2
	Velocity *Vec2
}

func RectFromCircle(c *Circle) Rect {
	length := c.Radius * 2
	radiusVec := Vector(c.Radius, c.Radius)

	lengthVec := Vector(length, length)

	cornerVec := SubVecVec(c.Center, &radiusVec)
	return Rect {
		&cornerVec,
		&lengthVec,
		c.Velocity,
	}
}

func CollidePointRect(pt Vec2, rc Rect) bool {
	return (pt.x > rc.Corner.x &&
			pt.x < rc.Corner.x + rc.Size.x &&
			pt.y > rc.Corner.y &&
			pt.y < rc.Corner.y + rc.Size.y)
}

func CollidePointCircle(pt Vec2, cr Circle) bool {
	if !CollidePointRect(pt, RectFromCircle(&cr)) {
		return false
	}

	hypot := SubVecVec(cr.Center, &pt)
	return hypot.Mag() < cr.Radius
}

func CollideRectRect(rect1 Rect, rect2 Rect) bool {
	return rect1.Corner.x < rect2.Corner.x + rect2.Size.x &&
		   rect1.Corner.x + rect1.Size.x > rect2.Corner.x &&
		   rect1.Corner.y < rect2.Corner.y + rect2.Size.y &&
		   rect1.Corner.y + rect1.Size.y > rect2.Corner.y
}

func CollideCircleCircle(c1 Circle, c2 Circle) bool{
	if (CollideRectRect(RectFromCircle(&c1), RectFromCircle(&c2))) {
		hypot := SubVecVec(c1.Center, c2.Center)
		return hypot.Mag() < (c1.Radius + c2.Radius)
	}

	return false
}

func ResolveCircleCircle(c1 Circle, c2 Circle) (Vec2, Vec2) {
	return VZero(), VZero()
}