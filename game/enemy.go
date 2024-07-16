package game

import (
	"roguedef/trait"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
	*trait.Transform
	sprite    *trait.Sprite
	velocity  *trait.Velocity
	intersect trait.Intersector
}

func (e *Enemy) Intersect() trait.Intersector {
	return e.intersect
}

func (e *Enemy) Update() {
	e.velocity.Update()
	e.MoveTo(e.Pos)
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	e.sprite.Draw(screen)
}

func NewEnemy() *Enemy {
	enemyImage, _, err := ebitenutil.NewImageFromFile("resources/images/gopher.png")
	if err != nil {
		panic(err)
	}
	transform := trait.NewTransform()
	transform.Scale = transform.Scale.MulScalar(0.5)

	return &Enemy{
		Transform: transform,
		sprite:    trait.NewSprite(enemyImage).WithTransform(transform),
		velocity:  trait.NewVelocity().WithTransform(transform).With(Vec2{X: 0, Y: 1}),
		intersect: trait.NewCircle().WithTransform(transform).FromImage(enemyImage),
	}
}
