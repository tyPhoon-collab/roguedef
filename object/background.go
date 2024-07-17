package object

import (
	"roguedef/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Background struct {
	*system.Sprite
}

func (b *Background) Register(g *Game, o *system.Object) {}

func (b *Background) Draw(screen *ebiten.Image) {
	b.Sprite.Draw(screen)
}

func (b *Background) Priority() int {
	return -1
}

func NewBackground() *Background {
	image, _, err := ebitenutil.NewImageFromFile("resources/images/background.png")

	if err != nil {
		panic(err)
	}

	transform := system.NewTransform()
	transform.Scale = transform.Scale.MulScalar(5.0)

	return &Background{
		Sprite: system.NewSprite(image).WithTransform(transform),
	}
}
