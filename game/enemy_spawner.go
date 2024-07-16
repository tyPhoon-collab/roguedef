package game

import (
	"time"
)

type EnemySpawner struct {
	game *Game
}

func (s *EnemySpawner) Update() {
	if s.game.frameCount%20 == 0 {
		s.spawnEnemy()
	}
}

func (s *EnemySpawner) spawnEnemy() {
	enemy := NewEnemy()

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
