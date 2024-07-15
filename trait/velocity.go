package trait

type Velocity struct {
	Velocity Vec2
	*Transform
}

func (v *Velocity) Update() {
	v.Move(v.Velocity)
}

func (v *Velocity) Set(velocity Vec2) {
	v.Velocity = velocity
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
