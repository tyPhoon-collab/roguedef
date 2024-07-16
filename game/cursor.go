package game

import (
	"roguedef/trait"

	"github.com/hajimehoshi/ebiten/v2"
)

type Cursor struct {
	*trait.Transform
	intersect trait.Intersector
}

func (p *Cursor) Update() {
	x, y := ebiten.CursorPosition()

	p.MoveTo(Vec2{X: float64(x), Y: float64(y)})
}

func (p *Cursor) Draw(screen *ebiten.Image) {
	p.intersect.Draw(screen)
}

func (p *Cursor) Intersect() trait.Intersector {
	return p.intersect
}

func NewCursor() *Cursor {
	transform := trait.NewTransform()
	return &Cursor{
		Transform: transform,
		intersect: trait.NewCircle().WithTransform(transform).WithRadius(5),
	}
}
