package system

type Velocity struct {
	Velocity vec
	*Transform
}

func (v *Velocity) Update() {
	v.Move(v.Velocity.MulScalar(TimeScale))
}

func (v *Velocity) With(velocity vec) *Velocity {
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
		Velocity: vec{},
	}).WithTransform(nil)
}
