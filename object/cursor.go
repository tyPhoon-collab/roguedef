package object

import (
	"roguedef/system"

	"github.com/hajimehoshi/ebiten/v2"
)

type Cursor struct {
	*system.Transform
	intersect system.Intersector
}

func (p *Cursor) Register(g *Game, o *system.Object) {}

func (p *Cursor) Update() {
	x, y := ebiten.CursorPosition()

	p.MoveTo(Vec2{X: float64(x), Y: float64(y)})
}

func (p *Cursor) Draw(screen *ebiten.Image) {
	p.intersect.Draw(screen)
}

func (p *Cursor) Intersect() system.Intersector {
	return p.intersect
}

func NewCursor() *Cursor {
	transform := system.NewTransform()
	return &Cursor{
		Transform: transform,
		intersect: system.NewCircle().WithTransform(transform).WithRadius(5),
	}
}
