package system

import (
	"roguedef/vector"
)

type vec = vector.Vec2
type gameBuilder = func(*Scene) *Game
type Routes = map[string]gameBuilder
