package object

import (
	"roguedef/system"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type BulletSpawner struct {
	player *Player
	game   *Game
}

func (r *BulletSpawner) Register(g *Game, o *system.Object) {
	r.game = g
}

func (r *BulletSpawner) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		r.spawnBullet()
	}
}

func (r *BulletSpawner) spawnBullet() {
	target := r.calculateTarget()
	dir, err := r.player.Pos.DirTo(target)
	if err != nil {
		return
	}

	bullet := NewBullet(dir.MulScalar(10))

	bullet.Pos = r.player.Pos

	obj := r.game.AddObjectWithData(bullet)

	r.game.AddTaskAfter(3*time.Second, func() error {
		r.game.RemoveObject(obj.ID)
		return nil
	})
}

func (r *BulletSpawner) calculateTarget() Vec2 {
	nearestDistance := -1.0
	nearestEnemy := Vec2{X: -1, Y: -1}

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
	return nearestEnemy
}

func NewBulletSpawner(player *Player) *BulletSpawner {
	return &BulletSpawner{player: player}
}
