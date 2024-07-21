package game

import (
	"roguedef/domain"
	"roguedef/resources"
	"roguedef/system"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	*system.Transform
	*domain.AttackStatus
	game      *Game
	object    *system.Object
	sprite    *system.Sprite
	velocity  *system.Velocity
	intersect system.Intersector
}

func (b *Bullet) Attack(status *domain.Status) {
	status.Hp -= b.Damage
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

func (b *Bullet) Priority() int {
	return 5
}

func (b *Bullet) Intersect() system.Intersector {
	return b.intersect
}

func (b *Bullet) OnIntersect(other *system.Object) {
	if _, ok := other.Data.(*Enemy); ok {
		b.game.RemoveObject(b.object.ID)
	}
}

func NewBullet(vel Vec2, status *domain.AttackStatus) *Bullet {
	img := system.LoadImage(resources.BulletImage)
	transform := system.NewTransform()
	transform.Scale = transform.Scale.MulScalar(2)

	return &Bullet{
		Transform:    transform,
		AttackStatus: status,
		sprite:       system.NewSprite(img).WithTransform(transform),
		velocity:     system.NewVelocity().WithTransform(transform).With(vel),
		intersect:    system.NewCircle().WithTransform(transform).FromImage(img),
	}
}
