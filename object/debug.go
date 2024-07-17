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
	phaseManager   *PhaseManager
	player         *Player
}

func (d *Debug) Register(g *Game, o *system.Object) {
	d.game = g
	d.phaseManager = g.ObjectByTag("phase_manager").Data.(*PhaseManager)
	d.player = g.ObjectByTag("player").Data.(*Player)
}

func (d *Debug) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		d.showIntersects = !d.showIntersects
	}
}

func (d *Debug) Draw(screen *ebiten.Image) {
	tps := ebiten.ActualTPS()
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %0.2f\nPhase: %d\nExp: %d\nLevel: %d",
		tps,
		d.phaseManager.phase,
		d.player.exp,
		d.player.Level(),
	))

	if d.showIntersects {
		for _, o := range d.game.Intersects() {
			o.Draw(screen)
		}
	}
}

func NewDebug() *Debug {
	return &Debug{}
}
