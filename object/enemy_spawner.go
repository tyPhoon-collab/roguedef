package object

import (
	"roguedef/system"
)

type EnemySpawner struct {
	game       *Game
	spawnRange Rect
	player     *Player
	*system.Looper
}

func (s *EnemySpawner) Register(g *Game, o *system.Object) {
	s.game = g
}

func (s *EnemySpawner) Update() {
	s.Looper.Update()
}

func (s *EnemySpawner) addEnemy() {
	enemy := NewEnemy().WithPlayer(s.player)
	enemy.Pos = s.spawnRange.RandomPoint()
	s.game.AddObjectWithData(enemy).WithTag("enemy")
}

func (s *EnemySpawner) WithPlayer(player *Player) *EnemySpawner {
	s.player = player
	return s
}

func NewEnemySpawner(spawnRange Rect) *EnemySpawner {
	s := &EnemySpawner{
		spawnRange: spawnRange,
	}
	s.Looper = system.NewLooper(-1, s.addEnemy)

	return s
}
