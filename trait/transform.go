package trait

import (
	"roguedef/vector"
)

type Vec2 = vector.Vec2

type Transform struct {
	Pos   Vec2
	Rot   float64
	Scale Vec2
}

func (p *Transform) Move(delta Vec2) {
	p.Pos = p.Pos.Add(delta)
}

func (p *Transform) MoveTo(x, y float64) {
	p.Pos = Vec2{X: x, Y: y}
}

func (p *Transform) String() string {
	return p.Pos.String()
}

func NewTransform() *Transform {
	return &Transform{
		Pos:   Vec2{X: 0, Y: 0},
		Rot:   0,
		Scale: Vec2{X: 1, Y: 1},
	}
}
