package game

import (
	"roguedef/trait"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct {
	*trait.Sprite
	*trait.Velocity

	intersect trait.Intersector
}

func (b *Bullet) Update() {
	b.Velocity.Update()

	b.Sprite.MoveTo(b.Velocity.Pos)
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	b.Sprite.Draw(screen)
}

func (b *Bullet) Intersect() trait.Intersector {
	return b.intersect
}

func NewBullet() *Bullet {
	bulletImage, _, err := ebitenutil.NewImageFromFile("resources/images/gopher.png")

	if err != nil {
		panic(err)
	}

	transform := trait.NewTransform()
	transform.Scale = transform.Scale.MulScalar(0.2)

	return &Bullet{
		Sprite:    trait.NewSprite(bulletImage).WithTransform(transform),
		Velocity:  trait.NewVelocity().WithTransform(transform),
		intersect: trait.NewCircle().WithTransform(transform).FromImage(bulletImage),
	}
}
