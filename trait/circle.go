package trait

import (
	"fmt"
	"image/color"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	ebitenVector "github.com/hajimehoshi/ebiten/v2/vector"
)

type Circle struct {
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
	return c.Pos.Sub(other.Pos).Len() < c.Radius+other.Radius
}

func (c *Circle) Trans() *Transform {
	return c.Transform
}

func (c *Circle) Draw(screen *ebiten.Image) {
	x := float32(c.Pos.X)
	y := float32(c.Pos.Y)
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

func calculateRadiusFromImage(image *ebiten.Image) float64 {
	return float64(image.Bounds().Dx()) / 2
}

func NewCircle(radius float64) *Circle {
	return (&Circle{
		Radius: radius,
	}).WithTransform(nil)
}

func NewCircleFromImage(image *ebiten.Image) *Circle {
	circle := NewCircle(calculateRadiusFromImage(image))

	x := float64(image.Bounds().Dx()) / 2
	y := float64(image.Bounds().Dy()) / 2

	circle.Move(Vec2{X: x, Y: y})

	return circle
}
