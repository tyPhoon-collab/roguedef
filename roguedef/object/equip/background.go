package equip

import (
	"image/color"
	"roguedef/system"
)

type Background struct {
	*system.Sprite
}

func (b *Background) Register(g *system.Game, o *system.Object) {}

func NewBackground(width int, height int) *Background {
	b := &Background{system.NewColorSprite(width, height, color.Gray{60}).WithAlignment(system.AlignTopLeft)}

	return b
}
