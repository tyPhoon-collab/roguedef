package object

import (
	"roguedef/system"
	"time"
)

type EnemySpawner struct {
	game       *Game
	spawnRange Rect
}

func (s *EnemySpawner) Register(g *Game, o *system.Object) {
	s.game = g
}

func (s *EnemySpawner) Update() {
	if s.game.FrameCount()%60 == 0 {
		s.spawnEnemy()
	}
}

func (s *EnemySpawner) spawnEnemy() {
	enemy := NewEnemy()
	enemy.Pos = s.spawnRange.RandomPoint()

	obj := s.game.AddObjectWithData(enemy)
	s.game.AddTaskAfter(10*time.Second, func() error {
		s.game.RemoveObject(obj.ID)
		return nil
	})
}

func NewEnemySpawner(spawnRange Rect) *EnemySpawner {
	return &EnemySpawner{
		spawnRange: spawnRange,
	}
}
