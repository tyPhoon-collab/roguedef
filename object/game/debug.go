package game

import (
	"fmt"
	"roguedef/domain"
	"roguedef/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Debug struct {
	showInfo      bool
	game          *Game
	ui            *UI
	phaseManager  *PhaseManager
	player        *Player
	bulletSpawner *BulletSpawner
	enemySpawner  *EnemySpawner
}

func (d *Debug) Register(g *Game, o *system.Object) {
	d.game = g
	d.ui = g.ObjectByTag("ui").Data.(*UI)
	d.phaseManager = g.ObjectByTag("phase_manager").Data.(*PhaseManager)
	d.player = g.ObjectByTag("player").Data.(*Player)
	d.bulletSpawner = g.ObjectByTag("bullet_spawner").Data.(*BulletSpawner)
	d.enemySpawner = g.ObjectByTag("enemy_spawner").Data.(*EnemySpawner)
}

func (d *Debug) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		d.showInfo = !d.showInfo
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		go d.ui.WaitShowGameOver()
	}
}

func (d *Debug) Draw(screen *ebiten.Image) {
	if !d.showInfo {
		return
	}

	tps := ebiten.ActualTPS()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
		"TPS: %0.2f\nPhase: %d\nExp: %d\nLevel: %d\nBulletFreq: %d\nSpawnFreq: %d\nExpToNextLevel: %d\nEnemyQueueLen: %d\n",
		tps,
		d.phaseManager.phase,
		d.player.expManager.Exp(),
		d.player.expManager.Level(),
		d.bulletSpawner.Frequency.Milliseconds(),
		d.enemySpawner.Frequency.Milliseconds(),
		domain.ExpToNextLevel(d.player.expManager.Level()),
		d.enemySpawner.queue.Len(),
	), 10, 200)

	for _, o := range d.game.Intersects() {
		o.Draw(screen)
	}
}

func (d *Debug) Priority() int {
	return 100
}

func NewDebug() *Debug {
	return &Debug{}
}
