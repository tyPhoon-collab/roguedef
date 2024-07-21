package system

import "roguedef/ds"

type vec = ds.Vec2
type gameBuilder = func(*Scene) *Game
type Routes = map[string]gameBuilder
