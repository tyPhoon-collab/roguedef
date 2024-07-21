package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"roguedef/domain"
	"roguedef/domain/upgrade"
	"roguedef/resources"
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
	v := p.ui.WaitShowUpgradeSelection()
	switch v {
	case upgrade.UpgradeFrequency:
		system.ScaleDuration(&p.bulletSpawner.Frequency, 0.75)
	case upgrade.UpgradePower:
		p.bulletSpawner.bDamage = system.MulIntByFloat(p.bulletSpawner.bDamage, 1.5)
	case upgrade.UpgradeSpeed:
		p.bulletSpawner.bSpeed *= 1.5
	}
}

func NewPlayer(pos Vec2) *Player {
	img := system.LoadImage(resources.PlayerBackImage)
	transform := system.NewTransform()
	transform.Scale = transform.Scale.MulScalar(4)
	transform.MoveTo(pos)

	player := &Player{
		Transform:  transform,
		sprite:     system.NewSprite(img).WithTransform(transform),
		expManager: domain.NewExpManager(),
	}

	return player
}
