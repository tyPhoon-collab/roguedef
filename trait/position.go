package trait

import (
	"roguedef/vector"
)

type Vec2 = vector.Vec2

type HasPosition struct {
	Pos Vec2
}

func (p *HasPosition) Move(delta Vec2) {
	p.Pos = p.Pos.Add(delta)
}

func (p *HasPosition) MoveTo(x, y float64) {
	p.Pos = Vec2{X: x, Y: y}
}
