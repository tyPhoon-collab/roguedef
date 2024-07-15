package trait

import (
	"fmt"
	"image/color"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	ebitenVector "github.com/hajimehoshi/ebiten/v2/vector"
)

type Circle struct {
	Offset Vec2
	Radius float64

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
	center := c.Center()
	otherCenter := other.Center()
	return center.Sub(otherCenter).Len() < c.Radius+other.Radius
}

func (c *Circle) Trans() *Transform {
	return c.Transform
}

func (c *Circle) Center() Vec2 {
	return c.Pos.Add(c.Offset)
}

func (c *Circle) Draw(screen *ebiten.Image) {
	center := c.Center()
	x := float32(center.X)
	y := float32(center.Y)
	radius := float32(c.Radius)
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
	c.Radius = radius
	return c
}

func (c *Circle) FromImage(image *ebiten.Image) *Circle {
	circle := calculateRadiusFromImage(image)
	center := calculateCenterFromImage(image)

	c.Radius = circle
	c.Offset = center

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
