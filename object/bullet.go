package object

import (
	"roguedef/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct {
	*system.Transform
	game      *Game
	object    *system.Object
	sprite    *system.Sprite
	velocity  *system.Velocity
	intersect system.Intersector
}

func (b *Bullet) Register(g *Game, o *system.Object) {
	b.game = g
	b.object = o
}

func (b *Bullet) Update() {
	b.velocity.Update()
	if b.game.IsOutside(b.Pos) {
		b.game.RemoveObject(b.object.ID)
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	b.sprite.Draw(screen)
}

func (b *Bullet) Intersect() system.Intersector {
	return b.intersect
}

func (b *Bullet) OnIntersect(other *system.Object) {
	if _, ok := other.Data.(*Enemy); ok {
		b.game.RemoveObject(b.object.ID)
	}
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
