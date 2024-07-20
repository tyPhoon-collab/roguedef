package system

import (
	"roguedef/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

type vec = vector.Vec2
type gameBuilder = func(*Scene) ebiten.Game
type Routes = map[string]gameBuilder
