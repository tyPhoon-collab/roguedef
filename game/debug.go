package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Debug struct {
	showIntersects bool
	game           Game
}

func (d *Debug) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		d.showIntersects = !d.showIntersects
	}
}

func (d *Debug) Draw(screen *ebiten.Image) {
	if d.showIntersects {
		for _, o := range d.game.intersects {
			o.Draw(screen)
		}
	}
}

func NewDebug(game Game) *Debug {
	return &Debug{
		game: game,
	}
}
