package game

import (
	"roguedef/system"
	"time"
)

type EnemySpawner struct {
	game       *Game
	spawnRange Rect
}

func (s *EnemySpawner) Register(o *system.Object) {}

func (s *EnemySpawner) Update() {
	if s.game.frameCount%60 == 0 {
		s.spawnEnemy()
	}
}

func (s *EnemySpawner) spawnEnemy() {
	enemy := NewEnemy(s.game)
	enemy.Pos = s.spawnRange.RandomPoint()

	obj := s.game.AddObjectWithData(enemy)
	s.game.AddTaskAfter(10*time.Second, func() error {
		s.game.RemoveObject(obj.ID)
		return nil
	})
}

func NewEnemySpawner(game *Game) *EnemySpawner {
	return &EnemySpawner{
		game: game,
	}
}
