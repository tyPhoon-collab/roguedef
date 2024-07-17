package object

import (
	"roguedef/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Debug struct {
	showIntersects bool
	player         *Player
	game           *Game
}

func (d *Debug) Register(g *Game, o *system.Object) {
	d.game = g
}

func (d *Debug) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		d.showIntersects = !d.showIntersects
	}
}

func (d *Debug) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, d.player.String())

	if d.showIntersects {
		for _, o := range d.game.Intersects() {
			o.Draw(screen)
		}
	}
}

func NewDebug(player *Player) *Debug {
	return &Debug{
		player: player,
	}
}
