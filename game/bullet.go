package game

import (
	"roguedef/trait"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct {
	*trait.Transform
	sprite    *trait.Sprite
	velocity  *trait.Velocity
	intersect trait.Intersector
}

func (b *Bullet) Update() {
	b.velocity.Update()
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	b.sprite.Draw(screen)
}

func (b *Bullet) Intersect() trait.Intersector {
	return b.intersect
}

func NewBullet(vel Vec2) *Bullet {
	bulletImage, _, err := ebitenutil.NewImageFromFile("resources/images/gopher.png")

	if err != nil {
		panic(err)
	}

	transform := trait.NewTransform()
	transform.Scale = transform.Scale.MulScalar(0.2)

	return &Bullet{
		Transform: transform,
		sprite:    trait.NewSprite(bulletImage).WithTransform(transform),
		velocity:  trait.NewVelocity().WithTransform(transform).With(vel),
		intersect: trait.NewCircle().WithTransform(transform).FromImage(bulletImage),
	}
}
