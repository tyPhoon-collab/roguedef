package game

import (
	"fmt"
	"roguedef/domain"
	"roguedef/ds"
	"roguedef/system"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	*system.Transform
	*domain.EnemyStatus
	game        *Game
	ui          *UI
	object      *system.Object
	velocity    *system.Velocity
	intersector system.Intersector
	player      *Player
	drawer      system.Drawer
}

func (e *Enemy) Register(g *Game, o *system.Object) {
	e.game = g
	e.object = o

	e.ui = g.ObjectByTag("ui").Data.(*UI)
}

func (e *Enemy) Intersect() system.Intersector {
	return e.intersector
}

func (e *Enemy) Update() {
	if e.Hp <= 0 {
		e.game.RemoveObject(e.object.ID)
	}
	e.velocity.Scale = e.Speed
	e.velocity.Update()
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	e.drawer.Draw(screen)
}

func (e *Enemy) OnIntersect(other *system.Object) {
	if o, ok := other.Data.(domain.Attacker); ok {
		ctx := o.Attack(&e.EnemyStatus.Status)
		x, y := ds.Unpack[int](e.Pos)
		rf := e.ui.AddTextAt(x, y, fmt.Sprintf("%d", ctx.AppliedDamage))
		e.game.AddTaskAfter(1*time.Second, func() error {
			rf()
			return nil
		})
	}
}

func (e *Enemy) PredictPos(proceed float64) Vec2 {
	return e.Pos.Add(e.velocity.ScaledVel().MulScalar(proceed))
}

func (e *Enemy) OnRemove() {
	if e.player != nil {
		e.player.AddExp(e.Exp)
	}
}

func (e *Enemy) WithPlayer(player *Player) *Enemy {
	e.player = player
	return e
}

func (o *Enemy) WithStatus(status *domain.EnemyStatus) *Enemy {
	o.EnemyStatus = status
	return o
}

func NewEnemy(transform *system.Transform, drawer system.Drawer, intersector system.Intersector) *Enemy {
	return &Enemy{
		Transform:   transform,
		drawer:      drawer,
		velocity:    system.NewVelocity().WithTransform(transform).With(Vec2{X: 0, Y: 1}),
		intersector: intersector,
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
