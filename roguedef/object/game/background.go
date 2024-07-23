package game

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

func (b *Background) Priority() int {
	return -1
}

func NewBackground() *Background {
	return &Background{
		system.NewColorSprite(320, 640, color.Gray{60}).WithAlignment(system.AlignTopLeft),
	}
}
