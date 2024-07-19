package object

import (
	"roguedef/system"
	"time"
)

type EnemySpawner struct {
	game       *Game
	spawnRange Rect
	player     *Player
	*system.Looper
}

func (s *EnemySpawner) Register(g *Game, o *system.Object) {
	s.game = g
	s.player = g.ObjectByTag("player").Data.(*Player)
}

func (s *EnemySpawner) Update() {
	s.Looper.Update()
}

func (s *EnemySpawner) addEnemy() {
	enemy := NewEnemy().WithPlayer(s.player)
	enemy.Pos = s.spawnRange.RandomPoint()
	s.game.AddObjectWithData(enemy).WithTag("enemy")
}

func NewEnemySpawner(spawnRange Rect, frequency time.Duration) *EnemySpawner {
	s := &EnemySpawner{
		spawnRange: spawnRange,
	}
	s.Looper = system.NewLooper(frequency, s.addEnemy)

	return s
}
