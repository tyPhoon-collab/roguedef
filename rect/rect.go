package rect

import (
	"math/rand/v2"
	"roguedef/vector"
)

type Vec2 = vector.Vec2

type Rect struct {
	Min Vec2
	Max Vec2
}

func (r Rect) RandomPoint() Vec2 {
	return Vec2{
		X: r.Min.X + rand.Float64()*(r.Max.X-r.Min.X),
		Y: r.Min.Y + rand.Float64()*(r.Max.Y-r.Min.Y),
	}
}
