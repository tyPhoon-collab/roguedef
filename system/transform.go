package system

type Transform struct {
	Pos   vec
	Rot   float64
	Scale vec
}

func (p *Transform) Move(delta vec) {
	p.Pos = p.Pos.Add(delta)
}

func (p *Transform) MoveTo(vec vec) {
	p.Pos = vec
}

func (p *Transform) String() string {
	return p.Pos.String()
}

func (p *Transform) Copy() *Transform {
	return &Transform{
		Pos:   p.Pos,
		Rot:   p.Rot,
		Scale: p.Scale,
	}
}

func (p *Transform) Priority() int {
	return int(p.Pos.Y)
}

func NewTransform() *Transform {
	return &Transform{
		Pos:   vec{X: 0, Y: 0},
		Rot:   0,
		Scale: vec{X: 1, Y: 1},
	}
}
