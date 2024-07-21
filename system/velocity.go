package system

type Velocity struct {
	Vel   vec
	Scale float64
	*Transform
}

func (v *Velocity) Update() {
	v.Move(v.ScaledVel())
}

func (v *Velocity) ScaledVel() vec {
	return v.Vel.MulScalar(TimeScale * v.Scale)
}

func (v *Velocity) With(velocity vec) *Velocity {
	v.Vel = velocity
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
		Vel:   vec{},
		Scale: 1,
	}).WithTransform(nil)
}
