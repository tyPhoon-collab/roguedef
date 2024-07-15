package game

import (
	"roguedef/trait"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Built struct {
	*trait.Sprite
	*trait.Velocity

	trait.Intersector
}

func (b *Built) Update() {
	b.Velocity.Update()

	b.Sprite.MoveTo(b.Velocity.Pos)
}

func (b *Built) Draw(screen *ebiten.Image) {
	b.Sprite.Draw(screen)
}

func NewBuilt() *Built {
	builtImage, _, err := ebitenutil.NewImageFromFile("resources/images/gopher.png")

	if err != nil {
		panic(err)
	}

	transform := trait.NewTransform()

	return &Built{
		Sprite:      trait.NewSprite(builtImage).WithTransform(transform),
		Velocity:    trait.NewVelocity().WithTransform(transform),
		Intersector: trait.NewCircle().WithTransform(transform).FromImage(builtImage),
	}
}
