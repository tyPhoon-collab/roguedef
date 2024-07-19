package object

import (
	"roguedef/domain"
	"roguedef/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
	*system.Transform
	*domain.Status
	game      *Game
	object    *system.Object
	sprite    *system.Sprite
	velocity  *system.Velocity
	intersect system.Intersector
	player    *Player
}

func (e *Enemy) Register(g *Game, o *system.Object) {
	e.game = g
	e.object = o
}

func (e *Enemy) Intersect() system.Intersector {
	return e.intersect
}

func (e *Enemy) Update() {
	if e.Hp <= 0 {
		e.game.RemoveObject(e.object.ID)
	}
	e.velocity.Update()
	e.MoveTo(e.Pos)
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	e.sprite.Draw(screen)
}

func (e *Enemy) OnIntersect(other *system.Object) {
	if o, ok := other.Data.(domain.Attacker); ok {
		o.Attack(e.Status)
	}
}

func (e *Enemy) OnRemove() {
	if e.player != nil {
		e.player.AddExp(10)
	}
}

func (e *Enemy) WithPlayer(player *Player) *Enemy {
	e.player = player
	return e
}

func (o *Enemy) WithStatusModifier(modifier func(*domain.Status)) *Enemy {
	if o.Status == nil {
		panic("status is nil")
	}
	modifier(o.Status)
	return o
}

func NewEnemy() *Enemy {
	enemyImage, _, err := ebitenutil.NewImageFromFile("resources/images/gopher.png")
	if err != nil {
		panic(err)
	}
	transform := system.NewTransform()
	transform.Scale = transform.Scale.MulScalar(0.5)

	return &Enemy{
		Transform: transform,
		sprite:    system.NewSprite(enemyImage).WithTransform(transform),
		velocity:  system.NewVelocity().WithTransform(transform).With(Vec2{X: 0, Y: 1}),
		intersect: system.NewCircle().WithTransform(transform).FromImage(enemyImage),
		Status:    domain.NewStatus(10),
	}
}
