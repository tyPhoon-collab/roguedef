package trait

import (
	"fmt"
	"image/color"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	ebitenVector "github.com/hajimehoshi/ebiten/v2/vector"
)

type Circle struct {
	offset Vec2
	radius float64

	*Transform
}

func (c *Circle) Intersects(other Intersector) bool {
	switch other := other.(type) {
	case *Circle:
		return c.IntersectsCircle(other)
	default:
		fmt.Println("type is not associated:", reflect.TypeOf(other))
		return false
	}
}

func (c *Circle) IntersectsCircle(other *Circle) bool {
	center := c.ScaledCenter()
	otherCenter := other.ScaledCenter()

	radius := c.ScaledRadius()
	otherRadius := other.ScaledRadius()
	return center.Sub(otherCenter).Len() < radius+otherRadius
}

func (c *Circle) Trans() *Transform {
	return c.Transform
}

func (c *Circle) ScaledCenter() Vec2 {
	return c.Pos.Add(c.offset.Mul(c.Scale))
}

func (c *Circle) ScaledRadius() float64 {
	return c.radius * c.Scale.X
}

func (c *Circle) Draw(screen *ebiten.Image) {
	center := c.ScaledCenter()
	x := float32(center.X)
	y := float32(center.Y)
	radius := float32(c.ScaledRadius())
	ebitenVector.StrokeCircle(screen, x, y, radius, 2, color.White, false)
}

func (c *Circle) WithTransform(transform *Transform) *Circle {
	if transform == nil {
		transform = NewTransform()
	}
	c.Transform = transform
	return c
}

func (c *Circle) WithRadius(radius float64) *Circle {
	c.radius = radius
	return c
}

func (c *Circle) FromImage(image *ebiten.Image) *Circle {
	circle := calculateRadiusFromImage(image)
	center := calculateCenterFromImage(image)

	c.radius = circle
	c.offset = center

	return c
}

func calculateRadiusFromImage(image *ebiten.Image) float64 {
	return float64(image.Bounds().Dx()) / 2
}

func calculateCenterFromImage(image *ebiten.Image) Vec2 {
	return Vec2{
		X: float64(image.Bounds().Dx()) / 2,
		Y: float64(image.Bounds().Dy()) / 2,
	}
}

func NewCircle() *Circle {
	return (&Circle{}).WithTransform(nil)
}
