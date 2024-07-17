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
	bulletSpawner *BulletSpawner
}

func (p *Player) Register(g *Game, o *system.Object) {
	p.bulletSpawner = g.ObjectByTag("bullet_spawner").Data.(*BulletSpawner)
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.Draw(screen)
}

func (p *Player) String() string {
	return "Player: " + p.Pos.String()
}

func (p *Player) AddExp(exp int) {
	p.exp += exp

	level := p.Level()
	p.bulletSpawner.SetFrequency(time.Duration(1000.0/(level*5)+100) * time.Millisecond)
}

func (p *Player) Level() int {
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
