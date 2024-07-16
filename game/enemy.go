package game

import (
	"roguedef/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
	*system.Transform
	game      *Game
	object    *system.Object
	sprite    *system.Sprite
	velocity  *system.Velocity
	intersect system.Intersector
}

func (e *Enemy) Register(o *system.Object) {
	e.object = o
}

func (e *Enemy) Intersect() system.Intersector {
	return e.intersect
}

func (e *Enemy) Update() {
	e.velocity.Update()
	e.MoveTo(e.Pos)
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	e.sprite.Draw(screen)
}

func (e *Enemy) OnIntersect(other *system.Object) {
	if _, ok := other.Data.(*Bullet); ok {
		e.game.RemoveObject(e.object.ID)
	}
}

func NewEnemy(game *Game) *Enemy {
	enemyImage, _, err := ebitenutil.NewImageFromFile("resources/images/gopher.png")
	if err != nil {
		panic(err)
	}
	transform := system.NewTransform()
	transform.Scale = transform.Scale.MulScalar(0.5)

	return &Enemy{
		Transform: transform,
		game:      game,
		sprite:    system.NewSprite(enemyImage).WithTransform(transform),
		velocity:  system.NewVelocity().WithTransform(transform).With(Vec2{X: 0, Y: 1}),
		intersect: system.NewCircle().WithTransform(transform).FromImage(enemyImage),
	}
}
