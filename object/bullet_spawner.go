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
	bullet := NewBullet(Vec2{X: 0, Y: -10})

	bullet.Pos = r.player.Pos

	obj := r.game.AddObjectWithData(bullet)

	r.game.AddTaskAfter(3*time.Second, func() error {
		r.game.RemoveObject(obj.ID)
		return nil
	})
}

func NewBulletSpawner(player *Player) *BulletSpawner {
	return &BulletSpawner{player: player}
}
