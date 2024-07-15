package trait

import "github.com/hajimehoshi/ebiten/v2"

type Intersector interface {
	Intersects(other Intersector) bool
	Draw(screen *ebiten.Image)
	Trans() *Transform
}
