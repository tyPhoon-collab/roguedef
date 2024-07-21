package game

import (
	"roguedef/domain"
	"roguedef/system"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	*system.Transform
	*domain.Status
	game        *Game
	object      *system.Object
	velocity    *system.Velocity
	intersector system.Intersector
	player      *Player
	drawer      system.Drawer
}

func (e *Enemy) Register(g *Game, o *system.Object) {
	e.game = g
	e.object = o
}

func (e *Enemy) Intersect() system.Intersector {
	return e.intersector
}

func (e *Enemy) Update() {
	if e.Hp <= 0 {
		e.game.RemoveObject(e.object.ID)
	}
	e.velocity.Update()
	e.MoveTo(e.Pos)
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	e.drawer.Draw(screen)
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

func NewEnemy(transform *system.Transform, drawer system.Drawer, intersector system.Intersector) *Enemy {
	return &Enemy{
		Transform:   transform,
		drawer:      drawer,
		velocity:    system.NewVelocity().WithTransform(transform).With(Vec2{X: 0, Y: 1}),
		intersector: intersector,
		Status:      domain.NewStatus(10),
	}
}

func NewEnemyFromResource(data []byte) *Enemy {
	transform := system.NewTransform()
	transform.Scale = transform.Scale.MulScalar(0.5)

	img := system.LoadImage(data)

	return NewEnemy(
		transform,
		system.NewSprite(img).WithTransform(transform),
		system.NewCircleFromImage(img).WithTransform(transform),
	)
}

func NewEnemyTriangle(size float64) *Enemy {
	transform := system.NewTransform()

	return NewEnemy(
		transform,
		system.NewVectorDrawerTriangle(size).WithTransform(transform),
		system.NewCircle(size/2).WithTransform(transform),
	)
}

func NewEnemySquare(size float64) *Enemy {
	transform := system.NewTransform()

	return NewEnemy(
		transform,
		system.NewVectorDrawerSquare(size).WithTransform(transform),
		system.NewCircle(size/2).WithTransform(transform),
	)
}
