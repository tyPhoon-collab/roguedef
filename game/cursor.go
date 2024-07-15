package game

import (
	"roguedef/trait"

	"github.com/hajimehoshi/ebiten/v2"
)

type Cursor struct {
	trait.Intersector
	trait.Object
}

func (p *Cursor) Update() {
	x, y := ebiten.CursorPosition()

	p.Trans().MoveTo(Vec2{X: float64(x), Y: float64(y)})
}

func (p *Cursor) Draw(screen *ebiten.Image) {
	p.Intersector.Draw(screen)
}

func NewCursor() *Cursor {
	return &Cursor{
		Intersector: trait.NewCircle().WithRadius(5),
	}
}
