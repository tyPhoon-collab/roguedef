package system

import "github.com/hajimehoshi/ebiten/v2"

type Intersector interface {
	Intersects(other Intersector) bool
	Draw(screen *ebiten.Image)
}

type IntersectHandler interface {
	OnIntersect(other *Object)
}

type IntersectHolder interface {
	Intersect() Intersector
}
