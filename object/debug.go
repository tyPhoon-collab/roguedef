package object

import (
	"fmt"
	"roguedef/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Debug struct {
	showIntersects bool
	game           *Game
	levelManager   *LevelManager
}

func (d *Debug) Register(g *Game, o *system.Object) {
	d.game = g

	for o := range d.game.ObjectsByTag("level_manager") {
		d.levelManager = o.Data.(*LevelManager)
	}
}

func (d *Debug) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		d.showIntersects = !d.showIntersects
	}
}

func (d *Debug) Draw(screen *ebiten.Image) {
	tps := ebiten.ActualTPS()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nLevel: %d", tps, d.levelManager.level))

	if d.showIntersects {
		for _, o := range d.game.Intersects() {
			o.Draw(screen)
		}
	}
}

func NewDebug() *Debug {
	return &Debug{}
}
