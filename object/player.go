package object

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"roguedef/system"
)

type Player struct {
	*system.Transform
	sprite *system.Sprite
}

func (p *Player) Register(g *Game, o *system.Object) {}

func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.Draw(screen)
}

func (p *Player) String() string {
	return "Player: " + p.Pos.String()
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
