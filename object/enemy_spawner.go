package object

import (
	"roguedef/system"
	"time"
)

type EnemySpawner struct {
	game           *Game
	spawnRange     Rect
	frequency      time.Duration
	timeAccumulate time.Duration
}

func (s *EnemySpawner) Register(g *Game, o *system.Object) {
	s.game = g
}

func (s *EnemySpawner) Update() {
	s.timeAccumulate += system.DeltaTime
	s.spawn()
}

func (s *EnemySpawner) spawn() {
	if s.timeAccumulate >= s.frequency {
		s.addEnemy()
		s.timeAccumulate -= s.frequency
		s.spawn() // spawn again
	}
}

func (s *EnemySpawner) addEnemy() {
	enemy := NewEnemy()
	enemy.Pos = s.spawnRange.RandomPoint()

	obj := s.game.AddObjectWithData(enemy).WithTag("enemy")
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
