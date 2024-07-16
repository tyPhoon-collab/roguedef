package game

import (
	"roguedef/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct {
	*system.Transform
	sprite    *system.Sprite
	velocity  *system.Velocity
	intersect system.Intersector
}

func (b *Bullet) Register(o *system.Object) {}

func (b *Bullet) Update() {
	b.velocity.Update()
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	b.sprite.Draw(screen)
}

func (b *Bullet) Intersect() system.Intersector {
	return b.intersect
}

func NewBullet(vel Vec2) *Bullet {
	bulletImage, _, err := ebitenutil.NewImageFromFile("resources/images/gopher.png")

	if err != nil {
		panic(err)
	}

	transform := system.NewTransform()
	transform.Scale = transform.Scale.MulScalar(0.2)

	return &Bullet{
		Transform: transform,
		sprite:    system.NewSprite(bulletImage).WithTransform(transform),
		velocity:  system.NewVelocity().WithTransform(transform).With(vel),
		intersect: system.NewCircle().WithTransform(transform).FromImage(bulletImage),
	}
}
