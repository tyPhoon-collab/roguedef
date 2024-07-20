package game

import (
	"roguedef/domain"
	"roguedef/system"
	"time"
)

type BulletSpawner struct {
	player  *Player
	game    *Game
	bSpeed  float64
	bDamage int
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

	dir, err := r.player.Pos.DirTo(target)
	if err != nil {
		return
	}

	bullet := NewBullet(dir.MulScalar(r.bSpeed), &domain.AttackStatus{Damage: r.bDamage})
	bullet.Pos = r.player.Pos

	r.game.AddObjectWithData(bullet)
}

func (r *BulletSpawner) calculateTarget() (Vec2, bool) {
	nearestDistance := -1.0
	nearestEnemy := Vec2{}

	for o := range r.game.ObjectsByTag("enemy") {
		enemy := o.Data.(*Enemy)

		distance := enemy.Pos.Distance(r.player.Pos)

		if nearestDistance == -1.0 {
			nearestDistance = distance
			nearestEnemy = enemy.Pos
		} else if nearestDistance > distance {
			nearestDistance = distance
			nearestEnemy = enemy.Pos
		}
	}

	if nearestDistance == -1.0 {
		return Vec2{}, false
	}
	return nearestEnemy, true
}

func NewBulletSpawner(player *Player) *BulletSpawner {
	s := &BulletSpawner{player: player, bSpeed: 10, bDamage: 10}

	s.Looper = system.NewLooper(time.Second, s.addBullet)

	return s
}
