package game

import (
	"roguedef/domain"
	"roguedef/resources"
	"roguedef/system"
	"time"
)

type BulletSpawner struct {
	*system.Sprite
	transform *system.Transform
	game      *Game
	bSpeed    float64
	bDamage   int
	*system.Looper
}

func (r *BulletSpawner) Register(g *Game, o *system.Object) {
	r.game = g
}

func (r *BulletSpawner) Update() {
	r.Looper.Update()
}

func (r *BulletSpawner) addBullet() {
	target, ok := r.calculateTarget()

	if !ok {
		return
	}

	distance := r.Pos.Distance(target.Pos)
	t := distance / r.bSpeed
	predictedPos := target.PredictPos(t)

	dir, err := r.Pos.DirTo(predictedPos)
	if err != nil {
		return
	}

	bullet := NewBullet(dir.MulScalar(r.bSpeed), &domain.AttackStatus{Damage: r.bDamage})
	bullet.Pos = r.Pos

	r.game.AddObjectWithData(bullet)
}

func (r *BulletSpawner) calculateTarget() (*Enemy, bool) {
	nearestDistance := -1.0
	var nearestEnemy *Enemy

	// get max y enemy
	for o := range r.game.ObjectsByTag("enemy") {
		enemy := o.Data.(*Enemy)
		distance := enemy.Pos.Y

		if nearestDistance == -1.0 {
			nearestDistance = distance
			nearestEnemy = enemy
		} else if nearestDistance < enemy.Pos.Y {
			nearestDistance = distance
			nearestEnemy = enemy
		}
	}

	if nearestDistance == -1.0 {
		return nil, false
	}
	return nearestEnemy, true
}

func NewBulletSpawner(pos Vec2) *BulletSpawner {
	img := system.LoadImage(resources.PendantImage)
	transform := system.NewTransform()
	transform.MoveTo(pos)
	transform.Scale = transform.Scale.MulScalar(2)
	sprite := system.NewSprite(img).WithTransform(transform)
	// sprite := system.NewEmptySprite().WithTransform(transform)

	s := &BulletSpawner{bSpeed: 10, bDamage: 10, Sprite: sprite, transform: transform}

	s.Looper = system.NewLooper(time.Second, s.addBullet)

	return s
}
