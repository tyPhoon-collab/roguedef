package system

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawer interface {
	Draw(screen *ebiten.Image)
	Priority() int // Draw priority. This func will be implemented by [Transform]. So basically, you don't have to think about this.
}
