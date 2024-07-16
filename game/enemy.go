package game

import "roguedef/trait"

type Enemy struct {
	*trait.Sprite
	*trait.Velocity
	intersect trait.Intersector
}

func (e *Enemy) Intersect() trait.Intersector {
	return e.intersect
}
