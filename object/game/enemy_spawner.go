package game

import (
	"roguedef/domain"
	"roguedef/system"
	"time"
)

type EnemySpawner struct {
	game         *Game
	spawnRange   Rect
	player       *Player
	phaseManager *PhaseManager
	*system.Looper
}

func (s *EnemySpawner) Register(g *Game, o *system.Object) {
	s.game = g
	s.player = g.ObjectByTag("player").Data.(*Player)
	s.phaseManager = g.ObjectByTag("phase_manager").Data.(*PhaseManager)
}

func (s *EnemySpawner) Update() {
	s.Looper.Update()
}

func (s *EnemySpawner) addEnemy() {
	phase := s.phaseManager.phase

	// enemy := NewEnemyFromResource(resources.GopherImage)

	t, status := domain.EnemyBlueprintByPhase(phase)
	domain.ModifyStatusByPhase(phase, &status)

	var enemy *Enemy

	switch t {
	case domain.EnemyTypeSquare:
		enemy = NewEnemySquare(24)
	case domain.EnemyTypeTriangle:
		enemy = NewEnemyTriangle(24)
	}

	// enemy := NewEnemyTriangle(24).
	enemy = enemy.
		WithPlayer(s.player).
		WithStatus(&status)

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
