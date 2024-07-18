package system

type Velocity struct {
	Velocity Vec2
	*Transform
}

func (v *Velocity) Update() {
	v.Move(v.Velocity.MulScalar(TimeScale))
}

func (v *Velocity) With(velocity Vec2) *Velocity {
	v.Velocity = velocity
	return v
}

func (v *Velocity) WithTransform(transform *Transform) *Velocity {
	if transform == nil {
		transform = NewTransform()
	}
	v.Transform = transform
	return v
}

func NewVelocity() *Velocity {
	return (&Velocity{
		Velocity: Vec2{},
	}).WithTransform(nil)
}
