package title

import (
	"image/color"
	"roguedef/system"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	*system.Sprite
}

func (b *Background) Register(g *Game, o *system.Object) {}

func (b *Background) Draw(screen *ebiten.Image) {
	b.Sprite.Draw(screen)
}

func NewBackground(width int, height int) *Background {
	b := &Background{system.NewColorSprite(width, height, color.Gray{60}).WithAlignment(system.AlignTopLeft)}

	return b
}
