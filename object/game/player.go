package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"roguedef/domain"
	"roguedef/domain/upgrade"
	"roguedef/system"
)

type Player struct {
	*system.Transform
	sprite        *system.Sprite
	bulletSpawner *BulletSpawner
	ui            *UI
	expManager    *domain.ExpManager
}

func (p *Player) Register(g *Game, o *system.Object) {
	p.bulletSpawner = g.ObjectByTag("bullet_spawner").Data.(*BulletSpawner)
	p.ui = g.ObjectByTag("ui").Data.(*UI)
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.Draw(screen)
}

func (p *Player) AddExp(exp int) {
	changed := p.expManager.AddExp(exp)

	if changed {
		go p.onLevelChanged()
	}
}

func (p *Player) onLevelChanged() {
	system.TimeScale = 0
	v := <-p.ui.ShowUpgradeSelection()
	switch v {
	case upgrade.UpgradeFrequency:
		system.ScaleDuration(&p.bulletSpawner.Frequency, 0.75)
	case upgrade.UpgradePower:
		p.bulletSpawner.bDamage = system.MulIntByFloat(p.bulletSpawner.bDamage, 1.5)
	case upgrade.UpgradeSpeed:
		p.bulletSpawner.bSpeed *= 1.5
	}
	system.TimeScale = 1
}

func NewPlayer(pos Vec2) *Player {
	playerImage, _, err := ebitenutil.NewImageFromFile("resources/images/player_back.png")

	if err != nil {
		panic(err)
	}

	transform := system.NewTransform()

	transform.Scale = transform.Scale.MulScalar(0.15)
	transform.MoveTo(pos)

	player := &Player{
		Transform:  transform,
		sprite:     system.NewSprite(playerImage).WithTransform(transform),
		expManager: domain.NewExpManager(),
	}

	return player
}
