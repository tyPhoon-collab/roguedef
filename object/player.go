package object

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"roguedef/system"
)

type Player struct {
	*system.Transform
	sprite        *system.Sprite
	exp           int
	level         int
	bulletSpawner *BulletSpawner
	ui            *UI
}

func (p *Player) Register(g *Game, o *system.Object) {
	p.bulletSpawner = g.ObjectByTag("bullet_spawner").Data.(*BulletSpawner)
	p.ui = g.ObjectByTag("ui").Data.(*UI)
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.Draw(screen)
}

func (p *Player) AddExp(exp int) {
	p.exp += exp
	p.checkLevel()
}

func (p *Player) checkLevel() {
	level := p.calculateLevel()

	if level != p.level {
		go p.setLevel(level)
	}
}

func (p *Player) setLevel(level int) {
	p.level = level

	system.TimeScale = 0
	<-p.ui.ShowUpgradeSelectionPopup()
	system.TimeScale = 1
	p.bulletSpawner.SetFrequency(time.Duration(1000.0/(level*5)+100) * time.Millisecond)
}

func (p *Player) calculateLevel() int {
	return p.exp/100 + 1
}

func NewPlayer(pos Vec2) *Player {
	playerImage, _, err := ebitenutil.NewImageFromFile("resources/images/gopher.png")

	if err != nil {
		panic(err)
	}

	transform := system.NewTransform()

	transform.MoveTo(pos)

	return &Player{
		Transform: transform,
		sprite:    system.NewSprite(playerImage).WithTransform(transform),
	}
}
